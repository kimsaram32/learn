# 20260109 - Tree를 이해하는 mental model

tree 공부하면서 느낀 것들을 그대로 정리하기...

## Search tree: 높이 유지

Search tree들의 많은 연산 (search, insertion, deletion...) 은 해당 트리의 높이에 비례한다.

따라서 search tree들의 기본 아이디어는 높이를 낮게 유지하여 연산을 효율적으로 수행하는 것이다.
구체적으로는 logarithmic 하게, 즉 N개의 node에 대하여 높이가 $O(\log N)$이 되도록 한다.

그러한 높이 유지 방법에 따라 여러 종류의 search tree가 생겨난 것이라고 설명할 수 있다.

- AVL trees: balance factor 유지
- 2-3 trees: subtree들의 height이 같도록 유지
- Red-black trees: 색 기반 property 유지

## Tree: operate - violate - fix

많은 tree DS들이 다음과 같이 동작한다.

1. 특정 property들을 가진다.
   - e.g. search tree: height을 낮게 유지하기 위한 property
   - e.g. binary heap: heap property
2. tree modification이 일어나는 연산들은 다음과 같이 동작한다.
   1. operate: "base" property 들을 깨지 않는 선에서 연산을 수행한다.
      보통 기반이 되는 더 단순한 tree DS의 property를 유지한다.
      - e.g. heap: complete BT를 유지하는 선에서
      - e.g. BST-based: BST의 inorder property? 를 유지하는 선에서
   2. violate: 특정 property들이 violate 된다.
   3. fix: 그것을 해결하기 위한 추가 연산을 실행한다. 즉 "fix the tree" 형태
