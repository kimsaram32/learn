# 20250918 - OAuth2 (recap)

예전에 배웠던 거 한번 더 대강 구체화하고 넘어가기

## 탄생 의도

서비스 외부의 3rd party application (client) 가 protected resource를 접근할 수 있게 하고 싶을 때,
resource의 소유자인 resource owner의 credential을 client에게 넘기지 않고도
resource owner로부터의 resource 접근 권한을 얻을 수 있게 하기 위해 만들어졌다.

내부 서비스에선 credential 기반 인증을 해도 상관 없지만
외부 접근자인 client에도 같은 방식을 사용하는 것은 보안상 매우 위험하다.

따라서 resource owner의 인증 자체는 client 밖에서 이루어지게 하며,
client는 credential 대신 resource owner로부터의 허가를 나타내는 authorization grant를 바탕으로 접근 권한을 얻는다.

## 용어

- client: 3rd party service (어떤 형태든 상관 없다)
- authorization server: authentication 및 authorization을 담당하는 server
- resource server: access token을 통해 protected resource를 제공하는 server
- resource owner: resource에 대한 접근 권한을 grant 해줄 수 있는 주체

authorization server와 resource server는 둘 다 같은 서비스 상에 속하므로
구현 시 한 server에서 구현될 수도 있고, 실제로 다른 server상에 존재할 수도 있다.

## 대략적인 flow

1. client는 resource owner로부터 authorization grant를 얻는다.
2. client는 얻은 authorization grant를 통해 authorization server로부터 access token을 받는다.
3. 이후 client는 access token을 통해 resource server로부터 resource owner가 grant해준 resource에 접근할 수 있다.

## Sources

<https://datatracker.ietf.org/doc/html/rfc6749>
