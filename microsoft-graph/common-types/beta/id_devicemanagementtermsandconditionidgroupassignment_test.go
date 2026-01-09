package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionIdGroupAssignmentId{}

func TestNewDeviceManagementTermsAndConditionIdGroupAssignmentID(t *testing.T) {
	id := NewDeviceManagementTermsAndConditionIdGroupAssignmentID("termsAndConditionsId", "termsAndConditionsGroupAssignmentId")

	if id.TermsAndConditionsId != "termsAndConditionsId" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsId'", id.TermsAndConditionsId, "termsAndConditionsId")
	}

	if id.TermsAndConditionsGroupAssignmentId != "termsAndConditionsGroupAssignmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsGroupAssignmentId'", id.TermsAndConditionsGroupAssignmentId, "termsAndConditionsGroupAssignmentId")
	}
}

func TestFormatDeviceManagementTermsAndConditionIdGroupAssignmentID(t *testing.T) {
	actual := NewDeviceManagementTermsAndConditionIdGroupAssignmentID("termsAndConditionsId", "termsAndConditionsGroupAssignmentId").ID()
	expected := "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments/termsAndConditionsGroupAssignmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTermsAndConditionIdGroupAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionIdGroupAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments/termsAndConditionsGroupAssignmentId",
			Expected: &DeviceManagementTermsAndConditionIdGroupAssignmentId{
				TermsAndConditionsId:                "termsAndConditionsId",
				TermsAndConditionsGroupAssignmentId: "termsAndConditionsGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments/termsAndConditionsGroupAssignmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIdGroupAssignmentID(v.Input)
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

		if actual.TermsAndConditionsGroupAssignmentId != v.Expected.TermsAndConditionsGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsGroupAssignmentId", v.Expected.TermsAndConditionsGroupAssignmentId, actual.TermsAndConditionsGroupAssignmentId)
		}

	}
}

func TestParseDeviceManagementTermsAndConditionIdGroupAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionIdGroupAssignmentId
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
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiD/gRoUpAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments/termsAndConditionsGroupAssignmentId",
			Expected: &DeviceManagementTermsAndConditionIdGroupAssignmentId{
				TermsAndConditionsId:                "termsAndConditionsId",
				TermsAndConditionsGroupAssignmentId: "termsAndConditionsGroupAssignmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsId/groupAssignments/termsAndConditionsGroupAssignmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiD/gRoUpAsSiGnMeNtS/tErMsAnDcOnDiTiOnSgRoUpAsSiGnMeNtId",
			Expected: &DeviceManagementTermsAndConditionIdGroupAssignmentId{
				TermsAndConditionsId:                "tErMsAnDcOnDiTiOnSiD",
				TermsAndConditionsGroupAssignmentId: "tErMsAnDcOnDiTiOnSgRoUpAsSiGnMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiD/gRoUpAsSiGnMeNtS/tErMsAnDcOnDiTiOnSgRoUpAsSiGnMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIdGroupAssignmentIDInsensitively(v.Input)
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

		if actual.TermsAndConditionsGroupAssignmentId != v.Expected.TermsAndConditionsGroupAssignmentId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsGroupAssignmentId", v.Expected.TermsAndConditionsGroupAssignmentId, actual.TermsAndConditionsGroupAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementTermsAndConditionIdGroupAssignmentId(t *testing.T) {
	segments := DeviceManagementTermsAndConditionIdGroupAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTermsAndConditionIdGroupAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
