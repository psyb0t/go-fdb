package fdb

import (
    "io/ioutil"
    "time"
    "strconv"

    "github.com/psyb0t/go-genutils"
)

func (c *Collection) AddKey(key string, value string) error {
    comp_val, err := genutils.Compress([]byte(value))

    if err != nil {
        return err
    }

    keyts := time.Now().UnixNano()
    keyfile := c.Path + "/" + strconv.FormatInt(keyts, 10) + "~" + key + key_val_file_ext

    err = ioutil.WriteFile(keyfile, comp_val, 0644)

    if err != nil {
        return err
    }

    c.KeyInMem(key, keyfile)

    return nil
}
