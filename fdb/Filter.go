package fdb

import (
    "regexp"
)

func (c *Collection) Filter(list []string, regex string) []string {
    var filter_result []string

    re, err := regexp.Compile(regex)

    if err != nil {
        return filter_result
    }

    for i:=0; i<len(list); i++ {
        if re.MatchString(list[i]) {
            filter_result = append(filter_result, list[i])
        }
    }

    return filter_result
}
