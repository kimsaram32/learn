# 20251219 - Overlay networks

overlay network는 하위의 네트워크 (underlay network라고 부른다) 위에서 구동되는
논리적인 네트워크이다.

overlay network에서 동작하는 host들은 그것을 실제 네트워크와 같이 사용할 수 있으나,
실제 네트워킹 operation은 underlay network로 translate 되어 이루어진다.

다음과 같은 예시가 있다.

- VPN은 overlay network로써 구현되는 경우가 많다.
  e.g. WireGuard의 IP -> UDP encapsulate 방식
  사용자 입장에선 IP packet을 WireGuard의 interface로 보내는 것과 같다.
  
- VXLAN: L2 packet을 UDP packet에 encapsulate 하는 방식
