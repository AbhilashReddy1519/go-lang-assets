package main

import "fmt"

// slicesDemo runs examples showing common slice patterns and pitfalls.
func slicesDemo() {
    fmt.Println("== creation ==")
    var nilSlice []int
    empty := []int{}
    lit := []int{1, 2, 3}
    made := make([]int, 3, 5)
    fmt.Println("nilSlice", nilSlice, "nil?", nilSlice == nil)
    fmt.Println("empty", empty, "nil?", empty == nil)
    fmt.Println("lit", lit)
    fmt.Println("made", made, "len", len(made), "cap", cap(made))

    fmt.Println("== append & growth ==")
    s := []int{}
    for i := 0; i < 10; i++ {
        s = append(s, i)
        fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    }

    fmt.Println("== slicing arrays & shared backing array ==")
    arr := [5]int{10, 20, 30, 40, 50}
    a := arr[1:4]
    fmt.Println("a before", a)
    a[0] = 99
    fmt.Println("a after", a, "arr now", arr)

    fmt.Println("== copying to detach backing array ==")
    src := []int{1, 2, 3, 4}
    dst := make([]int, len(src))
    copy(dst, src)
    dst[0] = 100
    fmt.Println("src", src, "dst", dst)

    fmt.Println("== deleting an element ==")
    del := []int{1, 2, 3, 4, 5}
    i := 2 // remove element at index 2 (value 3)
    del = append(del[:i], del[i+1:]...)
    fmt.Println("after delete", del)

    fmt.Println("== reslicing up to capacity ==")
    base := make([]int, 2, 6)
    base[0], base[1] = 1, 2
    r := base[:4]
    fmt.Println("base len/cap", len(base), cap(base), "r len/cap", len(r), cap(r), r)

    fmt.Println("== multi-dimensional slices ==")
    mat := make([][]int, 3)
    for i := range mat {
        mat[i] = make([]int, 2)
        for j := range mat[i] {
            mat[i][j] = i*10 + j
        }
    }
    fmt.Println(mat)

    fmt.Println("== range loop pitfalls (address of loop var) ==")
    names := []string{"a", "b", "c"}
    ptrs := make([]*string, 0, len(names))
    for _, n := range names {
        // n is reused across iterations; taking &n would give same address
        v := n
        ptrs = append(ptrs, &v)
    }
    for _, p := range ptrs {
        fmt.Println(*p)
    }

    fmt.Println("== in-place filter (remove odd numbers) ==")
    nums := []int{1, 2, 3, 4, 5, 6}
    j := 0
    for _, v := range nums {
        if v%2 == 0 {
            nums[j] = v
            j++
        }
    }
    nums = nums[:j]
    fmt.Println("evens", nums)

    fmt.Println("== passing slices to functions ==")
    s2 := []int{1, 2, 3}
    fmt.Println("before mutate", s2)
    mutateSlice(s2)
    fmt.Println("after mutate", s2)

    fmt.Println("== slice of pointers vs slice of structs ==")
    type Thing struct{ Name string }
    aStructs := []Thing{{"x"}, {"y"}}
    aPtrs := []*Thing{{Name: "x"}, {Name: "y"}}
    aStructs[0].Name = "changed" // modifies copy
    aPtrs[0].Name = "changedPtr"
    fmt.Println(aStructs, aPtrs)
}

func mutateSlice(s []int) {
    if len(s) > 0 {
        s[0] = 999
    }
}
