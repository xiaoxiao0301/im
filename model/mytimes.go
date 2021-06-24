package model

import "time"

type Mytimes time.Time

func (this Mytimes) MarshalJSON() ([]byte, error) {
	if time.Time(this).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(this).Format("2006-01-02 15:04:05") + `"`), nil
}
