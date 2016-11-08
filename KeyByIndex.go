package fdb

import (
    "errors"
    "io/ioutil"
)

func (c *Collection) KeyByIndex(index int) (string, error) {
    if !c.IndexExists(index) {
        return "", errors.New("Index does not exist")
    }

    key, err := ioutil.ReadFile(c.Index2KeyPath(index))

    if err != nil {
        return "", err
    }

    return string(key), nil
}
