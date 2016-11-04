package fdb

import (
    "strconv"
    "io/ioutil"
)

func (c *Collection) UpdateIndexAt(index int) error {
    keynofpath := c.Path + "/" + keyno_file

    err := ioutil.WriteFile(
        keynofpath, []byte(strconv.Itoa(index)), 0644)

    if err != nil {
        return err
    }

    c.IndexAt = index

    return nil
}
