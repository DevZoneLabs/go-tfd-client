package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type BaseStat struct {
	ID    string  `json:"stat_id"`
	Value float64 `json:"stat_value"`
}

func (b *BaseStat) UnmarshalJSON(data []byte) error {
	type Alias BaseStat
	aux := &struct {
		Value interface{} `json:"stat_value"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.Value.(type) {
	case string:
		parsedValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		b.Value = parsedValue
	case float64:
		b.Value = v
	default:
		return fmt.Errorf("unexpected type for stat_value: %T", v)
	}

	return nil
}
