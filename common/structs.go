package common

import "gorm.io/gorm"

type Card struct {
    gorm.Model
    Name       string `json:"name"`
    NameShort  string `json:"name_short"`
    Type       string `json:"type"`
    Value      string `json:"value"`
    ValueInt   int    `json:"value_int"`
    MeaningUp  string `json:"meaning_up"`
    MeaningRev string `json:"meaning_rev"`
    Description string `json:"desc"`
    Image      string `json:"image"`
}
