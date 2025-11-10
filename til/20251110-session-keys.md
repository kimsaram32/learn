# 20251110 - Session keys (short)

## 정의

- session key는 특정 session 동안 encryption을 위해 사용되는 session key이다.
- session마다 새로 생성되어야 하는데, 생성된 키는 일반적으로 asymmetric encryption 방식을 통해 교환된다.

## 쓰는 이유

왜 asymmetric key만 사용하지 않고, complexity를 추가하는가? 에 대한 대답이다.

asymmetric key는 computing 비용이 비싸다. 즉 모든 message마다 사용하기엔 부적합하다.

때문에 이보다 비용이 적은 symmetric key를 사용해 performance를 높인다.
