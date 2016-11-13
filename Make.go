package fdb

import (
    "os"

    "github.com/psyb0t/go-genutils"
)

func (c *Collection) Make() error {
    _, err := os.Stat(c.Path)

    err = genutils.MkDirAll(c.Path)

    if err != nil {
        return err
    }

    return nil
}
