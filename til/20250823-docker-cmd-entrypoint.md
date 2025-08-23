# 20250823 - Dockerfile CMD vs ENTRYPOINT

`ENTRYPOINT`는 default executable을 지정한다.
command를 넘기면 이것 뒤의 args로 들어간다.
-> 매번 똑같은 base command가 필요할 때 사용한다.

`CMD`는 container 실행 시 (docker run) command를 넘기지 않았을 때의 default command이다.

## Sources

- <https://spacelift.io/blog/docker-entrypoint-vs-cmd#what-is-the-difference-between-docker-entrypoint-and-cmd>
