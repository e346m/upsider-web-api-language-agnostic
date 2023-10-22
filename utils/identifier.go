package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type Identifier struct {
	defaultEntropySource *ulid.MonotonicEntropy
}

func NewIdentifier() *Identifier {
	return &Identifier{}
}

func (i *Identifier) NewIdString() string {
	t := time.Now()
	r := rand.New(rand.NewSource(t.UnixNano()))
	i.defaultEntropySource = ulid.Monotonic(r, 0)
	u := ulid.MustNew(ulid.Timestamp(t), i.defaultEntropySource)
	return u.String()
}

func (i *Identifier) StringToBinary(s string) ([]byte, error) {
	u, err := ulid.Parse(s)
	if err != nil {
		return nil, err
	}

	b, err := u.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return b, err
}

func (i *Identifier) BinaryToString(b []byte) (string, error) {
	var ret [16]byte
	copy(ret[:], b)
	s := ulid.ULID(ret).String()

	if _, err := ulid.Parse(s); err != nil {
		return "", err
	}

	return s, nil
}
