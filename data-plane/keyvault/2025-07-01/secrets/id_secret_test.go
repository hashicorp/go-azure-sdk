package secrets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SecretId{}

func TestNewSecretID(t *testing.T) {
	id := NewSecretID("secretName")

	if id.SecretName != "secretName" {
		t.Fatalf("Expected %q but got %q for Segment 'SecretName'", id.SecretName, "secretName")
	}
}

func TestFormatSecretID(t *testing.T) {
	actual := NewSecretID("secretName").ID()
	expected := "/secrets/secretName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSecretID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecretId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/secrets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/secrets/secretName",
			Expected: &SecretId{
				SecretName: "secretName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/secrets/secretName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecretID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SecretName != v.Expected.SecretName {
			t.Fatalf("Expected %q but got %q for SecretName", v.Expected.SecretName, actual.SecretName)
		}

	}
}

func TestParseSecretIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecretId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/secrets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sEcReTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/secrets/secretName",
			Expected: &SecretId{
				SecretName: "secretName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/secrets/secretName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sEcReTs/sEcReTnAmE",
			Expected: &SecretId{
				SecretName: "sEcReTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sEcReTs/sEcReTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecretIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SecretName != v.Expected.SecretName {
			t.Fatalf("Expected %q but got %q for SecretName", v.Expected.SecretName, actual.SecretName)
		}

	}
}

func TestSegmentsForSecretId(t *testing.T) {
	segments := SecretId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SecretId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
