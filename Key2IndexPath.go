package fdb

func (c *Collection) Key2IndexPath(key string) string {
    return c.Path + "/" + key_to_index_dir +
        "/" + key
}
