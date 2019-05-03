// +build !race

package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordSuccess(t *testing.T) {
	password := CreatePassword("secret", 5)
	assert.Equal(t, true, ComparePassword(password, []byte("secret")))
}

func TestPasswordFailure(t *testing.T) {
	password := CreatePassword("secret", 5)
	assert.Equal(t, false, ComparePassword(password, []byte("secretx")))
}

func TestBCryptFailure(t *testing.T) {
	assert.Panics(t, func() { CreatePassword("secret", 12312) })
}
