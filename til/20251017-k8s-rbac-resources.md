# 20251017 - K8s RBAC Resource

K8s에서의 RBAC 기반 authorization 구현용 resource들

## Role과 ClusterRole

Role과 ClusterRole은 RBAC에서의 role을 지정한다.

- `rules`를 작성하여, 접근을 허용할 리소스를 선택한다.
- deny 규칙은 작성할 수 없다 (접근 허용만 가능하다).
- Role의 경우, 선택되는 리소스는 해당 Role의 네임스페이스 내 리소스로 한정된다.
- ClusterRole의 경우, 선택되는 리소스는 전체 클러스터 범위이다.

## RoleBinding과 ClusterRoleBinding

RoleBinding과 ClusterRoleBinding은 Role 및 ClusterBinding을 사용자(들)에게 할당한다.

- role 할당 대상은 단일 유저 또는 set of users가 될 수 있다.
- RoleBinding에선 Role 또는 ClusterRole을 bind할 수 있다.
  이때 ClusterRole을 지정할 경우, 리소스 선택 범위가 RoleBinding이 속한 네임스페이스로 줄어든다.
  - ClusterRole 지정을 통해, 여러 네임스페이스에서의 access rule 공유를 구현할 수 있다.
- ClusterRoleBinding에선 ClusterRole을 bind할 수 있다.

## Sources

<https://kubernetes.io/docs/reference/access-authn-authz/rbac/>
