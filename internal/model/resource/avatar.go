package resource

import (
	"database/sql/driver"
	"encoding/json"
)

type Avatar struct {
	resource
}

func (a *Avatar) Scan(v any) error {
	b := v.([]byte)
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
