package fdb

import (
    "errors"
    "strconv"
    "io/ioutil"
)

func (c *Collection) IndexByKey(key string) (int, error) {
    if !c.KeyExists(key) {
        return 0, errors.New("Key does not exist")
    }

    val, err := ioutil.ReadFile(c.Key2IndexPath(key))

    if err != nil {
        return 0, err
    }

    index, err := strconv.Atoi(string(val))

    if err != nil {
        return 0, err
    }

    return index, nil
}
