# 20260120 - PID 1의 signal handling

## 복습: signal dispositions

Linux의 signal은 프로세스에게 특정 이벤트 정보를 전달하는 신호 메커니즘이다.

이때 signal이 전송되었을 때 프로세스가 실행해야 할 동작을 signal disposition이라고 부른다.

signal 마다 disposition을 지정해줄 수 있으며 (`sigaction()`, `signal()` 사용), 그렇게 하지 않은 경우엔 signal 마다 default disposition을 가진다.
이는 다음 중 하나이다.
- Ign: 시그널을 무시한다.
- Term: 프로세스를 종료한다.
- Core: core dump 파일을 생성하고 프로세스를 종료한다.
- Stop: 프로세스를 정지한다.
- Cont: 정지했던 프로세스를 재개한다.

## PID 1의 signal들은 default disposition이 Ign

위와 같이 일반적으로 프로세스의 signal은 disposition 미지정 시 적용되는
default dispositon이 signal 별로 존재한다.

그러나 PID 1 (init process) 의 경우에는
default disposition이 항상 Ign으로 설정된다.

즉 disposition을 설정하지 않은 시그널들을 받았을 경우 이들은 그냥 무시된다.

## Sources

- signal(7)
- [dumb-init: Why you need an init system](https://github.com/Yelp/dumb-init?tab=readme-ov-file#why-you-need-an-init-system)
