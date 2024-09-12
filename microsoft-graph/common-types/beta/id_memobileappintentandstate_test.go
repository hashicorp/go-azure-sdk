package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMobileAppIntentAndStateId{}

func TestNewMeMobileAppIntentAndStateID(t *testing.T) {
	id := NewMeMobileAppIntentAndStateID("mobileAppIntentAndStateIdValue")

	if id.MobileAppIntentAndStateId != "mobileAppIntentAndStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MobileAppIntentAndStateId'", id.MobileAppIntentAndStateId, "mobileAppIntentAndStateIdValue")
	}
}

func TestFormatMeMobileAppIntentAndStateID(t *testing.T) {
	actual := NewMeMobileAppIntentAndStateID("mobileAppIntentAndStateIdValue").ID()
	expected := "/me/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMobileAppIntentAndStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMobileAppIntentAndStateId
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
			Input: "/me/mobileAppIntentAndStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue",
			Expected: &MeMobileAppIntentAndStateId{
				MobileAppIntentAndStateId: "mobileAppIntentAndStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMobileAppIntentAndStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobileAppIntentAndStateId != v.Expected.MobileAppIntentAndStateId {
			t.Fatalf("Expected %q but got %q for MobileAppIntentAndStateId", v.Expected.MobileAppIntentAndStateId, actual.MobileAppIntentAndStateId)
		}

	}
}

func TestParseMeMobileAppIntentAndStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMobileAppIntentAndStateId
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
			Input: "/me/mobileAppIntentAndStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpInTeNtAnDsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue",
			Expected: &MeMobileAppIntentAndStateId{
				MobileAppIntentAndStateId: "mobileAppIntentAndStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mobileAppIntentAndStates/mobileAppIntentAndStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpInTeNtAnDsTaTeS/mObIlEaPpInTeNtAnDsTaTeIdVaLuE",
			Expected: &MeMobileAppIntentAndStateId{
				MobileAppIntentAndStateId: "mObIlEaPpInTeNtAnDsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mObIlEaPpInTeNtAnDsTaTeS/mObIlEaPpInTeNtAnDsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMobileAppIntentAndStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MobileAppIntentAndStateId != v.Expected.MobileAppIntentAndStateId {
			t.Fatalf("Expected %q but got %q for MobileAppIntentAndStateId", v.Expected.MobileAppIntentAndStateId, actual.MobileAppIntentAndStateId)
		}

	}
}

func TestSegmentsForMeMobileAppIntentAndStateId(t *testing.T) {
	segments := MeMobileAppIntentAndStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMobileAppIntentAndStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
