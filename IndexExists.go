package fdb

import (
    "os"
)

func (c *Collection) IndexExists(index int) bool {
    _, err := os.Stat(c.Index2KeyPath(index))

    if os.IsNotExist(err) {
        return false
    }

    if err != nil {
        panic(err)
    }

    return true
}
