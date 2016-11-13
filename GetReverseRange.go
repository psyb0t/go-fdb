package fdb

func (c *Collection) GetReverseRange(from int, to int) ([]string, error) {
    if len(c.Keys) < to {
        to = len(c.Keys)
    }

    if from > to {
        from = to
    }

    reverse_keys := []*string{}

    for i:=len(c.Keys)-1; i>=0; i-- {
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
