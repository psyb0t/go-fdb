package fdb

func (c *Collection) KeyExists(key string) bool {
    if _, exists := c.Key2File[key]; exists {
        return true
    }

    return false
}
