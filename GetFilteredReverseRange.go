package fdb

import "regexp"

func (c *Collection) GetFilteredReverseRange(
  from int, to int, regex string) ([]string, error) {
    if len(c.Keys) < to {
        to = len(c.Keys)
    }

    if from > to {
        from = to
    }

    re, err := regexp.Compile(regex)

    if err != nil {
        return []string{}, err
    }

    reverse_keys := []*string{}

    for i:=len(c.Keys)-1; i>=0; i-- {
        if !re.MatchString(c.Keys[i]) {
            continue
        }

        reverse_keys = append(reverse_keys, &c.Keys[i])
    }

    range_keys := reverse_keys[from:to]

    var values []string
    for _, key_name := range range_keys {
        value, err := c.KeyValue(*key_name)

        if err != nil {
            continue
        }

        values = append(values, value)
    }

    return values, nil
}
