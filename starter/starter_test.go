package starter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSayHello(t *testing.T){
	greeting := SayHello("Anupam")
	assert.Equal(t, "Hello Anupam. Welcome!", greeting)
}

func TestOddOrEven(t *testing.T) {
	val := OddOrEven(45)
	assert.Equal(t, 1,val)
}