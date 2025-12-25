# 20251225 - X11과 Wayland의 compositor architecture

## 둘의 방식 비교

핵심은 compositor의 처리 방식이라고 생각한다.

X11은 compositor가 일반적인 application 프로그램으로써 display server 위에서 구동되는 형태이다.
즉 근본적으론 compositor와 다른 application 들이 구분되지 않는다.

반면 Wayland는 compositor와 application들이 직접적으로 소통하는 프로토콜의 형태이다.
중간에 display server를 거치지 않는다.

compositor가 특수한 형태의 application 인가, 아니면 완전히 dedicated 된 컴포넌트인가?
의 차이로 생각해볼 수 있다.

## X11의 문제점을 해결하는 관점에서의 설명

X11의 architecture는 현대의 modern compositor와의 인터랙션 방식을 핸들링하기엔 덜 효율적인 것으로 보인다.

내가 이해한 바가 맞다면, 최초의 display server가 수행하던 기능들을 compositor에게 맡기게 되었다.
이렇게 되면 display server는 compositor 와 application 사이의 unnecessary middleman이
되어버리는 문제가 발생한다.

Wayland는 compositor와 application이 직접적으로 소통하게 함으로써 이를 해결하는 것으로 보인다.

## Sources

<https://wayland.freedesktop.org/architecture.html>
