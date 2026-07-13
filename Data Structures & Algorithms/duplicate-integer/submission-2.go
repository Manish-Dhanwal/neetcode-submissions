func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, k := range nums{
        _, ok := m[k]
        if ok{
            return true
        }
        m[k] = true
    }
    return false
}
