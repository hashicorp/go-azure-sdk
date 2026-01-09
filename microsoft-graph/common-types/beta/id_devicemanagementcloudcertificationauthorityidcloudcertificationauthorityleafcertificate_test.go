package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{}

func TestNewDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(t *testing.T) {
	id := NewDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID("cloudCertificationAuthorityId", "cloudCertificationAuthorityLeafCertificateId")

	if id.CloudCertificationAuthorityId != "cloudCertificationAuthorityId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudCertificationAuthorityId'", id.CloudCertificationAuthorityId, "cloudCertificationAuthorityId")
	}

	if id.CloudCertificationAuthorityLeafCertificateId != "cloudCertificationAuthorityLeafCertificateId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudCertificationAuthorityLeafCertificateId'", id.CloudCertificationAuthorityLeafCertificateId, "cloudCertificationAuthorityLeafCertificateId")
	}
}

func TestFormatDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(t *testing.T) {
	actual := NewDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID("cloudCertificationAuthorityId", "cloudCertificationAuthorityLeafCertificateId").ID()
	expected := "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId
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
			Input: "/deviceManagement/cloudCertificationAuthority",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId",
			Expected: &DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{
				CloudCertificationAuthorityId:                "cloudCertificationAuthorityId",
				CloudCertificationAuthorityLeafCertificateId: "cloudCertificationAuthorityLeafCertificateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudCertificationAuthorityId != v.Expected.CloudCertificationAuthorityId {
			t.Fatalf("Expected %q but got %q for CloudCertificationAuthorityId", v.Expected.CloudCertificationAuthorityId, actual.CloudCertificationAuthorityId)
		}

		if actual.CloudCertificationAuthorityLeafCertificateId != v.Expected.CloudCertificationAuthorityLeafCertificateId {
			t.Fatalf("Expected %q but got %q for CloudCertificationAuthorityLeafCertificateId", v.Expected.CloudCertificationAuthorityLeafCertificateId, actual.CloudCertificationAuthorityLeafCertificateId)
		}

	}
}

func TestParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId
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
			Input: "/deviceManagement/cloudCertificationAuthority",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy/cLoUdCeRtIfIcAtIoNaUtHoRiTyId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy/cLoUdCeRtIfIcAtIoNaUtHoRiTyId/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtE",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId",
			Expected: &DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{
				CloudCertificationAuthorityId:                "cloudCertificationAuthorityId",
				CloudCertificationAuthorityLeafCertificateId: "cloudCertificationAuthorityLeafCertificateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/cloudCertificationAuthorityLeafCertificate/cloudCertificationAuthorityLeafCertificateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy/cLoUdCeRtIfIcAtIoNaUtHoRiTyId/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtE/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtEiD",
			Expected: &DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{
				CloudCertificationAuthorityId:                "cLoUdCeRtIfIcAtIoNaUtHoRiTyId",
				CloudCertificationAuthorityLeafCertificateId: "cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy/cLoUdCeRtIfIcAtIoNaUtHoRiTyId/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtE/cLoUdCeRtIfIcAtIoNaUtHoRiTyLeAfCeRtIfIcAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudCertificationAuthorityId != v.Expected.CloudCertificationAuthorityId {
			t.Fatalf("Expected %q but got %q for CloudCertificationAuthorityId", v.Expected.CloudCertificationAuthorityId, actual.CloudCertificationAuthorityId)
		}

		if actual.CloudCertificationAuthorityLeafCertificateId != v.Expected.CloudCertificationAuthorityLeafCertificateId {
			t.Fatalf("Expected %q but got %q for CloudCertificationAuthorityLeafCertificateId", v.Expected.CloudCertificationAuthorityLeafCertificateId, actual.CloudCertificationAuthorityLeafCertificateId)
		}

	}
}

func TestSegmentsForDeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId(t *testing.T) {
	segments := DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCloudCertificationAuthorityIdCloudCertificationAuthorityLeafCertificateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
