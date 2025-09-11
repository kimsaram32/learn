# 20250911 - Data integrity를 위해 사용되는 cryptographic hashing

secure communication의 요소 중 하나인 data integrity를 cryptographic hashing을 통해 보장할 수 있다.

## 기본 원리

sender는 메시지 (plaintext / ciphertext) 와 더불어 메시지의 hash 값을 함께 보낸다.
receiver는 메시지와 hash 부분을 분리하고, 받은 메시지의 hash 값이 받은 hash와 일치하는지 비교한다.

만약 중간에 메시지가 변조되었다면 hash가 여전히 일치하는 일은 사실상 없다.
cryptographic hashing의 특성 때문이다.

## Authentication key

그러나 이런 basic hashing은 패킷 전체를 위조하는 공격을 막을 수 없다.
hashing algorithm은 거의 대부분 public이므로, attacker가 이것을 알아내기만 하면
임의의 메시지에 스스로 hashing을 거쳐 보낼 수 있기 때문이다.

이를 방지하기 위해 hashing 시 메시지 자체에 authentication key를 섞는다.
이 key는 sender와 receiver만 알고 있어야 한다.
- 이러면 두 end system만이 올바른 hash를 생성할 수 있기에 attacker의 hash 생성을 방지할 수 있다.
- 메시지 + authentication key의 hash를 MAC (Message Authentication Code) 라고 부른다.

## Todo

- HMAC

## Sources

- Computer Networking - a Top-Down Approach (8th ed.) 8.3.1
