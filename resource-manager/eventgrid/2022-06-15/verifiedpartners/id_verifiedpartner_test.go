package verifiedpartners

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VerifiedPartnerId{}

func TestNewVerifiedPartnerID(t *testing.T) {
	id := NewVerifiedPartnerID("verifiedPartnerValue")

	if id.VerifiedPartnerName != "verifiedPartnerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'VerifiedPartnerName'", id.VerifiedPartnerName, "verifiedPartnerValue")
	}
}

func TestFormatVerifiedPartnerID(t *testing.T) {
	actual := NewVerifiedPartnerID("verifiedPartnerValue").ID()
	expected := "/providers/Microsoft.EventGrid/verifiedPartners/verifiedPartnerValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVerifiedPartnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VerifiedPartnerId
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
			Input: "/providers/Microsoft.EventGrid",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.EventGrid/verifiedPartners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.EventGrid/verifiedPartners/verifiedPartnerValue",
			Expected: &VerifiedPartnerId{
				VerifiedPartnerName: "verifiedPartnerValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.EventGrid/verifiedPartners/verifiedPartnerValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVerifiedPartnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.VerifiedPartnerName != v.Expected.VerifiedPartnerName {
			t.Fatalf("Expected %q but got %q for VerifiedPartnerName", v.Expected.VerifiedPartnerName, actual.VerifiedPartnerName)
		}

	}
}

func TestParseVerifiedPartnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VerifiedPartnerId
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
			Input: "/providers/Microsoft.EventGrid",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.EventGrid/verifiedPartners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId/vErIfIeDpArTnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.EventGrid/verifiedPartners/verifiedPartnerValue",
			Expected: &VerifiedPartnerId{
				VerifiedPartnerName: "verifiedPartnerValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.EventGrid/verifiedPartners/verifiedPartnerValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId/vErIfIeDpArTnErS/vErIfIeDpArTnErVaLuE",
			Expected: &VerifiedPartnerId{
				VerifiedPartnerName: "vErIfIeDpArTnErVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.eVeNtGrId/vErIfIeDpArTnErS/vErIfIeDpArTnErVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVerifiedPartnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.VerifiedPartnerName != v.Expected.VerifiedPartnerName {
			t.Fatalf("Expected %q but got %q for VerifiedPartnerName", v.Expected.VerifiedPartnerName, actual.VerifiedPartnerName)
		}

	}
}

func TestSegmentsForVerifiedPartnerId(t *testing.T) {
	segments := VerifiedPartnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VerifiedPartnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
