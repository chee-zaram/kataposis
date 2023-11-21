package config

import (
	"os"
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv(pgUser, "test_user")
	os.Setenv(pgPass, "test_pass")
	os.Setenv(pgAddr, "test_addr")
	os.Setenv(pgDBName, "test_db")

	defer unsetEnv()

	config, err := NewConfig()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify the values are set correctly
	if config.PGUser != "test_user" {
		t.Errorf("Expected PGUser to be 'test_user', got '%s'", config.PGUser)
	}

	if config.PGPass != "test_pass" {
		t.Errorf("Expected PGPass to be 'test_pass', got '%s'", config.PGPass)
	}

	if config.PGAddr != "test_addr" {
		t.Errorf("Expected PGAddr to be 'test_addr', got '%s'", config.PGAddr)
	}

	if config.PGDBName != "test_db" {
		t.Errorf("Expected PGDB to be 'test_db', got '%s'", config.PGDBName)
	}
}

func TestNewConfig_WithoutEnv(t *testing.T) {
	// Set all variables and then unset them one after another to test in the
	// for loop below.
	os.Setenv(pgUser, "test_user")
	os.Setenv(pgPass, "test_pass")
	os.Setenv(pgAddr, "test_addr")
	os.Setenv(pgDBName, "test_db")

	defer unsetEnv()

	values := map[string]string{
		pgUser:   "user",
		pgPass:   "pass",
		pgAddr:   "addr",
		pgDBName: "name",
	}

	for k, v := range values {
		os.Unsetenv(k)
		val, err := NewConfig()
		if err == nil {
			t.Fatalf("Expected %s to be unset, got '%s'", k, val)
		}

		if !strings.Contains(err.Error(), v) {
			t.Fatalf("Expected error to contain '%s', got '%s'", v, err.Error())
		}

		os.Setenv(k, "placeholder")
	}
}

func unsetEnv() {
	os.Unsetenv(pgUser)
	os.Unsetenv(pgPass)
	os.Unsetenv(pgAddr)
	os.Unsetenv(pgDBName)
}
