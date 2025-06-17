package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudCertificationAuthorityId{}

func TestNewDeviceManagementCloudCertificationAuthorityID(t *testing.T) {
	id := NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId")

	if id.CloudCertificationAuthorityId != "cloudCertificationAuthorityId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudCertificationAuthorityId'", id.CloudCertificationAuthorityId, "cloudCertificationAuthorityId")
	}
}

func TestFormatDeviceManagementCloudCertificationAuthorityID(t *testing.T) {
	actual := NewDeviceManagementCloudCertificationAuthorityID("cloudCertificationAuthorityId").ID()
	expected := "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCloudCertificationAuthorityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudCertificationAuthorityId
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
			// Valid URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId",
			Expected: &DeviceManagementCloudCertificationAuthorityId{
				CloudCertificationAuthorityId: "cloudCertificationAuthorityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudCertificationAuthorityID(v.Input)
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

	}
}

func TestParseDeviceManagementCloudCertificationAuthorityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudCertificationAuthorityId
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
			// Valid URI
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId",
			Expected: &DeviceManagementCloudCertificationAuthorityId{
				CloudCertificationAuthorityId: "cloudCertificationAuthorityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudCertificationAuthority/cloudCertificationAuthorityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy/cLoUdCeRtIfIcAtIoNaUtHoRiTyId",
			Expected: &DeviceManagementCloudCertificationAuthorityId{
				CloudCertificationAuthorityId: "cLoUdCeRtIfIcAtIoNaUtHoRiTyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdCeRtIfIcAtIoNaUtHoRiTy/cLoUdCeRtIfIcAtIoNaUtHoRiTyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudCertificationAuthorityIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementCloudCertificationAuthorityId(t *testing.T) {
	segments := DeviceManagementCloudCertificationAuthorityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCloudCertificationAuthorityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
