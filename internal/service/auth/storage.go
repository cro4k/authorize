package auth

import (
	"errors"
	"github.com/cro4k/common/randx"
	"sync"
	"time"
)

const (
	maxClientAllowed = 5
)

type nonceData struct {
	cid    string
	nonce  string
	expire int64
}

type tokenNonce struct {
	uid   string
	nonce []*nonceData
}

type Storage interface {
	Get(uid string, cid string) (string, error)
	Gen(uid string, cid string) (string, error)
	Del(uid string, cid ...string) error
}

type storage struct {
	mu    *sync.Mutex
	cache map[string]*tokenNonce
}

func newStorage() *storage {
	return &storage{
		mu:    new(sync.Mutex),
		cache: make(map[string]*tokenNonce),
	}
}

func (s *storage) Get(uid, cid string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var n *tokenNonce
	if n = s.cache[uid]; n == nil {
		return "", errors.New("not found")
	}
	for _, v := range n.nonce {
		if v.cid == cid {
			return v.nonce, nil
		}
	}
	return "", errors.New("not found")
}

func (s *storage) Gen(uid, cid string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var n *tokenNonce
	if n = s.cache[uid]; n == nil {
		n = new(tokenNonce)
	}
	var no string
	var contain bool
	for _, v := range n.nonce {
		if v.cid == cid {
			v.nonce = randx.String(16)
			v.expire = time.Now().Unix() + 86400*7
			no = v.nonce
			contain = true
			break
		}
	}
	if !contain {
		if len(n.nonce) >= maxClientAllowed {
			n.nonce = n.nonce[1:]
		}
		no = randx.String(16)
		n.nonce = append(n.nonce, &nonceData{
			cid:    cid,
			nonce:  no,
			expire: time.Now().Unix() + 86400*7,
		})
	}
	s.cache[uid] = n
	return no, nil
}

func (s *storage) Del(uid string, cid ...string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(cid) == 0 {
		delete(s.cache, uid)
		return nil
	}
	n := s.cache[uid]
	if n == nil {
		return errors.New("not found")
	}
	for i, v := range n.nonce {
		if in(v.cid, cid) {
			if i == 0 {
				n.nonce = n.nonce[1:]
			} else if i == len(n.nonce)-1 {
				n.nonce = n.nonce[:len(n.nonce)-1]
			} else {
				n.nonce = append(n.nonce[:i], n.nonce[i+1:]...)
			}
			break
		}
	}
	return nil
}

func in(id string, ids []string) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
