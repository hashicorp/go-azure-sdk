package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAgreementAcceptanceId{}

func TestNewMeAgreementAcceptanceID(t *testing.T) {
	id := NewMeAgreementAcceptanceID("agreementAcceptanceId")

	if id.AgreementAcceptanceId != "agreementAcceptanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementAcceptanceId'", id.AgreementAcceptanceId, "agreementAcceptanceId")
	}
}

func TestFormatMeAgreementAcceptanceID(t *testing.T) {
	actual := NewMeAgreementAcceptanceID("agreementAcceptanceId").ID()
	expected := "/me/agreementAcceptances/agreementAcceptanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeAgreementAcceptanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAgreementAcceptanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/agreementAcceptances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/agreementAcceptances/agreementAcceptanceId",
			Expected: &MeAgreementAcceptanceId{
				AgreementAcceptanceId: "agreementAcceptanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/agreementAcceptances/agreementAcceptanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAgreementAcceptanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestParseMeAgreementAcceptanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeAgreementAcceptanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/agreementAcceptances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aGrEeMeNtAcCePtAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/agreementAcceptances/agreementAcceptanceId",
			Expected: &MeAgreementAcceptanceId{
				AgreementAcceptanceId: "agreementAcceptanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/agreementAcceptances/agreementAcceptanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeId",
			Expected: &MeAgreementAcceptanceId{
				AgreementAcceptanceId: "aGrEeMeNtAcCePtAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeAgreementAcceptanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestSegmentsForMeAgreementAcceptanceId(t *testing.T) {
	segments := MeAgreementAcceptanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeAgreementAcceptanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
