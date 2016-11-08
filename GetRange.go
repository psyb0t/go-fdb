package fdb

import (
)

func (c *Collection) GetRange(from int, to int) ([]string, error) {
    if len(c.Indexes) < to {
        to = len(c.Indexes)
    }

    range_indexes := c.Indexes[from:to]

    var values []string
    for _, key_index := range range_indexes {
        value, err := c.GetByIndex(key_index)

        if err != nil {
            continue
        }

        values = append(values, value)
    }

    return values, nil
}
