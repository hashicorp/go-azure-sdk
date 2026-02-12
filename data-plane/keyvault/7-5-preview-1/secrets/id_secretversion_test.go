package secrets

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SecretversionId{}

func TestNewSecretversionID(t *testing.T) {
	id := NewSecretversionID("https://endpoint-url.example.com", "secretName", "secretversion")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.SecretName != "secretName" {
		t.Fatalf("Expected %q but got %q for Segment 'SecretName'", id.SecretName, "secretName")
	}

	if id.Secretversion != "secretversion" {
		t.Fatalf("Expected %q but got %q for Segment 'Secretversion'", id.Secretversion, "secretversion")
	}
}

func TestFormatSecretversionID(t *testing.T) {
	actual := NewSecretversionID("https://endpoint-url.example.com", "secretName", "secretversion").ID()
	expected := "https://endpoint-url.example.com/secrets/secretName/secretversion"
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
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/secrets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/secrets/secretName",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/secrets/secretName/secretversion",
			Expected: &SecretversionId{
				BaseURI:       "https://endpoint-url.example.com",
				SecretName:    "secretName",
				Secretversion: "secretversion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/secrets/secretName/secretversion/extra",
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

		if actual.Secretversion != v.Expected.Secretversion {
			t.Fatalf("Expected %q but got %q for Secretversion", v.Expected.Secretversion, actual.Secretversion)
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
			Input: "https://endpoint-url.example.com/secrets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEcReTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/secrets/secretName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEcReTs/sEcReTnAmE",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/secrets/secretName/secretversion",
			Expected: &SecretversionId{
				BaseURI:       "https://endpoint-url.example.com",
				SecretName:    "secretName",
				Secretversion: "secretversion",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/secrets/secretName/secretversion/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEcReTs/sEcReTnAmE/sEcReTvErSiOn",
			Expected: &SecretversionId{
				BaseURI:       "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				SecretName:    "sEcReTnAmE",
				Secretversion: "sEcReTvErSiOn",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/sEcReTs/sEcReTnAmE/sEcReTvErSiOn/extra",
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

		if actual.Secretversion != v.Expected.Secretversion {
			t.Fatalf("Expected %q but got %q for Secretversion", v.Expected.Secretversion, actual.Secretversion)
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
