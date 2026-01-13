# 20260113 - Concept: snowflake servers

Snowflake server란 서버 인스턴스를 생성한 이후 내부적으로 추가 설정을 진행한 형태의 서버를 부르는 말이다.

이러한 서버들은 초기 생성 이후 여러 ad-hoc 설정을 적용하면서 생긴다.

다음과 같은 관리 측면에서의 문제점을 가진다.

- undocumented 환경이기 때문에 서버의 전체 설정을 쉽게 파악하기 어렵다.
- 시간이 지날수록 설정을 변경하기 어려워진다.
- 쉽게 재현할 수 없는 서버의 형태가 된다. image 생성 등을 통해 가능하기는 하지만
  여전히 간단하지 않은 형태인 것 같음

-----

thoughts: ad-hoc configuration을 적용한 형태의 서버가 가질 수 있는 문제점을 요약하기 위해 생긴 용어라고 생각한다.
automation tool 설명할 때 좋을 듯

## Sources

- [Snowflake Server - Martin Fowler](https://martinfowler.com/bliki/SnowflakeServer.html)
