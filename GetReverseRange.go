package fdb

func (c *Collection) GetReverseRange(from int, to int) ([]string, error) {
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
        value, err := c.GetByIndex(key_index)

        if err != nil {
            continue
        }

        values = append(values, value)
    }

    return values, nil
}
