package pricings

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PricingId{}

func TestNewPricingID(t *testing.T) {
	id := NewPricingID("scopeIdValue", "pricingValue")

	if id.ScopeId != "scopeIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ScopeId'", id.ScopeId, "scopeIdValue")
	}

	if id.PricingName != "pricingValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PricingName'", id.PricingName, "pricingValue")
	}
}

func TestFormatPricingID(t *testing.T) {
	actual := NewPricingID("scopeIdValue", "pricingValue").ID()
	expected := "/scopeIdValue/providers/Microsoft.Security/pricings/pricingValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePricingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PricingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue/providers/Microsoft.Security/pricings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/scopeIdValue/providers/Microsoft.Security/pricings/pricingValue",
			Expected: &PricingId{
				ScopeId:     "scopeIdValue",
				PricingName: "pricingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/scopeIdValue/providers/Microsoft.Security/pricings/pricingValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePricingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScopeId != v.Expected.ScopeId {
			t.Fatalf("Expected %q but got %q for ScopeId", v.Expected.ScopeId, actual.ScopeId)
		}

		if actual.PricingName != v.Expected.PricingName {
			t.Fatalf("Expected %q but got %q for PricingName", v.Expected.PricingName, actual.PricingName)
		}

	}
}

func TestParsePricingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PricingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sCoPeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sCoPeIdVaLuE/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue/providers/Microsoft.Security",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sCoPeIdVaLuE/pRoViDeRs/mIcRoSoFt.sEcUrItY",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/scopeIdValue/providers/Microsoft.Security/pricings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sCoPeIdVaLuE/pRoViDeRs/mIcRoSoFt.sEcUrItY/pRiCiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/scopeIdValue/providers/Microsoft.Security/pricings/pricingValue",
			Expected: &PricingId{
				ScopeId:     "scopeIdValue",
				PricingName: "pricingValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/scopeIdValue/providers/Microsoft.Security/pricings/pricingValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sCoPeIdVaLuE/pRoViDeRs/mIcRoSoFt.sEcUrItY/pRiCiNgS/pRiCiNgVaLuE",
			Expected: &PricingId{
				ScopeId:     "sCoPeIdVaLuE",
				PricingName: "pRiCiNgVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sCoPeIdVaLuE/pRoViDeRs/mIcRoSoFt.sEcUrItY/pRiCiNgS/pRiCiNgVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePricingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScopeId != v.Expected.ScopeId {
			t.Fatalf("Expected %q but got %q for ScopeId", v.Expected.ScopeId, actual.ScopeId)
		}

		if actual.PricingName != v.Expected.PricingName {
			t.Fatalf("Expected %q but got %q for PricingName", v.Expected.PricingName, actual.PricingName)
		}

	}
}

func TestSegmentsForPricingId(t *testing.T) {
	segments := PricingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PricingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
