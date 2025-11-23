# 20251123 - Emacs Org Mode에서 프로젝트 단위 todo 추가하기

프로젝트를 Org Mode의 파일 단위로 관리하는데, 할 일이 생길 때마다 매번 파일을 옮겨다니며 todo entry를 추가하는 것이 힘들었다.

org-capture의 custom capture template 기능을 사용하여, todo entry를 프로젝트 단위로 추가할 수 있게 하였다.

```elisp
(defun me:todo-project-find-location ()
(let ((project-file-name
 (read-file-name "Project file name: " me:note-project-directory "" t nil
         (lambda (name)
           (not (or (backup-file-name-p name)
                (auto-save-file-name-p name)))))))
  (unless (string-empty-p project-file-name)
    (find-file (expand-file-name project-file-name me:note-project-directory))    
    (goto-char (point-min))
    (unless (search-forward "* Todo" nil t)
(goto-char (point-max))
(insert "* Todo")))))

(add-to-list 'org-capture-templates
     `("x" "Todo (in a project)"
       plain (function me:todo-project-find-location)
       "** TODO %?"))
```

## `read-file-name`

프로젝트명을 입력을 받아 지정하기 위해 사용하였다.
- minibuffer 상에서 file autocompletion을 제공한다.
- directory 옵션: 프로젝트들을 넣는 디렉토리를 지정하였다.
- predicate 옵션: backup, auto save file을 제외하기 위해 지정하였다.

## capture 위치 지정

org-capture-templates 상에서 target을 function으로 지정하여,
사용자가 고른 파일로의 insertion이 일어나도록 했다.

또한 `* Todo` heading의 child로 삽입되도록 설정했다.
만약 `* Todo` heading이 없으면 이를 하단에 삽입된다.

## 배우고 싶은 점

- 프로젝트명을 받을 때 read-file-name을 사용하는 것만이 방법일까? 더 나은 autocompletion 방법은 없을까?
- 함수를 더 깔끔하게 짜고 싶다
