package locks

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LockId{}

func TestNewLockID(t *testing.T) {
	id := NewLockID("https://endpoint-url.example.com", "lockName")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.LockName != "lockName" {
		t.Fatalf("Expected %q but got %q for Segment 'LockName'", id.LockName, "lockName")
	}
}

func TestFormatLockID(t *testing.T) {
	actual := NewLockID("https://endpoint-url.example.com", "lockName").ID()
	expected := "https://endpoint-url.example.com/locks/lockName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLockID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LockId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/locks",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/locks/lockName",
			Expected: &LockId{
				BaseURI:  "https://endpoint-url.example.com",
				LockName: "lockName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/locks/lockName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLockID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.LockName != v.Expected.LockName {
			t.Fatalf("Expected %q but got %q for LockName", v.Expected.LockName, actual.LockName)
		}

	}
}

func TestParseLockIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LockId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/locks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lOcKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/locks/lockName",
			Expected: &LockId{
				BaseURI:  "https://endpoint-url.example.com",
				LockName: "lockName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/locks/lockName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lOcKs/lOcKnAmE",
			Expected: &LockId{
				BaseURI:  "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				LockName: "lOcKnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/lOcKs/lOcKnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLockIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.LockName != v.Expected.LockName {
			t.Fatalf("Expected %q but got %q for LockName", v.Expected.LockName, actual.LockName)
		}

	}
}

func TestSegmentsForLockId(t *testing.T) {
	segments := LockId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LockId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
