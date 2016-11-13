package fdb

func (c *Collection) Get(key string) (string, error) {
    keyval, err := c.KeyValue(key)

    if err != nil {
        return "", err
    }

    return keyval, nil
}
