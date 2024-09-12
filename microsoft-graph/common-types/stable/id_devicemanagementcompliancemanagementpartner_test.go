package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComplianceManagementPartnerId{}

func TestNewDeviceManagementComplianceManagementPartnerID(t *testing.T) {
	id := NewDeviceManagementComplianceManagementPartnerID("complianceManagementPartnerIdValue")

	if id.ComplianceManagementPartnerId != "complianceManagementPartnerIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ComplianceManagementPartnerId'", id.ComplianceManagementPartnerId, "complianceManagementPartnerIdValue")
	}
}

func TestFormatDeviceManagementComplianceManagementPartnerID(t *testing.T) {
	actual := NewDeviceManagementComplianceManagementPartnerID("complianceManagementPartnerIdValue").ID()
	expected := "/deviceManagement/complianceManagementPartners/complianceManagementPartnerIdValue"
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
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerIdValue",
			Expected: &DeviceManagementComplianceManagementPartnerId{
				ComplianceManagementPartnerId: "complianceManagementPartnerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerIdValue/extra",
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
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerIdValue",
			Expected: &DeviceManagementComplianceManagementPartnerId{
				ComplianceManagementPartnerId: "complianceManagementPartnerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/complianceManagementPartners/complianceManagementPartnerIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEmAnAgEmEnTpArTnErS/cOmPlIaNcEmAnAgEmEnTpArTnErIdVaLuE",
			Expected: &DeviceManagementComplianceManagementPartnerId{
				ComplianceManagementPartnerId: "cOmPlIaNcEmAnAgEmEnTpArTnErIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmPlIaNcEmAnAgEmEnTpArTnErS/cOmPlIaNcEmAnAgEmEnTpArTnErIdVaLuE/extra",
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
