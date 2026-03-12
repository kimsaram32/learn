# 20260312 - Quantifier 관련 rules of inference

[Introduction to Logic - P. Suppes (1957)](https://web.mit.edu/gleitz/www/Introduction%20to%20Logic%20-%20P.%20Suppes%20(1957)%20WW.pdf)
p. 98 기반으로 내 언어로 작성하기

-----

quantifier를 drop 하거나 (specification) add (generalization) 하는
것을 가능하게 하는 rules of inference 들이다.

1. US: from $(v)(S(v))$ derive $S(t)$.
   - for every variable occuring in $t$,
     all free occurrences of $v$ in $S(v)$ is not in
     the scope of any quantifier using that variable.

2. UG: from $S(v)$ derive $(v)(S(v))$.
   - v does not appear as a subscript.
   - v is not flagged.

3. ES: from $(\exists v)(S(v))$ derive $S(\alpha)$.
   - $\alpha$ does not already appear in the derivation.
   - subscript: denote all free variables in $S(\alpha)$ as a subscript
     in $\alpha$.

4. EG: from $S(\alpha)$ derive $(\exists v)(S(v))$.
   - v does not appear as a subscript.
   - no occurrences of $\alpha$ in the scope of a quantifier using $v$.
   - if $\alpha$ actually occurs in $S(\alpha)$, v is not flagged.

## Notation

- $v$ is any variable.
- $t$ is any term.
- $\alpha$ is any ambigious name.
- $S(v)$ is a sentence with zero or more occurrences of $v$.
- $S(t)$ is the sentence resulting from $S(v)$
  by substituting $t$ for every free occurrence of $v$ in $S$.
- $S(\alpha)$ is the sentence resulting from $S(v)$
  by substituting a for every free occurrence of $v$ in $S$.
