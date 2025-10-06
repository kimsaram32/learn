# 20251006 - SSH public key authentication

## SSH Authentication

로그인을 시도하는 local machine이 known이라는 것을 입증하는 과정이다.
인증 방법에는 여러 가지가 존재한다.

- password authentication
- public key authentication
- host-based authentication

## SSH public key authentication

local machine만이 private key를 가지고 있다는 전제하에 작동한다.

이러면 public key를 가지고 있다는 사실이 곧 authentication이 되므로, 이를 검증하는 방식이다.
이는 signature를 통해 달성한다.

다음과 같은 과정을 거친다 (일반적인 signature 검증법인듯..?).
1. remote machine에서 특정 message의 signing을 요청한다.
2. local machine에서 해당 message를 private key로 sign한 signature를 보낸다.
3. remote machine에서 해당 signature를 public key를 사용하여 검증한다.
4. 검증이 성공적이라면, 인증이 성공한다.

## 사용자별로 key pair 분리해야 함

key pair는 적어도 사용자별로는 분리되어야 한다.

사용자끼리 key pair를 공유하면 private key도 공유되는 것이다.
따라서 다음과 같은 문제점이 발생한다.

1. 사용자 중 한 명이 조직을 떠나게 되면 key pair를 다시 발급해서 모두 설정하게 만들어야 한다.
  'local machine 외엔 private key를 소지하지 않는다'가 전제인데,
  사람이 떠나면 그 사람의 machine은 이를 위배하게 만드므로 재발급해야 한다

2. '애초에 private key를 공유한다 = 보안이 약화된다' 이다.

분리했을 경우엔 이런 문제점이 발생하지 않는다.
조직을 떠난 사람의 public key만 remote machine에서 제거하면 된다.

## remote machine별로 분리할지는 tradeoff 적용

key pair를 remote machine별로 분리하는 것에 대해선 tradeoff가 존재한다.

분리할 경우 - security는 살리고, convenience는 약화한다.
- security: 강화된다. private key 하나를 털어도 그에 대응하는 remote machine만 털린다.
- convenience: remote machine이 많을 경우 설정이 불편해진다.
  machine마다 key 설정을 해주어야 한다.

공유할 경우 - security를 약화하고, convenience를 얻는다.
- security: 분리할 때보다 보안이 떨어질 수밖에 없다.
  private key 하나가 털리면, 이를 통해 인증하는 모든 remote machine이 털린다.
- convenience: local machine 설정이 간편해진다.
  machine마다 separate key를 배포하지 않아도 된다.

## Sources

- [RFC 4252](https://datatracker.ietf.org/doc/html/rfc4252)
- Stack Exchange
