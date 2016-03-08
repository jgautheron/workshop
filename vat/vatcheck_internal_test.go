package vatcheck

import (
	"errors"
	"io/ioutil"
	"log"
	"testing"

	"github.com/mattes/vat"
)

var (
	errMock     = errors.New("mock error to trip the circuit breaker")
	validNumber = "PL6492291353"
)

func init() {
	// Disable logs
	log.SetOutput(ioutil.Discard)
}

func mockCheckVAT(id string) (*vat.VATresponse, error) {
	return &vat.VATresponse{Valid: true}, nil
}

func TestFormatInvalid(t *testing.T) {
	if _, err := IsValid("fooo"); err != ErrInvalidFormat {
		t.Error("Format should be invalid")
	}
}

func TestIsValid(t *testing.T) {
	c.checkMethod = func(id string) (*vat.VATresponse, error) {
		return &vat.VATresponse{Valid: true}, nil
	}

	if isValid, _ := IsValid(validNumber); !isValid {
		t.Error("The number %s should be valid", validNumber)
	}
}

func TestInvalid(t *testing.T) {
	c.checkMethod = func(id string) (*vat.VATresponse, error) {
		return nil, errMock
	}

	if _, err := IsValid(validNumber); err == nil {
		t.Error("An error should be returned")
	}
}

func TestCircuitTripped(t *testing.T) {
	// Re-initialise the circuit breaker
	initCircuitBreaker()

	// How many failures before the circuit trips
	countToTrip := 2

	c.checkMethod = func(id string) (*vat.VATresponse, error) {
		return nil, errMock
	}

	// Trip the circuit
	for i := 0; i < countToTrip; i++ {
		if _, err := IsValid(validNumber); err != errMock {
			t.Error("The wrong error has been returned")
		}
	}

	if _, err := IsValid(validNumber); err != ErrCircuitTripped {
		t.Error("The wrong error has been returned")
	}
}
