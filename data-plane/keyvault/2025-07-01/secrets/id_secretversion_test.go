package secrets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SecretversionId{}

func TestNewSecretversionID(t *testing.T) {
	id := NewSecretversionID("https://endpoint_url", "secretName")

	if id.BaseURI != "https://endpoint_url" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint_url")
	}

	if id.SecretName != "secretName" {
		t.Fatalf("Expected %q but got %q for Segment 'SecretName'", id.SecretName, "secretName")
	}
}

func TestFormatSecretversionID(t *testing.T) {
	actual := NewSecretversionID("https://endpoint_url", "secretName").ID()
	expected := "/https://endpoint_url/secrets/secretName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSecretversionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecretversionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/secrets",
			Error: true,
		},
		{
			// Valid URI
			Input: "/https://endpoint_url/secrets/secretName",
			Expected: &SecretversionId{
				BaseURI:    "https://endpoint_url",
				SecretName: "secretName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/secrets/secretName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecretversionID(v.Input)
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

		if actual.SecretName != v.Expected.SecretName {
			t.Fatalf("Expected %q but got %q for SecretName", v.Expected.SecretName, actual.SecretName)
		}

	}
}

func TestParseSecretversionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SecretversionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/https://endpoint_url/secrets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/sEcReTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/https://endpoint_url/secrets/secretName",
			Expected: &SecretversionId{
				BaseURI:    "https://endpoint_url",
				SecretName: "secretName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/https://endpoint_url/secrets/secretName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/sEcReTs/sEcReTnAmE",
			Expected: &SecretversionId{
				BaseURI:    "hTtPs://eNdPoInT_UrL",
				SecretName: "sEcReTnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/hTtPs://eNdPoInT_UrL/sEcReTs/sEcReTnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSecretversionIDInsensitively(v.Input)
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

		if actual.SecretName != v.Expected.SecretName {
			t.Fatalf("Expected %q but got %q for SecretName", v.Expected.SecretName, actual.SecretName)
		}

	}
}

func TestSegmentsForSecretversionId(t *testing.T) {
	segments := SecretversionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SecretversionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
