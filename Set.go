package fdb

import (
    "strings"
    "errors"
    "io/ioutil"

    "github.com/psyb0t/go-genutils"
)

func (c *Collection) Set(key string, value string) error {
    if genutils.StringInSlice(key, reserved_cnames) {
        return errors.New("Reserved key name error")
    }

    if c.KeyExists(key) {
        return errors.New("Key already exists")
    }

    if strings.Index(key, "~") != -1 {
        return errors.New("Invalid character in strings: ~")
    }

    comp_val, err := genutils.Compress([]byte(value))

    if err != nil {
        return err
    }

    err = ioutil.WriteFile(c.KeyPath(key), comp_val, 0644)

    if err != nil {
        return err
    }

    c.Keys = append(c.Keys, key)

    err = c.UpdateIndexAt(c.IndexAt + 1)

    if err != nil {
        return err
    }

    return nil
}
