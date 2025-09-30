# 20250930 - Loopback interface

virtual network interface이다.
- 패킷을 호출한 host에 그대로 전송한다.
- 대부분 127.0.0.1 IP address를 사용한다.

host 내의 IP address를 써도 작동은 하지만, physical interface의 상태에 영향을 받으므로 reliable 하지 않다.
- down 될 수 있다
- (당연하지만) host마다 IP address가 달라진다
- DHCP 등으로 변경될 수 있다

그러나 loopback interface는 virtual 하게 software 레벨에서만 동작하므로,
어느 host에서도 항상 그리고 일관적으로 작동함을 보장한다.
