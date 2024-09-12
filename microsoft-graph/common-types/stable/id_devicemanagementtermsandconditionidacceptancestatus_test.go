package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTermsAndConditionIdAcceptanceStatusId{}

func TestNewDeviceManagementTermsAndConditionIdAcceptanceStatusID(t *testing.T) {
	id := NewDeviceManagementTermsAndConditionIdAcceptanceStatusID("termsAndConditionsIdValue", "termsAndConditionsAcceptanceStatusIdValue")

	if id.TermsAndConditionsId != "termsAndConditionsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsId'", id.TermsAndConditionsId, "termsAndConditionsIdValue")
	}

	if id.TermsAndConditionsAcceptanceStatusId != "termsAndConditionsAcceptanceStatusIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TermsAndConditionsAcceptanceStatusId'", id.TermsAndConditionsAcceptanceStatusId, "termsAndConditionsAcceptanceStatusIdValue")
	}
}

func TestFormatDeviceManagementTermsAndConditionIdAcceptanceStatusID(t *testing.T) {
	actual := NewDeviceManagementTermsAndConditionIdAcceptanceStatusID("termsAndConditionsIdValue", "termsAndConditionsAcceptanceStatusIdValue").ID()
	expected := "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses/termsAndConditionsAcceptanceStatusIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTermsAndConditionIdAcceptanceStatusID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionIdAcceptanceStatusId
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
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses/termsAndConditionsAcceptanceStatusIdValue",
			Expected: &DeviceManagementTermsAndConditionIdAcceptanceStatusId{
				TermsAndConditionsId:                 "termsAndConditionsIdValue",
				TermsAndConditionsAcceptanceStatusId: "termsAndConditionsAcceptanceStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses/termsAndConditionsAcceptanceStatusIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIdAcceptanceStatusID(v.Input)
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

		if actual.TermsAndConditionsAcceptanceStatusId != v.Expected.TermsAndConditionsAcceptanceStatusId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsAcceptanceStatusId", v.Expected.TermsAndConditionsAcceptanceStatusId, actual.TermsAndConditionsAcceptanceStatusId)
		}

	}
}

func TestParseDeviceManagementTermsAndConditionIdAcceptanceStatusIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTermsAndConditionIdAcceptanceStatusId
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
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe/aCcEpTaNcEsTaTuSeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses/termsAndConditionsAcceptanceStatusIdValue",
			Expected: &DeviceManagementTermsAndConditionIdAcceptanceStatusId{
				TermsAndConditionsId:                 "termsAndConditionsIdValue",
				TermsAndConditionsAcceptanceStatusId: "termsAndConditionsAcceptanceStatusIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/termsAndConditions/termsAndConditionsIdValue/acceptanceStatuses/termsAndConditionsAcceptanceStatusIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe/aCcEpTaNcEsTaTuSeS/tErMsAnDcOnDiTiOnSaCcEpTaNcEsTaTuSiDvAlUe",
			Expected: &DeviceManagementTermsAndConditionIdAcceptanceStatusId{
				TermsAndConditionsId:                 "tErMsAnDcOnDiTiOnSiDvAlUe",
				TermsAndConditionsAcceptanceStatusId: "tErMsAnDcOnDiTiOnSaCcEpTaNcEsTaTuSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tErMsAnDcOnDiTiOnS/tErMsAnDcOnDiTiOnSiDvAlUe/aCcEpTaNcEsTaTuSeS/tErMsAnDcOnDiTiOnSaCcEpTaNcEsTaTuSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTermsAndConditionIdAcceptanceStatusIDInsensitively(v.Input)
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

		if actual.TermsAndConditionsAcceptanceStatusId != v.Expected.TermsAndConditionsAcceptanceStatusId {
			t.Fatalf("Expected %q but got %q for TermsAndConditionsAcceptanceStatusId", v.Expected.TermsAndConditionsAcceptanceStatusId, actual.TermsAndConditionsAcceptanceStatusId)
		}

	}
}

func TestSegmentsForDeviceManagementTermsAndConditionIdAcceptanceStatusId(t *testing.T) {
	segments := DeviceManagementTermsAndConditionIdAcceptanceStatusId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTermsAndConditionIdAcceptanceStatusId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
