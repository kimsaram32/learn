# 20251125 - procfs 기본

## procfs

procfs는 Linux에서 제공하는 pesudo filesystem이다.
커널의 low level DS를 file interface로 제공한다.
주로 process의 정보를 읽기 위해 사용한다.

VFS 상에서 `/proc` 에 mount 되어 있다.

## 프로세스 정보 확인

process id 별로 subdirectory가 존재하며, 그 안에 각 프로세스의 정보가 담겨 있다.
e.g. init process는 `/proc/1`
- `/proc/pid/status`: 프로세스의 정보를 human-readable 하게 제공한다.
- `/proc/pid/fd`: 프로세스가 연 file descriptor들을 나타내는
  subdirectory이다.
  file descriptor를 파일명으로 하며, 실제 연 파일에 대한 symbolic link 이다.
- `/proc/self`: 현재 프로세스 id에 대응하는 directory로의 magic link 이다.

## Todo

단순 프로세스 정보 제공용. 이 아니기 때문에, 커널 관련 부분을 더 알아보아야 한다.
