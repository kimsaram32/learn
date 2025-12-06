# 20251206 - Emacs custom registers for Zettelkasten

## Goal

Emacs에서 Zettelkasten을 작성할 때, Zettel들 사이를 이동하는 용도로 register를 많이 사용한다.

이는 Org mode 네이티브하지 않은 기능이므로, Org 문서 상의 heading이 보이지 않고 단순 buffer의 point 기반이라서 불편했다.
custom register type을 작성해서 integration을 만들어보았다.

## Emacs register 동작 방식

Emacs의 register 들은 `register-alist` variable에 저장된다.
각 요소는 `(NAME . CONTENTS)` 형태로 되어 있다.
- name은 register 명 (1글자)이다.
- contents는 arbitrary type이 될 수 있다.

register 관련 함수를 generic의 형태로 작성하고, 이후 command에서 이를 참조하여 특정 register type에 대한 동작을 수행하는 형태이다.
- `jump-to-register`: `register-val-jump-to` generic function을 액세스한다.
- `insert-register`: `register-val-insert` generic function을 액세스한다.

즉 custom register를 만들어주려면 새로 타입을 작성한 후 해당 generic을 구현해주면 된다.

## Custom register 작성

```elisp
(cl-defstruct zk-zettel-register
  id)

(cl-defmethod register--type ((_regval zk-zettel-register)) 'zettel)

(cl-defmethod register-val-jump-to ((val zk-zettel-register) _arg)
  (org-id-goto (zk-zettel-register-id val)))

(cl-defmethod register-val-describe ((val zk-zettel-register) _verbose)
  (let* ((id (zk-zettel-register-id val))
         (marker (org-id-find id 'marker))
         (heading (if marker
                      (with-current-buffer (marker-buffer marker)
                        (save-excursion
                          (goto-char marker)
                          (org-get-heading t t t t)))
                    "Entry not found")))
    (princ (format "a zettel: %s (%s)" heading id))))

(cl-defmethod register-val-insert ((val zk-zettel-register))
  (let ((id (zk-zettel-register-id val)))
    (insert (format "[[id:%s][%s]]" id id))))
```

- `with-current-buffer` -> `save-excursion` 의 순서에 주의해야 한다.
  1. 호출 시점에선 현재 buffer가 register 상에서의 것이기 때문에, 그냥 `save-excursion` 호출 시 이 buffer 상에서의 state가 저장된다.
  2. 그러나 내가 원하는 것은 새로 jump 하는 buffer에서의 state restoration 이므로 해당 buffer로 `with-current-buffer` 를 먼저 호출해야 한다.

## 기존 command 사용하기

기존에 사용하던 `jump-to-register`, `insert-register` command를 그대로 사용하고 싶었다.
해당 커맨드마다 보여지는 register의 종류는 `register-command-info` generic에 저장된다.
즉 여기에 내가 추가한 register의 type을 넣어주어야 커맨드 상에서 이 register를 쓸 수 있다.

```elisp
(cl-defmethod register-command-info ((_command (eql insert-register)))
  (make-register-preview-info
   :types '(string number zettel)
   :msg "Insert register `%s'"
   :act 'insert
   :smatch t
   :noconfirm (memq register-use-preview '(nil never))))

(cl-defmethod register-command-info ((_command (eql jump-to-register)))
  (make-register-preview-info
   :types  '(window frame marker kmacro
             file buffer file-query zettel)
   :msg "Jump to register `%s'"
   :act 'jump
   :smatch t
   :noconfirm (memq register-use-preview '(nil never))))
```

기존 것을 override 하는 형태가 마음에 들지 않으나... 생각나는 대안이 없어서 그대로 진행했다. 고칠 필요가 있다.

## Storing

현재 zettel을 register 상에 store 하는 command를 작성 후,
기존의 `point-to-register` 에 사용되던 바인딩을 override 하였다.

```elisp
(defun zk-zettel-to-register (register)
  "Store current zettel at point in REGISTER.
Interactively, prompt for REGISTER using `register-read-with-preview'."
  (interactive (list (register-read-with-preview
                      "Store zettel in register: ")))
  (add-hook 'kill-buffer-hook 'register-swap-out nil t)
  (set-register register (make-zk-zettel-register :id (me:entry-zettel-id))))

(keymap-set me:zk-mode-map "C-x r SPC" #'zk-zettel-to-register)
```
