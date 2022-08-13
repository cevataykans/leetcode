var (
    start int
    end int
)

func longestPalindrome(s string) string {
    
    start = 0
    end = 0
    
    for i := range s {
        evalPalindrome(s, i, i)
        evalPalindrome( s, i, i + 1)
    }
    
    return s[start+1:end]
}

func evalPalindrome( s string, l, r int) {
    
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    
    if end - start < r - l {
        start = l
        end = r
    }
}