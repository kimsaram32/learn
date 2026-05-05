# 20260505 - Alternative Group Theory: Existence of the right-hand identity element

## Problem

[[https://web.mit.edu/gleitz/www/Introduction%20to%20Logic%20-%20P.%20Suppes%20(1957)%20WW.pdf][Introduction to Logic - P. Suppes (1957)]]

5-2 113p - 5번 문제

Axioms:

1. \((x)(y)(z)(x \circ (y \circ z) = (x \circ y) \circ z)\)
2. \((x)(y)(\exists z)(x = y \circ z)\)
3. \((x)(z)(\exists y)(x = y \circ z)\)

Theorem 1. \((\exists y)(x)(x \circ y = x)\)

## Formal derivation

\begin{tabular}{l l l}
\{1\} & (1) \((x)(y)(z)(x \circ (y \circ z) = (x \circ y) \circ z)\) & P (Ax. (1)) \\
\{2\} & (2) \((x)(y)(\exists z)(x = y \circ z)\) & P (Ax. (2)) \\
\{3\} & (3) \((x)(z)(\exists y)(x = y \circ z)\) & P (Ax. (3)) \\
\{3\} & (4) \((\exists y)(x = y \circ w)\) & 3 US \(x/x\) \(w/z\) \\
\{3\} & (5) \(x = \alpha_{xw} \circ w\) & 4 ES \\
\{2\} & (6) \((\exists z)(w = w \circ z)\) & 2 US \(w/x\) \(w/y\) \\
\{2\} & (7) \(w = w \circ \beta_w\) & 6 ES \\
\{1,2\} & (8) \(\alpha_{xw} \circ (w \circ \beta_w) = (\alpha_{xw} \circ w) \circ \beta_w\) & 1 US
\(\alpha_{xw}/x\) \(w/y\) \(\beta_w/z\) \\
\{1,2\} & (9) \(\alpha_{xw} \circ w = (\alpha_{xw} \circ w) \circ \beta_w\) & 7,8 I \\
\{1,2,3\} & (10) \(x = x \circ \beta_w\) & 5,9 I \\
\varnothing & (11) \(x = x\) & I \\
\{1,2,3\} & (12) \(x \circ \beta_w = x\) & 10,11 I \\
\{1,2,3\} & (13) \((x)(x \circ \beta_w = x)\) & 12 UG \\
\{1,2,3\} & (14) \((\exists y)(x)(x \circ y = x)\) & 13 EG \\
\end{tabular}

## Thoughts

일단 rules of inference 따르면 틀린 부분은 없긴 한데... confidence가 없다.
멘토가 없으니까 한계를 느낀다.

-----

문제를 처음 접하고 풀어내기까지 2달이라는 시간이 걸렸다 (사실 이 기간 동안
순수하게 이 문제에 투자한 시간은 많지 않았긴 하다. 어찌됐건...). 행복함...
한편으로는 허무하기도 하다. 왜 이런 간단한 생각을 이제껏 하지 못했을까.

Formal derivation 문제는 informal 한 생각에서 벗어나서, 순수하게 rule을
밀어붙이는 식으로 푸는 접근법 그리고 intuitive understanding에서 translate 하는
식의 접근법이 있다고 생각한다. 이번에는 직관의 힘을 빌리는 쪽을 택했다 (풀이법
자체를 여기에 적지는 못했다).

두 접근법 중 어느 것이 좋다/나쁘다를 따지고 싶지는 않다만 내가 더 끌리는 쪽은
전자인 것 같다. 그 자체로 complete하고 precise 한 체계 안에서 생각하고 싶다.
내가 프로그래밍을 해서 그런가... 논리도 이런 식으로 받아들이게 되는 것 같다.
