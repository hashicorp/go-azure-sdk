package agreementacceptance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeAgreementAcceptanceId{}

func TestNewMeAgreementAcceptanceID(t *testing.T) {
	id := NewMeAgreementAcceptanceID("agreementAcceptanceIdValue")

	if id.AgreementAcceptanceId != "agreementAcceptanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementAcceptanceId'", id.AgreementAcceptanceId, "agreementAcceptanceIdValue")
	}
}

func TestFormatMeAgreementAcceptanceID(t *testing.T) {
	actual := NewMeAgreementAcceptanceID("agreementAcceptanceIdValue").ID()
	expected := "/me/agreementAcceptances/agreementAcceptanceIdValue"
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
			Input: "/me/agreementAcceptances/agreementAcceptanceIdValue",
			Expected: &MeAgreementAcceptanceId{
				AgreementAcceptanceId: "agreementAcceptanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/agreementAcceptances/agreementAcceptanceIdValue/extra",
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
			Input: "/me/agreementAcceptances/agreementAcceptanceIdValue",
			Expected: &MeAgreementAcceptanceId{
				AgreementAcceptanceId: "agreementAcceptanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/agreementAcceptances/agreementAcceptanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeIdVaLuE",
			Expected: &MeAgreementAcceptanceId{
				AgreementAcceptanceId: "aGrEeMeNtAcCePtAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeIdVaLuE/extra",
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
