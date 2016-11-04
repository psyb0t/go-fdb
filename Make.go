package fdb

import (
    "os"
    "strconv"
    "io/ioutil"
)

func (c *Collection) Make() error {
    _, err := os.Stat(c.Path)

    if os.IsNotExist(err) {
        err := os.MkdirAll(c.Path, 0644)

        if err != nil {
            return err
        }

        return nil
    }

    if err != nil {
        return err
    }

    keynofpath := c.Path + "/" + keyno_file

    _, err = os.Stat(keynofpath)

    if os.IsNotExist(err) {
        err = ioutil.WriteFile(
            keynofpath, []byte(strconv.Itoa(c.IndexAt)), 0644)

        if err != nil {
            return err
        }
    }

    if err != nil {
        return err
    }

    return nil
}
