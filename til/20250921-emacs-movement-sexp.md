# 20250921 - Emacs에서의 S-expression movement

본래 Lisp 기준으로 만들어져서 S-expression 기반 움직임이라 부르지만
타 프로그래밍 언어에서 symmetric인 것들 (괄호, 따옴표...) 에도 적용될 수 있다

forward/backward
- `C-M-f` - sexp 기준 forward 이동
- `C-M-b` - sexp 기준 backward 이동

level로의 이동
- `C-M-u`
  - 한 level 위의 parentheses로 이동
  - 기본적으론 forward로 움직이지만 negative argument 주면 backward
- `C-M-d`
  - 한 level 아래의 parentheses로 이동
  - negative argument는 `C-M-u`와 반대 (기본적 backward -> negative 주면 forward)

level 안에서의 이동
- `C-M-n` - 같은 level에서의 다음 parentheses로 이동
- `C-M-p` - 같은 level에서의 이전 parentheses로 이동

## Thoughts

Vim의 hjkl에 대응하는 키가 Vim보다 더 불편해서 아쉬웠는데,
Emacs에선 그런 precise movement보단 syntactic movement를 더 중요시 여기기에 큰 문제가 없다고 한다.

즉 Vim 식의 movement가 아니니까 이것에 대응하는 키매핑을 찾지 말고 Emacs 식으로 생각하는 연습을 할 필요가 있음

## Sources

- Emacs
- Mastering Emacs
