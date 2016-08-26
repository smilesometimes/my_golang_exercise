
# my_golang_exercise
 Daily Exercise,small codes.

## 1 tree

```tree
.
├── go_stl
│   ├── algorithm
│   │   └── algorithm.go
│   └── main.go
├── pipe
│   ├── a.txt
│   └── pipe.go
└── README.md
```
## 2 pipe example
```shell
cat a.txt

just  hello  
text 中国
golang

cat a.txt | ./pipe | grep hello

hello pipe just  hello  
hello pipe text 中国
hello pipe golang
```
## 3 go_stl example
```go
package main

import (
    "fmt"
    "my_golang_exercise/go_stl/algorithm"
)

func main() {
    s1 := []int{1, 2, 4, 5, 9}
    s2 := []int{2, 6, 7, 8}
    l := len(s1) + len(s2)
    d := make([]int, l, l)
    //merge two sorted int slice.
    algorithm.Merge(s1, s2, d)

    fmt.Println(d)
    //[1 2 2 4 5 6 7 8 9]
}

watch -n 1 tree

```













