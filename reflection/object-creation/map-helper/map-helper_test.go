package maphelper

import (
	"object-creation/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertObjectToMap(t *testing.T) {
	usr := user.UserCreated{
		ID:   "123",
		Type: "UserCreated",
	}

	result := ToMap(&usr)
	assert.Equal(t, "123", result["ID"])
	assert.Equal(t, "UserCreated", result["Type"])
}
