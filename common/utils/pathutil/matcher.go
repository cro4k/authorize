package pathutil

import "regexp"

type Matcher interface {
	Match(s string) bool
}

type TextMatcher string

func (m TextMatcher) Match(s string) bool {
	return string(m) == "*" || string(m) == s
}

type RegexpMatcher struct {
	origin string
	reg    *regexp.Regexp
}

func NewRegexpMatcher(exp string) (*RegexpMatcher, error) {
	m := &RegexpMatcher{origin: exp}
	var err error
	m.reg, err = regexp.Compile(exp)
	return m, err
}

func (m *RegexpMatcher) Match(s string) bool {
	return m.reg.MatchString(s)
}

//type MatchDecoder interface {
//	Decode(v string) (Matcher, error)
//}
//
//type RegexpDecoder struct{}
//
//func (r *RegexpDecoder) Decode(v string) (Matcher, error) {
//
//}
