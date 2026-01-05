# 20260105 - Selection sort, heap sort의 추상적인 절차

selection sort와 heap sort를 추상적인 절차의 구현 관점에서 설명할 수 있다.

둘은 다음과 같이 동일한 추상 절차를 구현한다.

sequence A의 selection sort는 다음과 같이 진행한다.
1. let U be a set and S be a sequence, initiallly empty.
   U는 정렬에 남은 요소, S는 정렬된 sequence를 저장한다.
2. U <- A
3. while U is not empty do
   1. let x be minimum element in U
   2. remove x from U
   3. append x to S

둘 다 unsorted / sorted 부분을 유지하며,
unsorted 부분에서 제일 작은 요소를 골라 sorted에 append 하는 동작을 반복한다.

그러나 그것의 구현 방식은 단순 minimum을 scan하는 $O(n)$ 방식과 (selection)
heap 자료구조를 이용하여 효율적으로 minimum을 얻는 $O(\log n)$ 방식으로 (heap) 나뉜다.

즉 추상 절차로 뭉뚱그리면 같아지지만, 그것을 구현하는 구체적인 스텝에 따라
$O(n^2)$ 인지 $O(n \log n)$ 인지 갈리게 된다.
