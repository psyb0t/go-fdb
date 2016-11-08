package fdb

import (
    "fmt"
)

func (c *Collection) Index2KeyPath(index int) string {
    return c.Path + "/" + index_to_key_dir +
        "/" + fmt.Sprintf("%06d", index)
}
