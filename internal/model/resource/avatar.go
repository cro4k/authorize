package resource

import (
	"database/sql/driver"
	"encoding/json"
)

type Avatar struct {
	resource
}

func (a *Avatar) Scan(v any) error {
	if v == nil {
		return nil
	}
	b := v.([]byte)
	if len(b) == 0 {
		return nil
	}
	err := json.Unmarshal(b, a)
	if err != nil {
		return err
	}
	a.fix()
	return nil
}

func (a Avatar) Value() (driver.Value, error) {
	a.trim()
	return json.Marshal(a)
}

func (a *Avatar) Simple() Simple {
	s := Simple{}
	s.Path = a.Path
	s.Thumbnail = a.Thumbnail
	s.Type = Image
	return s
}
