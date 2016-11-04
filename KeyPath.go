package fdb

import (
    "strings"
)

func (c *Collection) KeyPath(key string) string {
    return c.Path + "/" + strings.Replace(key, "/", "_", -1) + keyfile_ext
}
