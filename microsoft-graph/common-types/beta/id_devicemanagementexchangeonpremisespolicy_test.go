package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementExchangeOnPremisesPolicyId{}

func TestNewDeviceManagementExchangeOnPremisesPolicyID(t *testing.T) {
	id := NewDeviceManagementExchangeOnPremisesPolicyID("deviceManagementExchangeOnPremisesPolicyId")

	if id.DeviceManagementExchangeOnPremisesPolicyId != "deviceManagementExchangeOnPremisesPolicyId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementExchangeOnPremisesPolicyId'", id.DeviceManagementExchangeOnPremisesPolicyId, "deviceManagementExchangeOnPremisesPolicyId")
	}
}

func TestFormatDeviceManagementExchangeOnPremisesPolicyID(t *testing.T) {
	actual := NewDeviceManagementExchangeOnPremisesPolicyID("deviceManagementExchangeOnPremisesPolicyId").ID()
	expected := "/deviceManagement/exchangeOnPremisesPolicies/deviceManagementExchangeOnPremisesPolicyId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementExchangeOnPremisesPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementExchangeOnPremisesPolicyId
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
			Input: "/deviceManagement/exchangeOnPremisesPolicies",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/exchangeOnPremisesPolicies/deviceManagementExchangeOnPremisesPolicyId",
			Expected: &DeviceManagementExchangeOnPremisesPolicyId{
				DeviceManagementExchangeOnPremisesPolicyId: "deviceManagementExchangeOnPremisesPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/exchangeOnPremisesPolicies/deviceManagementExchangeOnPremisesPolicyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementExchangeOnPremisesPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementExchangeOnPremisesPolicyId != v.Expected.DeviceManagementExchangeOnPremisesPolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementExchangeOnPremisesPolicyId", v.Expected.DeviceManagementExchangeOnPremisesPolicyId, actual.DeviceManagementExchangeOnPremisesPolicyId)
		}

	}
}

func TestParseDeviceManagementExchangeOnPremisesPolicyIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementExchangeOnPremisesPolicyId
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
			Input: "/deviceManagement/exchangeOnPremisesPolicies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eXcHaNgEoNpReMiSeSpOlIcIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/exchangeOnPremisesPolicies/deviceManagementExchangeOnPremisesPolicyId",
			Expected: &DeviceManagementExchangeOnPremisesPolicyId{
				DeviceManagementExchangeOnPremisesPolicyId: "deviceManagementExchangeOnPremisesPolicyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/exchangeOnPremisesPolicies/deviceManagementExchangeOnPremisesPolicyId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eXcHaNgEoNpReMiSeSpOlIcIeS/dEvIcEmAnAgEmEnTeXcHaNgEoNpReMiSeSpOlIcYiD",
			Expected: &DeviceManagementExchangeOnPremisesPolicyId{
				DeviceManagementExchangeOnPremisesPolicyId: "dEvIcEmAnAgEmEnTeXcHaNgEoNpReMiSeSpOlIcYiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/eXcHaNgEoNpReMiSeSpOlIcIeS/dEvIcEmAnAgEmEnTeXcHaNgEoNpReMiSeSpOlIcYiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementExchangeOnPremisesPolicyIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementExchangeOnPremisesPolicyId != v.Expected.DeviceManagementExchangeOnPremisesPolicyId {
			t.Fatalf("Expected %q but got %q for DeviceManagementExchangeOnPremisesPolicyId", v.Expected.DeviceManagementExchangeOnPremisesPolicyId, actual.DeviceManagementExchangeOnPremisesPolicyId)
		}

	}
}

func TestSegmentsForDeviceManagementExchangeOnPremisesPolicyId(t *testing.T) {
	segments := DeviceManagementExchangeOnPremisesPolicyId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementExchangeOnPremisesPolicyId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
