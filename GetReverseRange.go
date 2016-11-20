package fdb

import "github.com/psyb0t/go-genutils"

func (c *Collection) GetReverseRange(from int, to int) ([]string, error) {
    if len(c.Keys) < to {
        to = len(c.Keys)
    }

    if from > to {
        from = to
    }

    reverse_keys := genutils.RevStrSlice(c.Keys)

    range_keys := reverse_keys[from:to]

    var values []string
    for _, key_name := range range_keys {
        value, err := c.KeyValue(key_name)

        if err != nil {
            continue
        }

        values = append(values, value)
    }

    return values, nil
}

func reverseSlice(input []string) []string {
    if len(input) == 0 {
        return input
    }
    return append(reverseSlice(input[1:]), input[0])
}
