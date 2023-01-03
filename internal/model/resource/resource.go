package resource

import (
	"database/sql/driver"
	"encoding/json"
)

type Type int8

const (
	Others Type = iota
	Image
	Audio
	Video
	Document
)

type resource struct {
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
	Type      Type   `json:"type"`
}

func (r *resource) trim() {
	//r.Path = fsutil.Path(r.Path)
	//r.Thumbnail = fsutil.Path(r.Thumbnail)
}

func (r *resource) fix() {
	//r.Path = fsutil.URL(r.Path)
	//r.Thumbnail = fsutil.URL(r.Thumbnail)
}

type Resource struct {
	resource
	Name   string `json:"name"`
	MD5    string `json:"md5"`
	Ext    string `json:"ext"`
	Size   int    `json:"size"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Simple struct {
	resource
}

func (a *Resource) Scan(v any) error {
	b := v.([]byte)
	err := json.Unmarshal(b, a)
	if err != nil {
		return err
	}
	a.fix()
	return nil
}

func (a *Resource) Simple() Simple {
	s := Simple{}
	s.Path = a.Path
	s.Thumbnail = a.Thumbnail
	s.Type = a.Type
	return s
}

func (a Resource) Value() (driver.Value, error) {
	a.trim()
	return json.Marshal(a)
}

type Resources []Resource

func (a *Resources) Scan(v any) error {
	b := v.([]byte)
	err := json.Unmarshal(b, a)
	if err != nil {
		return err
	}
	for i := range *a {
		((*a)[i]).fix()
	}
	return nil
}

func (a Resources) Value() (driver.Value, error) {
	for i := range a {
		a[i].trim()
	}
	return json.Marshal(a)
}

func (a Resources) Simple() []Simple {
	var res = make([]Simple, 0, len(a))
	for _, v := range a {
		res = append(res, v.Simple())
	}
	return res
}
