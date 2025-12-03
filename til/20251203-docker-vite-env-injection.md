# 20251203 - Docker: Vite에서의 environment variable 주입 문제

## Problem

Vite를 사용하는 프론트엔드 코드베이스를 Docker 컨테이너로 실행하였다.
이때 환경변수를 옵션을 통해 넣어주었지만 적용이 되지 않고 있음을 확인하였다.

## 원인 1. Vite에선 빌드 타임에 환경변수 주입이 일어남

Vite에선 환경변수 주입이 빌드 타임에 끝난다.
즉 런타임, 즉 컨테이너 실행 시에 환경변수를 주입해도 이는 적용되지 않는다.

때문에 빌드 시 `ARG` 와 `ENV` 명령어, `--build-arg` option을 이용하여 해결하려고 하였다
([Stack Overflow](https://stackoverflow.com/questions/77486735/docker-with-vite-env-variables-are-undefined-inside-the-docker-container)).
그러나 여전히 주입되지 않았다.

## 원인 2. Vite는 .env 파일만 읽는 것으로 보임

[Vite 공식문서를 살펴본 결과, 환경변수 주입에는 .env 파일만 사용되는 것으로 보였다](https://vite.dev/guide/env-and-mode).
즉 프로세스의 환경변수는 사용하지 않는 것으로 보였다.

따라서 빌드 파이프라인에서 .env 파일을 생성 후 이미지를 빌드하는 것으로 해결하였다.

GitHub Actions workflow:
```
...
- name: Write .env file
    uses: SpicyPizza/create-envfile@v2.0
    with:
      envkey_VITE_REACT_APP_API_BASE_URL: ${{ secrets.SERVER_BASE_URL }}
...
```

원래는 쉘 상에서 직접 `.env` 파일을 만드는 스크립트를 했었는데, secret이 로그에서 그대로 노출된다는 문제가 있었다.
따라서 action을 사용하였다.

## Thoughts

Vite에서 프로세스의 환경변수를 주입하지 않는다는 게 이상하다...
내가 잘못 이해한 것이면 좋겠다.
