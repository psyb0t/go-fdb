package fdb

func (c *Collection) KeyValuePath(key string) string {
    return c.Path + "/" + key + key_val_file_ext
}
