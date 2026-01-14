# 20260114 - Emacs: fixing `next-line` (short)

## Problem

faster line movement를 위해 `next-line` 에 numeric argument를 먹이는
것을 시도해보았으나
제대로 작동하지 않았다.

## Prefix argument handling 문제인 줄 앎

`next-line` 의 `interactive` 내에서 numeric argument를 파싱하는 것이
잘못된 줄 알았다.

그래서 `advice-add` 로 `next-line` function을 override 하여,
numeric argument의 파싱 behavior를 변경하는 것으로 고치려고 했었다.

## `C-1` 바인딩 문제였음

advice 추가 후에도 작동하지 않길래 여러 가지를 실험해 보았다.
그 결과 `C-k C-1` 이 인식되지 않음을 알 수 있었다.

`C-1` 자체가 먹지 않는다는 소리였음

확인해보니 macOS shortcut (Mission Control 관련) 과 충돌하는 것이었다.
해당 shortcut을 disable 하는 것으로 해결하였다.
