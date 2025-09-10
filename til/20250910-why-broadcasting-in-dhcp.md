# 20250910 - DHCP에서 broadcasting 하는 이유

처음 subnet 들어왔을 때의 procedure 상으로 보면,
DHCP message의 src/dest address가 client는 항상 0.0.0.0 -> 255.255.255.255,
server는 server IP -> 255.255.255.255라고 설명된다.

결국 둘 다 broadcast로 보내는 건데 왜 그런지 적어봤다.

## DHCPDISCOVER

이건 당연한 것이, client는 새로 들어왔기에 subnet에 대한 정보가 하나도 없다.

## DHCPOFFER

dest에 client IP를 넣기 -> 안됨

어떻게든 client의 MAC address로 패킷을 보낸다고 해도 client는 아직 IP address 설정을 하지 않았다.
따라서 client로 간 패킷이 그냥 drop 될 것이다.

## DHCPREQUEST

dest에 server IP를 넣기 -> server 여러 개인 경우 이슈 발생

제공받은 IP를 사용할 DHCP server에만 request 메시지를 보내도 될 것 같아 보이지만,
DHCP server가 여러 개일 때 문제가 발생한다.

이러면 다른 server엔 아무 메시지가 가지 않게 되고,
이 서버들은 offer 메시지가 제대로 가지 못한 것과 본인에게 request를 하지 않은 것을 구분할 수 없다.

## DHCPACK

dest에 client IP 넣기 -> configure 했다는 보장이 없음

client가 configure 했다고 가정하고 dest에 client IP를 넣을 수 있을 것 같았다.
그러나 그 가정이 옳지 않은 듯 하다. 즉 이미 설정해놓았을 것이라는 보장이 없다.

## Todo

RFC의 section 4.1에 DHCPOFFER와 DHCPACK 메시지에서 broadcast/unicast를 정하는 조건이 있다.

이거 뜯어보면서 이론적으로 되는 경우 알아봐도 좋을 듯

## Sources

- <https://superuser.com/questions/1536810/why-are-dhcp-messages-broadcast>
- <https://datatracker.ietf.org/doc/html/rfc2131>
