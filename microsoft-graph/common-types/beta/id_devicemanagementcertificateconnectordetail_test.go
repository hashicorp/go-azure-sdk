package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCertificateConnectorDetailId{}

func TestNewDeviceManagementCertificateConnectorDetailID(t *testing.T) {
	id := NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsId")

	if id.CertificateConnectorDetailsId != "certificateConnectorDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateConnectorDetailsId'", id.CertificateConnectorDetailsId, "certificateConnectorDetailsId")
	}
}

func TestFormatDeviceManagementCertificateConnectorDetailID(t *testing.T) {
	actual := NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsId").ID()
	expected := "/deviceManagement/certificateConnectorDetails/certificateConnectorDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCertificateConnectorDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCertificateConnectorDetailId
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
			Input: "/deviceManagement/certificateConnectorDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/certificateConnectorDetails/certificateConnectorDetailsId",
			Expected: &DeviceManagementCertificateConnectorDetailId{
				CertificateConnectorDetailsId: "certificateConnectorDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/certificateConnectorDetails/certificateConnectorDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCertificateConnectorDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CertificateConnectorDetailsId != v.Expected.CertificateConnectorDetailsId {
			t.Fatalf("Expected %q but got %q for CertificateConnectorDetailsId", v.Expected.CertificateConnectorDetailsId, actual.CertificateConnectorDetailsId)
		}

	}
}

func TestParseDeviceManagementCertificateConnectorDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCertificateConnectorDetailId
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
			Input: "/deviceManagement/certificateConnectorDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cErTiFiCaTeCoNnEcToRdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/certificateConnectorDetails/certificateConnectorDetailsId",
			Expected: &DeviceManagementCertificateConnectorDetailId{
				CertificateConnectorDetailsId: "certificateConnectorDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/certificateConnectorDetails/certificateConnectorDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cErTiFiCaTeCoNnEcToRdEtAiLs/cErTiFiCaTeCoNnEcToRdEtAiLsId",
			Expected: &DeviceManagementCertificateConnectorDetailId{
				CertificateConnectorDetailsId: "cErTiFiCaTeCoNnEcToRdEtAiLsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cErTiFiCaTeCoNnEcToRdEtAiLs/cErTiFiCaTeCoNnEcToRdEtAiLsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCertificateConnectorDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CertificateConnectorDetailsId != v.Expected.CertificateConnectorDetailsId {
			t.Fatalf("Expected %q but got %q for CertificateConnectorDetailsId", v.Expected.CertificateConnectorDetailsId, actual.CertificateConnectorDetailsId)
		}

	}
}

func TestSegmentsForDeviceManagementCertificateConnectorDetailId(t *testing.T) {
	segments := DeviceManagementCertificateConnectorDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCertificateConnectorDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
