package fdb

func (c *Collection) KeyInMem(key string, keyfile string) {
    c.Keys = append(c.Keys, key)
    c.KeyFiles = append(c.KeyFiles, keyfile)

    c.Key2File[key] = keyfile
    c.File2Key[keyfile] = key
}
