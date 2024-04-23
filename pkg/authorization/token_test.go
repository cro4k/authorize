package authorization

import "testing"

func TestToken(t *testing.T) {
	token := NewToken("123", WithClientID("aaa"), WithClaims(map[string]interface{}{"aaa": "bbb"}))
	text, err := token.Encode("hello")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(text)
	var claims = map[string]any{}
	t2, err := ParseToken(text, WithClaims(&claims))
	if err != nil {
		t.Error(err)
	}
	if err := t2.Validate("hello"); err != nil {
		t.Error(err)
	}
}
