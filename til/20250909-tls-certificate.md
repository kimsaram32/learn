# 20250909 - TLS certificate

TLS에선 secure connection을 만들기 위해 certificate를 사용한다.
이는 일반적으로 CA (Certificate Authority) 에서 발급한다.
호스트에선 직접 발급하지 않고 CA에서 certificate를 불러오게 된다.

대표적인 CA로 Let's Encrypt가 있다.

테스팅 목적으로 self-signed certificate를 사용할 수도 있다.
이 경우 client는 certificate를 보낸 end의 identity를 검증할 수 없기 때문에
암호화는 여전히 이루어지지만 신뢰성이 떨어진다.
즉 순수 개발 목적으로만 사용되어야 하고, production엔 쓰면 안 된다.

## Todo

- TLS handshake 복습
- certificate 어떻게 발급받는지?
