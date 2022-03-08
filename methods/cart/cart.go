package cart

import (
	"errors"
	"methodsPractice/product"
	"os/user"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

type Cart struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt time.Time
	user.User
	Items []Item
	CurrencyCode string
	isLocked bool
}

type Item struct {
	product.Product
	Quantity uint8
}


// calculates the total price of the products in the cart
func (c *Cart)TotalPrice()(*money.Money, error) {
	totalAmount := money.New(0, "EUR")
	var err error
	for _, v := range c.Items {
	subtotal:=v.Product.Price.Multiply(int64(v.Quantity))
	totalAmount, err = totalAmount.Add(subtotal)
	if err != nil {
		return nil, err
	}
	}
	return totalAmount, nil
}

// lockes the cart to avoid further modifications after confirmation
func (c *Cart)Lock() error {
	
	if c.isLocked {
		return errors.New("ERROR: the cart is already locked")
	} else {
		c.isLocked = true;
		c.lockedAt = time.Now()
		return nil
	} 
	
}

func (c *Cart)delete() error {
	// to implement
	return nil
}

// unit tests for TotalPrice() and Lock() functions

func TestTotalPrice(t *testing.T) {
	items := []Item{
		{
		Product: product.Product {
			ID: "p-1254",
			Name: "Product Test",
			Price: money.New(1000, "EUR"),
		},
		Quantity: 2,
		},
		{
		Product: product.Product{
			ID: "p-1255",
			Name: "Product test 2",
			Price: money.New(2000,"EUR"),
		},
		Quantity: 1,	
		},
	}
	c:= Cart{
		ID: "1254",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User: user.User{},
		Items: items,
		CurrencyCode: "EUR",
	}
	actual, err := c.TotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, money.New(4000, "EUR"), actual)
}
func TestLock(t *testing.T) {
	c:= Cart {
		ID: "1254",
	}
	err := c.Lock()
	assert.NoError(t, err)
	assert.True(t, c.isLocked)
	assert.True(t, c.lockedAt.Unix() > 0)
}

func TestLockAlreadyLocked(t *testing.T){
	c:= Cart{
		ID: "1254",
		isLocked: true,
	}
	err := c.Lock()
	assert.Error(t,err)
}