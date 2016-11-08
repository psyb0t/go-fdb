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

    index_to_key_path := c.Path + "/" + index_to_key_dir

    err = genutils.MkDirAll(index_to_key_path)

    if err != nil {
        return err
    }

    key_to_index_path := c.Path + "/" + key_to_index_dir

    err = genutils.MkDirAll(key_to_index_path)

    if err != nil {
        return err
    }

    return nil
}
