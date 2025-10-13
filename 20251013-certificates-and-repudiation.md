# 20251013 - Public-key certificate, repudiation

## Definitions

- entity: 시스템에 참여하는 대상
- identity: entity를 구분해주는 set of attributes
- identifier (name): 특정 identity를 가진 entity를 나타내주는 문자열

## Public-key certificates

public-key certificate는 특정 identifier와 public key를 연관지어준다.

이는 다음과 같이 설명할 수도 있다.
- public key를 identifier에 bind 해 준다.
- 어떤 entity가 public key를 own한다는 사실을 증명한다.

## Repudiaton, non-repudation service

### Repudiation

repudiation: sender의 identity가 특정 entity로 확인된 동작이 있을 때,
entity가 그 동작을 수행했었음을 부인하는 것이다.

e.g. 메시지의 생성이 어떤 entity로부터였다고 확인이 되지만,
그 entity는 해당 메시지의 생성을 부인하는 것

### Non-repudiation service

non-repudiation service: repudiation이 불가능하도록 만드는 것이다.

e.g. message를 어떤 entity가 생성했다는 것을 확인할 수 있다면, entity는 이를 부인할 수 없다.

## Sources

[RFC 4949](https://datatracker.ietf.org/doc/html/rfc4949)

## Thoughts

요새 배운 게 있어도 귀찮아서 안 올리고 있다가, 밀린 걸 올리기 시작했다...

초심을 찾자
