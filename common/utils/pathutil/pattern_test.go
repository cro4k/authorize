package pathutil

import "testing"

func TestPattern(t *testing.T) {
	p, err := NewURLPattern("/api/*/r(\\d)")
	if err != nil {
		t.Error(err)
		return
	}
	if p.Match("/api/user/") {
		t.Fail()
	}
	if p.Match("/api/user/") {
		t.Fail()
	}
	if p.Match("/api/user/abc") {
		t.Fail()
	}
	if !p.Match("/api/user/1") {
		t.Fail()
	}
	if !p.Match("/api/book/123") {
		t.Fail()
	}
	if p.Match("/api/user/123/456") {
		t.Fail()
	}
}
