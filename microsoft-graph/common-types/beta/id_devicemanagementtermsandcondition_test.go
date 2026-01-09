package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionId{}

func TestNewDeviceManagementTermsAndConditionID(t *testing.T) {
	id := NewDeviceManagementTermsAndConditionID("termsAndConditionsId")

	if id.TermsAndConditionsId != "termsAndConditionsId" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsId'", id.TermsAndConditionsId, "termsAndConditionsId")
	}
}

func TestFormatDeviceManagementTermsAndConditionID(t *testing.T) {
	actual := NewDeviceManagementTermsAndConditionID("termsAndConditionsId").ID()
	expected := "/deviceManagement/termsAndConditions/termsAndConditionsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTermsAndConditionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId",
			Expected: &DeviceManagementTermsAndConditionId{
				TermsAndConditionsId: "termsAndConditionsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TermsAndConditionsId != v.Expected.TermsAndConditionsId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsId", v.Expected.TermsAndConditionsId, actual.TermsAndConditionsId)
		}

	}
}

func TestParseDeviceManagementTermsAndConditionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId",
			Expected: &DeviceManagementTermsAndConditionId{
				TermsAndConditionsId: "termsAndConditionsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiD",
			Expected: &DeviceManagementTermsAndConditionId{
				TermsAndConditionsId: "tErMsAnDcOnDiTiOnSiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.TermsAndConditionsId != v.Expected.TermsAndConditionsId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsId", v.Expected.TermsAndConditionsId, actual.TermsAndConditionsId)
		}

	}
}

func TestSegmentsForDeviceManagementTermsAndConditionId(t *testing.T) {
	segments := DeviceManagementTermsAndConditionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTermsAndConditionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
