package subscriptions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AliasId{}

func TestNewAliasID(t *testing.T) {
	id := NewAliasID("aliasName")

	if id.AliasName != "aliasName" {
		t.Fatalf("Expected %q but got %q for Segment 'AliasName'", id.AliasName, "aliasName")
	}
}

func TestFormatAliasID(t *testing.T) {
	actual := NewAliasID("aliasName").ID()
	expected := "/providers/Microsoft.Subscription/aliases/aliasName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAliasID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AliasId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Subscription",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Subscription/aliases",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Subscription/aliases/aliasName",
			Expected: &AliasId{
				AliasName: "aliasName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Subscription/aliases/aliasName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAliasID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AliasName != v.Expected.AliasName {
			t.Fatalf("Expected %q but got %q for AliasName", v.Expected.AliasName, actual.AliasName)
		}

	}
}

func TestParseAliasIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AliasId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Subscription",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Subscription/aliases",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/aLiAsEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Subscription/aliases/aliasName",
			Expected: &AliasId{
				AliasName: "aliasName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Subscription/aliases/aliasName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/aLiAsEs/aLiAsNaMe",
			Expected: &AliasId{
				AliasName: "aLiAsNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.sUbScRiPtIoN/aLiAsEs/aLiAsNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAliasIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AliasName != v.Expected.AliasName {
			t.Fatalf("Expected %q but got %q for AliasName", v.Expected.AliasName, actual.AliasName)
		}

	}
}

func TestSegmentsForAliasId(t *testing.T) {
	segments := AliasId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AliasId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
