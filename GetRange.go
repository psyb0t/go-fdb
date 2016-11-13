package fdb

func (c *Collection) GetRange(from int, to int) ([]string, error) {
    if len(c.Keys) < to {
        to = len(c.Keys)
    }

    if from > to {
        from = to
    }

    range_keys := c.Keys[from:to]

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
