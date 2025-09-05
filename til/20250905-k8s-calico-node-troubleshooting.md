# 20250905 - calico-node 작동 문제 해결하기

## 문제 발견

calico-node DaemonSet의 Pod가 각 node에서 잘 돌아가야 하는데,
한 노드에서만 터진 상태였다.

## BGP가 제대로 설정되지 않았다 -> BGP 이전의 문제이다

터진 Pod의 events를 봤더니 중간에 이런 게 있었다.
```
calico/node is not ready: BIRD is not ready: BGP not established with 192.168.219.126
```

더 파보니 BGP에 해당하는 179 port로 connection이 형성되지 않고 있었다.
`ping`은 되므로 네트워크가 아니라 프로세스 문제라고 예상했다.

`netstat -anpt | grep 179`를 돌려본 결과 예상대로 프로세스가 돌아가지 않고 있었다.
BGP 설정은 Calico에서 해주므로 이것 이전에서 문제가 발생하여 BGP 설정에 도달하지 못한 것이다.

## init container 상에서 문제가 발생했다 -> route 문제이다

BGP 설정 이전의 문제이다.
다시 events를 확인해봤더니 다음과 같은 부분이 있었다.
```
3m34s (x224 over 53m)    Warning   BackOff            Pod/calico-node-hjk2l       Back-off restarting failed container install-cni in pod calico-node-hjk2l_calico-system(117a7771-7bef-4af2-812a-038191fd0614)
```
install-cni init container에서 문제가 발생했음을 알 수 있다.

해당 로그를 확인했다.
```sh
$ kubectl logs calico-node-hjk2l install-cni -p
...
2025-09-05 11:27:55.559 [ERROR][1] cni-installer/token_watch.go 108: Unable to create token for CNI kubeconfig error=Post "https://10.96.0.1:443/api/v1/namespaces/calico-system/serviceaccounts/calico-cni-plugin/token": dial tcp 10.96.0.1:443: connect: no route to host
2025-09-05 11:27:55.559 [FATAL][1] cni-installer/install.go 499: Unable to create token for CNI kubeconfig error=Post "https://10.96.0.1:443/api/v1/namespaces/calico-system/serviceaccounts/calico-cni-plugin/token": dial tcp 10.96.0.1:443: connect: no route to host
```

Pod내의 route 상에서의 문제임을 알 수 있다.

## kube-proxy 문제

ChatGPT 말 따라서 kube-proxy Pod를 조회했다.
- Pod 상에서의 routing 문제엔 kube-proxy가 연관되어 있을 수 있다.

```sh
$ kubectl get po -n kube-system | grep proxy
kube-proxy-4mw2g                        1/1     Running   43 (125m ago)    28d
kube-proxy-7lcd8                        0/1     Error     3                24d
kube-proxy-7q8bp                        1/1     Running   103 (128m ago)   28d
kube-proxy-d2kcn                        1/1     Running   11 (127m ago)    24d

$ kubectl describe po kube-proxy-7lcd8 -n kube-system
...
Events:
  Type     Reason                  Age                   From     Message
  ----     ------                  ----                  ----     -------
  Normal   SandboxChanged          14m (x510 over 124m)  kubelet  Pod sandbox changed, it will be killed and re-created.
  Warning  FailedCreatePodSandBox  14m (x510 over 124m)  kubelet  Failed to create pod sandbox: rpc error: code = Unknown desc = failed to reserve sandbox name "kube-proxy-7lcd8_kube-system_738cd862-3810-448f-b906-9e3cd8310be4_4": name "kube-proxy-7lcd8_kube-system_738cd862-3810-448f-b906-9e3cd8310be4_4" is reserved for "dac527435fd80466a2dfb46fcb062c8989e166d8afb2edf1b9ee8e92b6565fa7"
  Normal   SandboxChanged          3m26s (x49 over 13m)  kubelet  Pod sandbox changed, it will be killed and re-created.
  Warning  FailedCreatePodSandBox  3m26s (x49 over 13m)  kubelet  Failed to create pod sandbox: rpc error: code = Unknown desc = failed to reserve sandbox name "kube-proxy-7lcd8_kube-system_738cd862-3810-448f-b906-9e3cd8310be4_4": name "kube-proxy-7lcd8_kube-system_738cd862-3810-448f-b906-9e3cd8310be4_4" is reserved for "dac527435fd80466a2dfb46fcb062c8989e166d8afb2edf1b9ee8e92b6565fa7"
```

sandbox name이라는 것이 중복됨을 알 수 있다.

## container runtime 문제 - forced shutdown

찾아보니 강제 shutdown 시 container runtime이 제대로 cleanup 되지 않아 발생하는 문제라고 한다.
지난번에 몇 번 그랬었던 것이 문제를 일으킨 것으로 보인다.

containerd를 재시작하고 남은 workload도 재시작하니 해결되었다.
