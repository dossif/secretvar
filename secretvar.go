package secretvar

import "encoding/json"

// SecretPlaceHolder default placeholder for secret string
var SecretPlaceHolder = "*****"

type SecretString string

// Get returns UnMasked string value!
func (s SecretString) Get() string {
	return string(s)
}

// String returns masked string value
func (s SecretString) String() string {
	return SecretPlaceHolder
}

// MarshalJSON returns masked string value
func (s SecretString) MarshalJSON() ([]byte, error) {
	return json.Marshal(SecretPlaceHolder)
}

// MarshalYAML returns masked string value
func (s SecretString) MarshalYAML() (interface{}, error) {
	return SecretPlaceHolder, nil
}
