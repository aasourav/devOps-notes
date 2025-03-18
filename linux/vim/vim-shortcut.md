. = this key is repeat last action
X = erase char before cursor
x = erase char in cursor
dd= delete the line
>G = is for set indentation for below all line
A = go to last in the line add append
a = go to nex to cursor and append
$ = go to last char of the line
C = remove text from current cursor position to last character and insert
s = remove currnt char and insert
S = remove current line and insert
I = go to the first line and insert
o = new line below and insert
O = new line top and insert
f = find char between current cursor to the last line
; = use last find forward cursor
. = use last find backword cursor

O = new line top and insert
f = find char between current cursor to the last line
; = use last find forward cursor
. = use last find backword cursor
/pattern = search forword pattern
?pattern = search backword pattern
n = use last search pattern forword
N = use last search pattern backword
:s/target/replacement = peform substitution for current line
& = use the previous substitution
db = delete backword from cursor's next back char to last char of word
x = delete char that cursor on it will delete left to right
dw = delete forword from cursor to next last char of word
b = move cursor backword wordwise(begening of the the next word)
w = move cursor forword (begining of the next word)
daw = delete current word cursor on also delete the white space
ctrl + a increment next number
ctrl + x decrement next number
cW  = delete current word from cursor on and insert
dap = delete entire paragraph

// insert mode
ctrl + w = delete back on word
ctrl + h = delete back one char
ctrl + u = delete back to start line

// visual mode
viw - select current word
v = enable character wise visual mode
V = enable line wise visual mode
ctrl + v = enable block wise visual mode
gv = reselect the last visual selection
vb = mark backword wordwise visual ( e for forword )
ve = mark forword wordwise visual (b for backword , we can move like this)
o =  go to the first or last side of visual area
j. = uppercase current word