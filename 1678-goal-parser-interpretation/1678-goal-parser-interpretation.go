func interpret(command string) string {
    
    i, j := 0, 1
    strMap := map[ string]string{
        "G": "G",
        "()": "o",
        "(al)": "al",
    }
    answer := strings.Builder{}
    
    for j <= len( command) {
        
        val, ok := strMap[ command[ i: j]]
        if ok {
            
            answer.WriteString( val)
            i = j
        }
        j++
    }
    return answer.String()
}