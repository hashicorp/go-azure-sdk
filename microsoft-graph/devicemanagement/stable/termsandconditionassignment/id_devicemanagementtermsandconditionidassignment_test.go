package termsandconditionassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionIdAssignmentId{}

func TestNewDeviceManagementTermsAndConditionIdAssignmentID(t *testing.T) {
	id := NewDeviceManagementTermsAndConditionIdAssignmentID("termsAndConditionsIdValue", "termsAndConditionsAssignmentIdValue")

	if id.TermsAndConditionsId != "termsAndConditionsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsId'", id.TermsAndConditionsId, "termsAndConditionsIdValue")
	}

	if id.TermsAndConditionsAssignmentId != "termsAndConditionsAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsAssignmentId'", id.TermsAndConditionsAssignmentId, "termsAndConditionsAssignmentIdValue")
	}
}

func TestFormatDeviceManagementTermsAndConditionIdAssignmentID(t *testing.T) {
	actual := NewDeviceManagementTermsAndConditionIdAssignmentID("termsAndConditionsIdValue", "termsAndConditionsAssignmentIdValue").ID()
	expected := "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments/termsAndConditionsAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTermsAndConditionIdAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionIdAssignmentId
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
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments/termsAndConditionsAssignmentIdValue",
			Expected: &DeviceManagementTermsAndConditionIdAssignmentId{
				TermsAndConditionsId:           "termsAndConditionsIdValue",
				TermsAndConditionsAssignmentId: "termsAndConditionsAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments/termsAndConditionsAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIdAssignmentID(v.Input)
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

		if actual.TermsAndConditionsAssignmentId != v.Expected.TermsAndConditionsAssignmentId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsAssignmentId", v.Expected.TermsAndConditionsAssignmentId, actual.TermsAndConditionsAssignmentId)
		}

	}
}

func TestParseDeviceManagementTermsAndConditionIdAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionIdAssignmentId
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
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe/aSsIgNmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments/termsAndConditionsAssignmentIdValue",
			Expected: &DeviceManagementTermsAndConditionIdAssignmentId{
				TermsAndConditionsId:           "termsAndConditionsIdValue",
				TermsAndConditionsAssignmentId: "termsAndConditionsAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/assignments/termsAndConditionsAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe/aSsIgNmEnTs/tErMsAnDcOnDiTiOnSaSsIgNmEnTiDvAlUe",
			Expected: &DeviceManagementTermsAndConditionIdAssignmentId{
				TermsAndConditionsId:           "tErMsAnDcOnDiTiOnSiDvAlUe",
				TermsAndConditionsAssignmentId: "tErMsAnDcOnDiTiOnSaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe/aSsIgNmEnTs/tErMsAnDcOnDiTiOnSaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIdAssignmentIDInsensitively(v.Input)
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

		if actual.TermsAndConditionsAssignmentId != v.Expected.TermsAndConditionsAssignmentId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsAssignmentId", v.Expected.TermsAndConditionsAssignmentId, actual.TermsAndConditionsAssignmentId)
		}

	}
}

func TestSegmentsForDeviceManagementTermsAndConditionIdAssignmentId(t *testing.T) {
	segments := DeviceManagementTermsAndConditionIdAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTermsAndConditionIdAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
