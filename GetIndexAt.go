package fdb

import (
    "io/ioutil"
    "strconv"
)

func (c *Collection) GetIndexAt() int {
    keynofpath := c.Path + "/" + keyno_file

    index, err := ioutil.ReadFile(keynofpath)

    if err != nil {
        return 0
    }

    indexint, err := strconv.Atoi(string(index))

    if err != nil {
        return 0
    }

    return indexint
}
