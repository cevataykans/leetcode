func generateParenthesis(n int) []string {
    
    paranthesis := make([]string, 0)
    return generate( paranthesis, n, n, "")
}

func generate( p []string, l, r int, s string) []string {
    
    if l < 0 || r < 0 {
        return p
    }
    
    if l == 0 && r == 0 {
        p = append( p, s)
        return p
    }
    
    p = generate( p, l - 1, r, s + "(")
    if r > l {
        p = generate( p, l, r - 1, s + ")")
    }
    return p
}