package fdb

import "regexp"

func (c *Collection) GetFilteredReverseRange(
  from int, to int, regex string) ([]string, error) {
    if len(c.Indexes) < to {
        to = len(c.Indexes)
    }

    if from > to {
        from = to
    }

    reverse_indexes := []int{}

    for i:=len(c.Indexes)-1; i>=0; i-- {
        reverse_indexes = append(reverse_indexes, c.Indexes[i])
    }

    range_indexes := reverse_indexes[from:to]

    var values []string
    for _, key_index := range range_indexes {
        key_name, err := c.KeyByIndex(key_index)

        if err != nil {
            continue
        }

        re, err := regexp.Compile(regex)

        if err != nil {
            return values, err
        }

        if !re.MatchString(key_name) {
            continue
        }

        value, err := c.GetByIndex(key_index)

        if err != nil {
            continue
        }

        values = append(values, value)
    }

    return values, nil
}
