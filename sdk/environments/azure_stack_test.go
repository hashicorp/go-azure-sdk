// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package environments

import (
	"log"
	"testing"
)

func TestIsAzureStack_BuiltInEnvironmentsShouldNotBeMarkedAsStack(t *testing.T) {
	testData := []*Environment{
		AzurePublic(),
		AzurePublicCanary(),
		AzureChina(),
		AzureUSGovernment(),
		AzureUSGovernmentL5(),
	}
	for _, v := range testData {
		log.Printf("Environment: %q", v.Name)
		if v.IsAzureStack() {
			t.Fatalf("expected IsAzureStack to be false but was true for Environment %q", v.Name)
		}
	}
}

func TestIsAzureStack(t *testing.T) {
	testData := []*Environment{
		{
			Name: "AzureStackCloud",
		},
		{
			Name: "Identity Provider Empty",
			Authorization: &Authorization{
				IdentityProvider: "",
			},
		},
		{
			Name: "Identity Provider Custom",
			Authorization: &Authorization{
				IdentityProvider: "Custom",
			},
		},
		{
			Name: "Tenant Empty",
			Authorization: &Authorization{
				IdentityProvider: "AAD",
				Tenant:           "",
			},
		},
		{
			Name: "Tenant Custom",
			Authorization: &Authorization{
				IdentityProvider: "AAD",
				Tenant:           "Custom",
			},
		},
	}
	for _, v := range testData {
		log.Printf("Environment %q", v.Name)
		if !v.IsAzureStack() {
			t.Fatalf("expected IsAzureStack to be true but was false for Environment %q", v.Name)
		}
	}
}
