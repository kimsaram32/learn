# 20251014 - K8s API access control (기초)

## API access control flow

1. 기본적으로 API server는 TLS 상에서 작동한다.
2. Authentication - 전체 request 기반 사용자 인증을 수행한다.
3. Authorization - resource 및 non-resource에 대한 사용자 접근 권한을 확인한다.
4. Admission control - request 자체를 수정하거나 거부할 수 있다.

## Authentication - types of users

K8s에 접근하는 user는 크게 두 종류로 구분할 수 있다.

1. normal users
  - K8s 내부적으로 관리하지 않는다.
  - 외부 provider에서 저장 및 관리를 담당한다.

2. service accounts
  - K8s에서 관리하는 managed user이다.

## Sources

- <https://kubernetes.io/docs/concepts/security/controlling-access/>
- <https://kubernetes.io/docs/reference/access-authn-authz/authentication/>
