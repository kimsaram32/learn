# 20250925 - RFC 보면서 이것저것 정리

RFC 읽으면서 익힌 것들 정리하기

실용적인 것보단 그냥 TMI가 되어버린 것 같다;

## Underlying transport protocol

HTTP/1.1
- transport protocol을 명시하지 않는다.
- URI scheme에 따라 특정 프로토콜 (e.g. TCP) 을 사용하라고 요구하긴 하지만 전체 프로토콜의 요구사항이 아니다.

HTTP/2
- TCP를 사용한다고 명시되어 있다.
- TLS 사용 여부는 정해놓지 않았다.
- 실질적으론 브라우저들이 TLS 사용을 강제하고, 웹서버에서도 HTTPS만 지원하기에, TLS 위에서 돌아간다고 생각하면 편하다

HTTP/3
- UDP 기반의 QUIC 프로토콜을 사용한다고 명시되어 있다.

## HTTP/2 시작 방법

브라우저에서 HTTP/2 지원을 한다고 해도, 웹서버에선 아닐 수 있기 때문에 확인이 필요하다.

TLS를 사용할 경우
- TLS handshake 과정에서 프로토콜 negotiation이 이루어진다.
- 이에 따라 웹서버에서 HTTP/2를 지원할 경우 그것을 사용할 수 있다.

TLS 사용하지 않을 경우
- HTTP/1.1에서 Upgrade 헤더를 이용할 수 있다.
- 또는 HTTP/2로 보내보고 실패하면 1.1로 fallback 한다.
- 이전에도 언급했듯이 TLS 사용이 사실상 요구사항이라 이게 중요한지는 모르겠다
