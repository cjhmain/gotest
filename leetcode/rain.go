/*
	力扣练习题：接雨水
*/
package leetcode

func find_top(height []int, li, ri int) int {
    if height[li] <= height[ri] {
        return height[li]
    } else {
        return height[ri]
    }
}

func quick_move(height []int, li, ri int) (int, int) {
    for li <= ri {
        hit := false
        if li + 1 < len(height) &&  height[li + 1] > height[li] {
            li = li + 1
            hit = true
        }
        if ri - 1 >= 0 && height[ri - 1] > height[ri] {
            ri = ri - 1
            hit = true
        }
        if !hit {
            break
        }
    }
    return li, ri
}

func find_left_boader(height []int, li, ri int) int {
    lin := -1
    for i := li + 1; i <= ri; i++ {
        if height[i] >= height[li] {
            lin = i
            break
        }
    }
    return lin
}

func find_right_boader(height []int, li, ri int) int{
    rin := -1
    for i := ri - 1; i >= li; i--{
        if height[i] >= height[ri] {
            rin = i
            break
        }
    }
    return rin
}

func add(height []int, l, r int) int {
    var rain int
    top := find_top(height, l, r)
    for i := l; i <= r; i++{
        v := top - height[i]
        if v > 0 {
            rain = rain + v
        }
    }
    return rain
}

func trap(height []int) int {
    size := len(height)
    if size <= 2 {
        return 0
    }

    var rain int

    li, ri := 0, size -1
    fli, fri := true, true
    for li <= ri && li + 1 <= ri {
        lin, rin := -1, -1
        li, ri = quick_move(height, li, ri)
        if fli {
            lin = find_left_boader(height, li, ri)
        }
        if fri {
            rin = find_right_boader(height, li, ri)
        }
        if lin < 0 && rin < 0 {
            rain = rain + add(height, li, ri)
            break
        }

        if lin >= 0 {
            rain = rain + add(height, li, lin)
            li = lin
        } else {
            fli = false
        }
        if rin >= 0 && rin >= li {
            rain = rain + add(height, rin, ri)
            ri = rin
        } else {
            fri = false
        }
    }
    return rain
}