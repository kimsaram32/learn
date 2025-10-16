# 20251016 - K8s Static Pods

## Static Pods

static Pod는 API server를 거치지 않고, kubelet에서 직접 생성하는 Pod이다.

일반적인 Pod 생성 프로세스에선,
kubelet이 API server에 올라간 Pod를 보고 이에 따라 생성을 진행한다.

그러나 static Pod는 API server 외부에 manifest가 올라간다.
kubelet은 그것을 바로 읽어서 생성 프로세스를 진행한다.

manifest들은 보통은 file system에 저장된다.

## Mirror Pod

생성은 API server를 거치지 않았더라도, 결국 Pod는 kubelet 상에서 존재하는 상태이다.

따라서 kubelet은 API server에 Pod resource를 올려 이를 알린다.
이렇게 올라간 mirror Pod는 API server 상에 read-only Pod로써 존재하게 된다.

## Use case: bootstrapping 시

클러스터 구축 시 주요 컴포넌트 (etcd, API server...) 들을 Pod로써 지정하는 경우가 있다.

일반 Pod들은 작동 중인 클러스터 없이는 올리는 것이 불가능하다.
따라서 이 경우 static Pod를 통해 컴포넌트들을 배포시킨다.

## Sources

<https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/>
