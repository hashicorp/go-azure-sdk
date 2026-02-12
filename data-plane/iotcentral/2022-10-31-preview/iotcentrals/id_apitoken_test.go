package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApiTokenId{}

func TestNewApiTokenID(t *testing.T) {
	id := NewApiTokenID("https://endpoint-url.example.com", "tokenId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.TokenId != "tokenId" {
		t.Fatalf("Expected %q but got %q for Segment 'TokenId'", id.TokenId, "tokenId")
	}
}

func TestFormatApiTokenID(t *testing.T) {
	actual := NewApiTokenID("https://endpoint-url.example.com", "tokenId").ID()
	expected := "https://endpoint-url.example.com/apiTokens/tokenId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseApiTokenID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApiTokenId
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
			Input: "https://endpoint-url.example.com/apiTokens",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/apiTokens/tokenId",
			Expected: &ApiTokenId{
				BaseURI: "https://endpoint-url.example.com",
				TokenId: "tokenId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/apiTokens/tokenId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApiTokenID(v.Input)
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

		if actual.TokenId != v.Expected.TokenId {
			t.Fatalf("Expected %q but got %q for TokenId", v.Expected.TokenId, actual.TokenId)
		}

	}
}

func TestParseApiTokenIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ApiTokenId
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
			Input: "https://endpoint-url.example.com/apiTokens",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/aPiToKeNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/apiTokens/tokenId",
			Expected: &ApiTokenId{
				BaseURI: "https://endpoint-url.example.com",
				TokenId: "tokenId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/apiTokens/tokenId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/aPiToKeNs/tOkEnId",
			Expected: &ApiTokenId{
				BaseURI: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				TokenId: "tOkEnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/aPiToKeNs/tOkEnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseApiTokenIDInsensitively(v.Input)
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

		if actual.TokenId != v.Expected.TokenId {
			t.Fatalf("Expected %q but got %q for TokenId", v.Expected.TokenId, actual.TokenId)
		}

	}
}

func TestSegmentsForApiTokenId(t *testing.T) {
	segments := ApiTokenId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ApiTokenId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
