package fdb

func (c *Collection) KeyPath(key string) string {
    return c.Path + "/" + c.Key2File[key]
}
