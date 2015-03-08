package api

import (
	"encoding/json"
	"strings"
	"time"
)

const Fmt = "2006-01-02T15:04:05"

type jTime time.Time

func (jt *jTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(Fmt, s)
	if err != nil {
		// find last "."
		i := strings.LastIndex(s, ".")
		if i == -1 {
			return err
		}

		// try to parse the slice of the string
		t, err = time.Parse(Fmt, s[0:i])
		if err != nil {
			return err
		}
	}
	*jt = (jTime)(t)
	return nil
}

func (jt jTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&jt).Format(Fmt))
}
