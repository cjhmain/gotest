/*
	有效括号
*/
package leetcode

func isValid(s string) bool {
    size := len(s)
    if size == 0 || size % 2 != 0 {
        return false
    }
    
    code := make(map[byte]int)
    code['('] = 1; code[')'] = 10
    code['['] = 2; code[']'] = 20
    code['{'] = 3; code['}'] = 30

    var index int = 0
    arr := make([]int, size)
    for i := 0; i < size; i++ {
        arr[index] = code[s[i]]
        if index > 0 && arr[index - 1] * 10 == arr[index]{
            arr = arr[:len(arr)-2]
            index = index - 1
        }else{
            index++
        }
    }
    if len(arr) == 0{
        return true
    }
    return false
}