# 20251020 - K8s authentication methods

K8s에서 사용하는 authentication 방법들에 대해서...

K8s에 접근하는 유저 종류는 크게 normal user와 service account로 나눌 수 있다.
이때 normal user는 K8s 외부에서 관리하므로, 인증 방법 역시 다양하다.

방법들이 여러 가지가 있지만 일단 정리한 것만 올려보기

## X.509 certificates

public-key certificate 기반으로 인증하는 방식이다.
인증 플로우는 일반적인 public-key certificate 이용 방식과 같다.

즉 대략 다음과 같다.
1. client에서 certificate를 제시한다.
2. certificate의 유효성 검증을 실시한다.
3. 성공했다면 client에서의 private key 기반 signature 인증을 실시한다.

이때 root CA는 직접 소유하게 된다.

## Webhook token authentication

client가 bearer token을 제시하면, 이것의 유효성 검증을 외부 provider에게 delegate 하는 방식이다.

1. client가 bearer token을 제시한다.
2. API server는 설정된 URL로 해당 token과 관련 정보를 담은 `TokenReview`를 전송한다.
3. 외부 provider에서 유효성 여부를 채운 `TokenReview`를 반환하고,
   API server는 이에 따라 사용자 인증을 진행한다.
