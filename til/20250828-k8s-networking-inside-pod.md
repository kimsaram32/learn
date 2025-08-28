# 20250828 - Pod 내부에서의 네트워크

한 Pod 내에서의 모든 container는 같은 network namespace를 공유한다.

- localhost를 통해 통신할 수 있다.
- container 끼리의 port가 중복되면 안 된다.

## 중복 port 실험

2개의 nginx image를 돌리는 container를 한 Pod에 넣었다.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
  - name: nginx0
    image: nginx
  - name: nginx1
    image: nginx
```

별도의 설정을 하지 않았으므로 두 container 모두 80번 port에 bind하려고 한다.

결국 한쪽은 port가 이미 사용 중이기 때문에 거부당한다.
```
2025/08/28 00:21:13 [emerg] 1#1: bind() to 0.0.0.0:80 failed (98: Address already in use)
```

## network 확인하기

busybox container 2개를 같은 Pod 안에서 돌린 후 `ip a` 커맨드로 주소를 확인해보았다.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test
spec:
  containers:
  - name: busybox0
    image: busybox
    command: ["/bin/sh"]
    args: ["-c", "while true; do sleep 10; done"]

  - name: busybox1
    image: busybox
    command: ["/bin/sh"]
    args: ["-c", "while true; do sleep 10; done"]
```

```sh
$ kubectl exec test -c busybox0 -- ip a
$ kubectl exec test -c busybox1 -- ip a
```

실행하면 ethernet 인터페이스에서 같은 ip address를 가지고 있음을 확인할 수 있다.
