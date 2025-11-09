# 20251109 - 시그널 종류

사실 훨씬 많긴 하지만 내가 공부한 것 까지만 정리해보기

## SIGINT

- SIGINT는 쉘에서 Ctrl-C를 눌렀을 때 보내지는 시그널이다.
- 프로세스에서 실행 중이던 작업을 중단하라는 인터럽트의 의미를 가진다고 볼 수 있다

## SIGTSTP

- SIGTSTP은 쉘에서 Ctrl-Z를 눌렀을 때 보내지는 시그널이다.
- 프로세스를 중단해달라는 요청을 할 때 보내는 시그널이라고 할 수 있다.

## SIGTERM

- SIGTERM은 프로세스의 graceful termination 용도로 전송한다.
- 즉 프로세스는 SIGTERM에 대한 핸들러를 설정해 종료 전 cleanup 작업을 지정할 수 있다.
- `kill` 명령어에서 기본적으로 보내는 시그널이다.

## SIGKILL

- SIGKILL은 항상 프로세스를 terminate하는 시그널이다.
- 프로세스에서 disposition을 지정할 수 없다.
- 프로세스를 강제로 종료하는 데 사용된다. 즉 최후의 수단으로 사용해야 한다.

## SIGSTOP

- SIGSTOP은 항상 프로세스 실행을 중단하는 시그널이다.
- 프로세스에서 disposition을 지정할 수 없다.
- SIGKILL과 마찬가지로 프로세스를 강제로 중단하는 데 사용된다.

## SIGCONT

- SIGCONT은 프로세스 실행을 재개하는 용도로 보낸다.
- 내부적으로 special하게 처리한다고 하는데 더 알아볼 필요가 있음

