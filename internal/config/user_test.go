// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/modern/sqlconfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddUser(t *testing.T) {
	defer func() { assert.NotNil(t, recover(), "The code did not panic as expected") }()
	AddUser(User{
		Name:               "",
		AuthenticationType: "basic",
		BasicAuth:          nil,
	})
}

func TestNegAddUser(t *testing.T) {
	assert.Panics(t, func() {
		AddUser(User{
			Name:               "",
			AuthenticationType: "basic",
			BasicAuth: &BasicAuthDetails{
				Username:          "",
				PasswordEncrypted: false,
				Password:          "",
			},
		})
	})
}

func TestNegAddUser2(t *testing.T) {
	assert.Panics(t, func() {
		GetUser("doesnotexist")
	})
}
