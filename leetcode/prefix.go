/*
	最长公共前缀
*/
package leetcode

func compare(s1, s2 []rune) string {
    var count int
    var max int = len(s1)
    if len(s2) < max {
        max = len(s2)
    }

    ans := make([]rune, max)
    for i := 0; i < max; i++ {
        if s1[i] == s2[i] {
            ans[i] = s1[i]
            count = count + 1
        }else{
            break
        }
    }
    if count > 0 {
        ans = ans[0:count]
        return string(ans)
    }
    return ""
}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
        return ""
    }else if len(strs) == 1{
        return strs[0]
    }

    ans := strs[0]
    for i := 1; i < len(strs); i++ {
        if ans == "" || strs[i] == "" {
            ans = ""
            break
        }
        ans = compare([]rune(ans), []rune(strs[i]))
    }
    return ans
}