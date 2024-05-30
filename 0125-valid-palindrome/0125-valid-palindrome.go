func isPalindrome(s string) bool {
    bytes := []byte(s)
    for i:=0; i<len(bytes); i++ {
        if 'A' <= bytes[i] && bytes[i] <= 'Z' {
            bytes[i] += 32
        } else if ('a' <= bytes[i] && bytes[i] <= 'z') || ('0' <= bytes[i] && bytes[i] <= '9') {

        } else {
            bytes = append(bytes[:i], bytes[i+1:]...)
            i--
        }
    }

    for i, j := 0, len(bytes)-1; 0 < len(bytes) && i < j; i, j = i+1, j-1 {
        if bytes[i] != bytes[j] {
            return false
            break;
        }
    }
    return true
}