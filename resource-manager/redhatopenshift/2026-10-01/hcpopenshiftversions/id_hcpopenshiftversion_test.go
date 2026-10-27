package hcpopenshiftversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &HcpOpenShiftVersionId{}

func TestNewHcpOpenShiftVersionID(t *testing.T) {
	id := NewHcpOpenShiftVersionID("12345678-1234-9876-4563-123456789012", "locationName", "hcpOpenShiftVersionName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationName")
	}

	if id.HcpOpenShiftVersionName != "hcpOpenShiftVersionName" {
		t.Fatalf("Expected %q but got %q for Segment 'HcpOpenShiftVersionName'", id.HcpOpenShiftVersionName, "hcpOpenShiftVersionName")
	}
}

func TestFormatHcpOpenShiftVersionID(t *testing.T) {
	actual := NewHcpOpenShiftVersionID("12345678-1234-9876-4563-123456789012", "locationName", "hcpOpenShiftVersionName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions/hcpOpenShiftVersionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseHcpOpenShiftVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *HcpOpenShiftVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions/hcpOpenShiftVersionName",
			Expected: &HcpOpenShiftVersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				LocationName:            "locationName",
				HcpOpenShiftVersionName: "hcpOpenShiftVersionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions/hcpOpenShiftVersionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseHcpOpenShiftVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.HcpOpenShiftVersionName != v.Expected.HcpOpenShiftVersionName {
			t.Fatalf("Expected %q but got %q for HcpOpenShiftVersionName", v.Expected.HcpOpenShiftVersionName, actual.HcpOpenShiftVersionName)
		}

	}
}

func TestParseHcpOpenShiftVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *HcpOpenShiftVersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/lOcAtIoNs/lOcAtIoNnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/lOcAtIoNs/lOcAtIoNnAmE/hCpOpEnShIfTvErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions/hcpOpenShiftVersionName",
			Expected: &HcpOpenShiftVersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				LocationName:            "locationName",
				HcpOpenShiftVersionName: "hcpOpenShiftVersionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.RedHatOpenShift/locations/locationName/hcpOpenShiftVersions/hcpOpenShiftVersionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/lOcAtIoNs/lOcAtIoNnAmE/hCpOpEnShIfTvErSiOnS/hCpOpEnShIfTvErSiOnNaMe",
			Expected: &HcpOpenShiftVersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				LocationName:            "lOcAtIoNnAmE",
				HcpOpenShiftVersionName: "hCpOpEnShIfTvErSiOnNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.rEdHaToPeNsHiFt/lOcAtIoNs/lOcAtIoNnAmE/hCpOpEnShIfTvErSiOnS/hCpOpEnShIfTvErSiOnNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseHcpOpenShiftVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.HcpOpenShiftVersionName != v.Expected.HcpOpenShiftVersionName {
			t.Fatalf("Expected %q but got %q for HcpOpenShiftVersionName", v.Expected.HcpOpenShiftVersionName, actual.HcpOpenShiftVersionName)
		}

	}
}

func TestSegmentsForHcpOpenShiftVersionId(t *testing.T) {
	segments := HcpOpenShiftVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("HcpOpenShiftVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %d unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
