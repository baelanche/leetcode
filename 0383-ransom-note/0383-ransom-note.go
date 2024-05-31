func canConstruct(ransomNote string, magazine string) bool {
    alphabets := [26]int{}
    for i, j := 0, 0; i<len(ransomNote) || j<len(magazine); i, j = i+1, j+1 {
        if i < len(ransomNote) {
            alphabets[ransomNote[i] - 'a']--
        }
        if j < len(magazine) {
            alphabets[magazine[j] - 'a']++
        }
    }

    for _, alphabet := range alphabets {
        if alphabet < 0 {
            return false
        }
    }
    return true
}