package types

import "encoding/json"

type Secret string

func (s Secret) String() string {
	return string(s)
}

func (s Secret) Masking() string {
	return "*****"
}

func (s Secret) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Masking())
}
