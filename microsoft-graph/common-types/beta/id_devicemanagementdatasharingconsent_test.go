package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDataSharingConsentId{}

func TestNewDeviceManagementDataSharingConsentID(t *testing.T) {
	id := NewDeviceManagementDataSharingConsentID("dataSharingConsentIdValue")

	if id.DataSharingConsentId != "dataSharingConsentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DataSharingConsentId'", id.DataSharingConsentId, "dataSharingConsentIdValue")
	}
}

func TestFormatDeviceManagementDataSharingConsentID(t *testing.T) {
	actual := NewDeviceManagementDataSharingConsentID("dataSharingConsentIdValue").ID()
	expected := "/deviceManagement/dataSharingConsents/dataSharingConsentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDataSharingConsentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDataSharingConsentId
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
			Input: "/deviceManagement/dataSharingConsents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/dataSharingConsents/dataSharingConsentIdValue",
			Expected: &DeviceManagementDataSharingConsentId{
				DataSharingConsentId: "dataSharingConsentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/dataSharingConsents/dataSharingConsentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDataSharingConsentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DataSharingConsentId != v.Expected.DataSharingConsentId {
			t.Fatalf("Expected %q but got %q for DataSharingConsentId", v.Expected.DataSharingConsentId, actual.DataSharingConsentId)
		}

	}
}

func TestParseDeviceManagementDataSharingConsentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDataSharingConsentId
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
			Input: "/deviceManagement/dataSharingConsents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dAtAsHaRiNgCoNsEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/dataSharingConsents/dataSharingConsentIdValue",
			Expected: &DeviceManagementDataSharingConsentId{
				DataSharingConsentId: "dataSharingConsentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/dataSharingConsents/dataSharingConsentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dAtAsHaRiNgCoNsEnTs/dAtAsHaRiNgCoNsEnTiDvAlUe",
			Expected: &DeviceManagementDataSharingConsentId{
				DataSharingConsentId: "dAtAsHaRiNgCoNsEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dAtAsHaRiNgCoNsEnTs/dAtAsHaRiNgCoNsEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDataSharingConsentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DataSharingConsentId != v.Expected.DataSharingConsentId {
			t.Fatalf("Expected %q but got %q for DataSharingConsentId", v.Expected.DataSharingConsentId, actual.DataSharingConsentId)
		}

	}
}

func TestSegmentsForDeviceManagementDataSharingConsentId(t *testing.T) {
	segments := DeviceManagementDataSharingConsentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDataSharingConsentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
