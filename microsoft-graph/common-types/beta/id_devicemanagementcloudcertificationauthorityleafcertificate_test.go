package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudCertificationAuthorityLeafCertificateId{}

func TestNewDeviceManagementCloudCertificationAuthorityLeafCertificateID(t *testing.T) {
	id := NewDeviceManagementCloudCertificationAuthorityLeafCertificateID("cloudCertificationAuthorityLeafCertificateId")

	if id.CloudCertificationAuthorityLeafCertificateId != "cloudCertificationAuthorityLeafCertificateId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudCertificationAuthorityLeafCertificateId'", id.CloudCertificationAuthorityLeafCertificateId, "cloudCertificationAuthorityLeafCertificateId")
	}
}

func TestFormatDeviceManagementCloudCertificationAuthorityLeafCertificateID(t *testing.T) {
	actual := NewDeviceManagementCloudCertificationAuthorityLeafCertificateID("cloudCertificationAuthorityLeafCertificateId").ID()
	expected := "/deviceManagement/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCloudCertificationAuthorityLeafCertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudCertificationAuthorityLeafCertificateId
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
			Input: "/deviceManagement/cloudCertificationAuthorityLeafCertificate",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId",
			Expected: &DeviceManagementCloudCertificationAuthorityLeafCertificateId{
				CloudCertificationAuthorityLeafCertificateId: "cloudCertificationAuthorityLeafCertificateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudCertificationAuthorityLeafCertificateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudCertificationAuthorityLeafCertificateId != v.Expected.CloudCertificationAuthorityLeafCertificateId {
			t.Fatalf("Expected %q but got %q for CloudCertificationAuthorityLeafCertificateId", v.Expected.CloudCertificationAuthorityLeafCertificateId, actual.CloudCertificationAuthorityLeafCertificateId)
		}

	}
}

func TestParseDeviceManagementCloudCertificationAuthorityLeafCertificateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudCertificationAuthorityLeafCertificateId
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
			Input: "/deviceManagement/cloudCertificationAuthorityLeafCertificate",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId",
			Expected: &DeviceManagementCloudCertificationAuthorityLeafCertificateId{
				CloudCertificationAuthorityLeafCertificateId: "cloudCertificationAuthorityLeafCertificateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtE/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtEiD",
			Expected: &DeviceManagementCloudCertificationAuthorityLeafCertificateId{
				CloudCertificationAuthorityLeafCertificateId: "cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtE/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudCertificationAuthorityLeafCertificateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudCertificationAuthorityLeafCertificateId != v.Expected.CloudCertificationAuthorityLeafCertificateId {
			t.Fatalf("Expected %q but got %q for CloudCertificationAuthorityLeafCertificateId", v.Expected.CloudCertificationAuthorityLeafCertificateId, actual.CloudCertificationAuthorityLeafCertificateId)
		}

	}
}

func TestSegmentsForDeviceManagementCloudCertificationAuthorityLeafCertificateId(t *testing.T) {
	segments := DeviceManagementCloudCertificationAuthorityLeafCertificateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCloudCertificationAuthorityLeafCertificateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
