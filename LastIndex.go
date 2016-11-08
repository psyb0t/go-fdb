package fdb

import (
    "os"
    "io/ioutil"
    "strconv"
)

func (c *Collection) LastIndex() (int, error) {
    last_index_fpath := c.Path + "/" + last_index_file

    _, err := os.Stat(last_index_fpath)

    if os.IsNotExist(err) {
        err = ioutil.WriteFile(last_index_fpath, []byte("0"), 0644)

        if err != nil {
            return 0, err
        }

        return 0, nil
    }

    if err != nil {
        return 0, err
    }

    value, err := ioutil.ReadFile(last_index_fpath)

    if err != nil {
        return 0, err
    }

    index, err := strconv.Atoi(string(value))

    if err != nil {
        return 0, err
    }

    return index, nil
}
