package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_Gets_An_Error_If_Id_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.ValidateFields(), "invalid id")
}

func Test_If_Gets_An_Error_If_Price_Small_Or_Equal_Than_Zero(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.ValidateFields(), "invalid price")
}

func Test_If_Gets_An_Error_If_Tax_Small_Or_Equal_Than_Zero(t *testing.T) {
	order := Order{ID: "123", Price: 34.0}
	assert.Error(t, order.ValidateFields(), "invalid tax")
}

func Test_Sucess(t *testing.T) {
	order := Order{ID: "123", Price: 34.0, Tax: 10.2}
	assert.NoError(t, order.ValidateFields())
	assert.Equal(t, order.Price, 34.0)
}
