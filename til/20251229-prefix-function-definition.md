# 20251229 - Prefix function definition

1. proper prefix:
   문자열 $S$의 proper prefix는 $S$의 모든 prefix 중 $S$가 아닌 것이다.

2. prefix function:
   prefix function은 어떤 문자열 $S$에 대한 배열 $\pi$이다.
   
   $0 \le i < n$인 $i$에 대하여
   $\pi[i]$는 $S[0...i]$의 proper prefix이면서 suffix인 substring 중 가장 긴 것의 길이이다.

## Informal explanation

이를 좀 더 informal 하게 설명하자면...
어떤 문자열의 prefix function은 그것의 prefix인 모든 substring과 연관되는 값을 저장하는 배열이다.
substring T에 대하여 그 값은 T의 substring 중 다음을 만족하는 것의 길이이다.

1. T의 prefix이면서 suffix이다.
2. T가 아니다 (<-> 길이가 T보다 작다).
