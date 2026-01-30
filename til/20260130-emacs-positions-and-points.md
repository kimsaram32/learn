# 20260130 - Emacs positions and points

## Positions

position은 기본적으로는 정수 값이며,
buffer에서 두 character 사이의 위치를 나타낸다.

- 최솟값은 첫 character 이전의 위치이므로 1이다.
- 최댓값은 마지막 character 다음의 위치이므로 (buffer size) + 1이다.

이때 buffer에 narrowing이 적용될 경우, 이러한 최소/최대 값은 더 제한될 수 있다.

UI 상에서 position을 표시할 때, highlight 되는 character는 point 다음의 character이다.
따라서 "character at point X" 는 엄밀히 말하면 "character after point X" 를 의미한다.

## Points

point는 매 buffer마다 저장되는 position 값이며,
현재 cursor 위치를 나타내는 데 사용한다.

## `point-min`, `point-max` function

`point-min`, `point-max` function은 각각 현재 buffer에서의
minimum point / maximum point 값을 리턴한다.

이것을 사용하지 않고 position의 일반적인 minimum / maximum 값인 1과 (buffer size) + 1을 하드코딩해서 사용하면
buffer가 narrowing 상태인 경우를 고려하지 않게 된다.
따라서 항상 `point-min`, `point-max` 함수를 사용해야 한다.

## Sources

- info manual
