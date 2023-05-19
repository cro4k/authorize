package pathutil

import "strings"

type Option func(o *URLPattern)

func applyOptions(options ...Option) *URLPattern {
	o := &URLPattern{
		regLeft:  "r(",
		regRight: ")",
	}
	for _, opt := range options {
		opt(o)
	}
	return o
}

func WithLeftRegPattern(pattern string) Option {
	return func(o *URLPattern) {
		if pattern != "" {
			o.regLeft = pattern
		}
	}
}

func WithRightRegPattern(pattern string) Option {
	return func(o *URLPattern) {
		if pattern != "" {
			o.regRight = pattern
		}
	}
}

type URLPattern struct {
	origin string
	split  []Matcher

	regLeft  string
	regRight string
}

func NewURLPattern(p string, options ...Option) (*URLPattern, error) {
	pattern := applyOptions(options...)
	ln := len(pattern.regLeft)
	rn := len(pattern.regRight)

	temp := strings.Split(p, "/")
	pattern.split = make([]Matcher, 0, len(temp))

	for _, v := range temp {
		if len(v) > ln+rn && v[:ln] == pattern.regLeft && v[len(v)-rn:] == pattern.regRight {
			matcher, err := NewRegexpMatcher(v[ln : len(v)-rn])
			if err != nil {
				return nil, err
			}
			pattern.split = append(pattern.split, matcher)
		} else {
			pattern.split = append(pattern.split, TextMatcher(v))
		}
	}
	return pattern, nil
}

func (p *URLPattern) Match(path string) bool {
	if path == "" {
		return false
	}
	temp := strings.Split(path, "/")
	if len(p.split) != len(temp) {
		return false
	}
	for i, v := range p.split {
		if !v.Match(temp[i]) {
			return false
		}
	}
	return true
}
