# 20251127 - lazy loading을 통해 zsh startup 속도 개선하기

zsh startup 속도가 느리다는 것을 느껴서 개선해보았다.

## 원인 찾기

1. 처음엔 그냥 config를 지워 가면서 소거법으로 찾았다.
2. zsh의 zprof 모듈 기반으로 profiling을 하였다 ([Profiling zsh startup](https://stevenvanbael.com/profiling-zsh-startup)).

## 원인: environment managers

`nvm`, `pyenv`, `jenv` 등이 속도 저하의 주된 원인임을 파악하였다.

## 해결: lazy loading

각 도구를 lazy load하는 방식으로 해결하였다.

1. 처음엔 함수를 통해 명령어를 override 하는 방식으로 직접 lazy load 하려고 하였다 ([GitHub Issue](https://github.com/nvm-sh/nvm/issues/2724#issuecomment-1336497491)).
2. 찾아보니 [zsh-lazyload](https://github.com/qoomon/zsh-lazyload) 플러그인이 있어서 해당 플러그인을 도입하여 해결하였다.
