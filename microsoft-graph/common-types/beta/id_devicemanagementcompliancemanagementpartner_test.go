package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComplianceManagementPartnerId{}

func TestNewDeviceManagementComplianceManagementPartnerID(t *testing.T) {
	id := NewDeviceManagementComplianceManagementPartnerID("complianceManagementPartnerId")

	if id.ComplianceManagementPartnerId != "complianceManagementPartnerId" {
		t.Fatalf("Expected %q but got %q for Segment 'ComplianceManagementPartnerId'", id.ComplianceManagementPartnerId, "complianceManagementPartnerId")
	}
}

func TestFormatDeviceManagementComplianceManagementPartnerID(t *testing.T) {
	actual := NewDeviceManagementComplianceManagementPartnerID("complianceManagementPartnerId").ID()
	expected := "/deviceManagement/complianceManagementPartners/complianceManagementPartnerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComplianceManagementPartnerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComplianceManagementPartnerId
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
			Input: "/deviceManagement/complianceManagementPartners",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerId",
			Expected: &DeviceManagementComplianceManagementPartnerId{
				ComplianceManagementPartnerId: "complianceManagementPartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComplianceManagementPartnerID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ComplianceManagementPartnerId != v.Expected.ComplianceManagementPartnerId {
			t.Fatalf("Expected %q but got %q for ComplianceManagementPartnerId", v.Expected.ComplianceManagementPartnerId, actual.ComplianceManagementPartnerId)
		}

	}
}

func TestParseDeviceManagementComplianceManagementPartnerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComplianceManagementPartnerId
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
			Input: "/deviceManagement/complianceManagementPartners",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEmAnAgEmEnTpArTnErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerId",
			Expected: &DeviceManagementComplianceManagementPartnerId{
				ComplianceManagementPartnerId: "complianceManagementPartnerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEmAnAgEmEnTpArTnErS/cOmPlIaNcEmAnAgEmEnTpArTnErId",
			Expected: &DeviceManagementComplianceManagementPartnerId{
				ComplianceManagementPartnerId: "cOmPlIaNcEmAnAgEmEnTpArTnErId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEmAnAgEmEnTpArTnErS/cOmPlIaNcEmAnAgEmEnTpArTnErId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComplianceManagementPartnerIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ComplianceManagementPartnerId != v.Expected.ComplianceManagementPartnerId {
			t.Fatalf("Expected %q but got %q for ComplianceManagementPartnerId", v.Expected.ComplianceManagementPartnerId, actual.ComplianceManagementPartnerId)
		}

	}
}

func TestSegmentsForDeviceManagementComplianceManagementPartnerId(t *testing.T) {
	segments := DeviceManagementComplianceManagementPartnerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComplianceManagementPartnerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
