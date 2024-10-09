package provideroperationsmetadata

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProviderOperationId{}

func TestNewProviderOperationID(t *testing.T) {
	id := NewProviderOperationID("providerOperationName")

	if id.ProviderOperationName != "providerOperationName" {
		t.Fatalf("Expected %q but got %q for Segment 'ProviderOperationName'", id.ProviderOperationName, "providerOperationName")
	}
}

func TestFormatProviderOperationID(t *testing.T) {
	actual := NewProviderOperationID("providerOperationName").ID()
	expected := "/providers/Microsoft.Authorization/providerOperations/providerOperationName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderOperationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderOperationId
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
			Input: "/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/providerOperations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Authorization/providerOperations/providerOperationName",
			Expected: &ProviderOperationId{
				ProviderOperationName: "providerOperationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/providerOperations/providerOperationName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderOperationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProviderOperationName != v.Expected.ProviderOperationName {
			t.Fatalf("Expected %q but got %q for ProviderOperationName", v.Expected.ProviderOperationName, actual.ProviderOperationName)
		}

	}
}

func TestParseProviderOperationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderOperationId
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
			Input: "/providers/Microsoft.Authorization",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Authorization/providerOperations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pRoViDeRoPeRaTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Authorization/providerOperations/providerOperationName",
			Expected: &ProviderOperationId{
				ProviderOperationName: "providerOperationName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Authorization/providerOperations/providerOperationName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pRoViDeRoPeRaTiOnS/pRoViDeRoPeRaTiOnNaMe",
			Expected: &ProviderOperationId{
				ProviderOperationName: "pRoViDeRoPeRaTiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.aUtHoRiZaTiOn/pRoViDeRoPeRaTiOnS/pRoViDeRoPeRaTiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderOperationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ProviderOperationName != v.Expected.ProviderOperationName {
			t.Fatalf("Expected %q but got %q for ProviderOperationName", v.Expected.ProviderOperationName, actual.ProviderOperationName)
		}

	}
}

func TestSegmentsForProviderOperationId(t *testing.T) {
	segments := ProviderOperationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderOperationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
