package fdb

import (
    "errors"
    "io/ioutil"

    "github.com/psyb0t/go-genutils"
)

func (c *Collection) Update(key string, value string) error {
    if !c.KeyExists(key) {
        return errors.New("Key does not exist")
    }

    comp_val, err := genutils.Compress([]byte(value))

    if err != nil {
        return err
    }

    err = ioutil.WriteFile(c.KeyValuePath(key), comp_val, 0644)

    if err != nil {
        return err
    }

    return nil
}
