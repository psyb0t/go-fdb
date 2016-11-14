package fdb

import (
    "errors"
    "io/ioutil"

    "github.com/psyb0t/go-genutils"
)

func (c *Collection) KeyValue(key string) (string, error) {
    if !c.KeyExists(key) {
        return "", errors.New("Key does not exist")
    }

    value, err := ioutil.ReadFile(c.KeyPath(key))

    if err != nil {
        return "caca", err
    }

    decomp_val, err := genutils.Decompress(value)

    if err != nil {
        return "", err
    }

    return string(decomp_val), nil
}
