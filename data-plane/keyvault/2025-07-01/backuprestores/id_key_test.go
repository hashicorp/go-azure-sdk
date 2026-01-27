package backuprestores

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &KeyId{}

func TestNewKeyID(t *testing.T) {
	id := NewKeyID("keyName")

	if id.KeyName != "keyName" {
		t.Fatalf("Expected %q but got %q for Segment 'KeyName'", id.KeyName, "keyName")
	}
}

func TestFormatKeyID(t *testing.T) {
	actual := NewKeyID("keyName").ID()
	expected := "/keys/keyName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseKeyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KeyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/keys",
			Error: true,
		},
		{
			// Valid URI
			Input: "/keys/keyName",
			Expected: &KeyId{
				KeyName: "keyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/keys/keyName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKeyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.KeyName != v.Expected.KeyName {
			t.Fatalf("Expected %q but got %q for KeyName", v.Expected.KeyName, actual.KeyName)
		}

	}
}

func TestParseKeyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *KeyId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/keys",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/kEyS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/keys/keyName",
			Expected: &KeyId{
				KeyName: "keyName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/keys/keyName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/kEyS/kEyNaMe",
			Expected: &KeyId{
				KeyName: "kEyNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/kEyS/kEyNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseKeyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.KeyName != v.Expected.KeyName {
			t.Fatalf("Expected %q but got %q for KeyName", v.Expected.KeyName, actual.KeyName)
		}

	}
}

func TestSegmentsForKeyId(t *testing.T) {
	segments := KeyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("KeyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
