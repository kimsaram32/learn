# 20251015 - Minikube에 새 user 추가하기

Minikube에 새 normal user를 추가하고, RBAC 기반으로 authorization까지 수행한다.

## 목표

Pod를 read할 수 있는 `test-user` 이름의 user를 추가한다.

## 계획

1. certificate 발급
2. Role 및 RoleBinding 추가
3. kubeconfig에 credentials 및 context 추가
4. 테스트

```sh
$ base=~/.minikube/profiles/minikube/test-user
```

## Certificate 발급

- root key: `~/.minikube/ca.key`
- root certificate: `~/.minikube/ca.crt`

- 생성할 client key: `~/.minikube/profiles/minikube/test-user.crt`
- 생성할 client certificate: `~/.minikube/profiles/minikube/test-user.crt`

```sh
$ openssl genrsa -out ${base}.key 4096

$ openssl req -new -key ${base}.key -sha256 \
  -subj "/CN=test-user/O=test-team" \
  -out ${base}.csr

$ openssl x509 -req -days 3653 -in ${base}.csr \
  -copy_extensions copyall \
  -sha256 -CA ~/.minikube/ca.crt \
  -CAkey ~/.minikube/ca.key \
  -CAcreateserial \
  -out ${base}.crt
```

## Role 및 RoleBinding 추가

```sh
$ cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: test-role
  namespace: default
rules:
- apiGroups:
    - ""
  resources:
    - pods
  verbs:
    - get
    - list
    - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: test-role-binding
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: test-role
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: User
  name: test-user
EOF
```

## kubeconfig에 credentials 및 context 추가

```sh
$ kubectl config set-credentials test-user \
  --client-certificate="${base}.crt" \
  --client-key="${base}.key"

$ kubectl config set-context test \
  --cluster minikube
  --namespace default
  --user test-user
```

## 테스트

```sh
$ kubectl use-context test

$ kubectl auth whoami
ATTRIBUTE                                           VALUE
Username                                            test-user
Groups                                              [test-team system:authenticated]
Extra: authentication.kubernetes.io/credential-id   [X509SHA256=d002dda0b5ca6bc5177fea7478d610838a5827220126d8fe7cf9f27776b996cf]

$ kubectl get po # ok
$ kubectl describe po <name> # ok
$ kubectl get po -n kube-system # forbidden
```
