package authorization

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	email := ""
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail2(t *testing.T) {
	email := "test.test@html.com"
	result := IsValidEmail(email)
	if result != true {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, true)
	}
}

func TestIsValidEmail3(t *testing.T) {
	email := "test.test@@@html.com"
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail4(t *testing.T) {
	email := "111111@html.com"
	result := IsValidEmail(email)
	if result != true {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, true)
	}
}

func TestIsValidEmail5(t *testing.T) {
	email := "user@.com"
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail6(t *testing.T) {
	email := "user@com"
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail7(t *testing.T) {
	email := "user@com."
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail8(t *testing.T) {
	email := "@com."
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail9(t *testing.T) {
	email := "user@example_com"
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail10(t *testing.T) {
	email := "user@example.c"
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidEmail11(t *testing.T) {
	email := "user@example.c"
	result := IsValidEmail(email)
	if result != false {
		t.Errorf("IsValidEmail(%s) = %t; want %t", email, result, false)
	}
}

func TestIsValidPassword(t *testing.T) {
	password := "1234567"
	result := IsValidPassword(password)
	if result != false {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, false)
	}
}

func TestIsValidPassword2(t *testing.T) {
	password := "12345678"
	result := IsValidPassword(password)
	if result != false {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, false)
	}
}

func TestIsValidPassword3(t *testing.T) {
	password := "12345678A"
	result := IsValidPassword(password)
	if result != true {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, true)
	}
}

func TestIsValidPassword4(t *testing.T) {
	password := "12345678Aa"
	result := IsValidPassword(password)
	if result != true {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, true)
	}
}

func TestIsValidPassword5(t *testing.T) {
	password := "12345678AA"
	result := IsValidPassword(password)
	if result != true {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, true)
	}
}

func TestIsValidPassword6(t *testing.T) {
	password := "12345678aa"
	result := IsValidPassword(password)
	if result != false {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, false)
	}
}

func TestIsValidPassword7(t *testing.T) {
	password := "AaAaAaAaAa"
	result := IsValidPassword(password)
	if result != false {
		t.Errorf("IsValidPassword(%s) = %t; want %t", password, result, false)
	}
}

func TestComparePasswords(t *testing.T) {
	password1 := "12345678Aa"
	hashedPassword1, _ := HashPassword(password1)

	password2 := "12345678Aa"
	result := ComparePasswords(hashedPassword1, password2)
	if result != true {
		t.Errorf("ComparePasswords(%s, %s) = %t; want %t", password1, password2, result, true)
	}
}
