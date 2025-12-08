# 20251208 - GitHub actions: environment variable injection 문제

## Problem

CD pipeline으로 생성된 프론트엔드 서버 image를 배포했는데, 환경변수가 제대로 들어가지 않음을 확인했다.

빌드는 다음과 같이 간단히 진행되고 있었다.
1. Vite 빌드를 위해 `.env` 파일을 생성한다. 값은 secret으로 들어간다.
2. Docker image를 빌드 후 push한다.

## .env 파일 생성까지는 잘 됨을 확인함

로컬에서 act 도구를 통해 `.env` 파일 생성까지 테스트해본 결과, 여기에는 문제가 없음을 확인할 수 있었다.

## 원인: docker/build-push-action의 context 옵션

1. [docker/build-push-action은 기본적으로 Git context 상에서 빌드한다](https://github.com/docker/build-push-action?tab=readme-ov-file#git-context).
2. path context를 사용하려면 `context` 옵션을 넘겨야 한다.
3. 이를 알지 못해 `context` 옵션을 사용하지 않았다.
   이러면 파일 시스템 상에 추가한 `.env` 파일은 반영되지 않는다.
   빌드 시 로컬 파일 대신 Git에 올라와 있는 스냅샷을 사용하기 때문이다.

해당 옵션 추가를 통해 해결하였다.

## Thoughts

- 액션을 쓰기 전 문서를 읽자...
- 가능한 경우 act를 통해 로컬 테스팅을 먼저 해보자
- 워크플로우 troubleshooting 시 action 단위로 문서를 읽어보면서 해결하면 좋은 듯 하다
