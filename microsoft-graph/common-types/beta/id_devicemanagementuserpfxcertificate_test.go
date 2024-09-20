package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserPfxCertificateId{}

func TestNewDeviceManagementUserPfxCertificateID(t *testing.T) {
	id := NewDeviceManagementUserPfxCertificateID("userPFXCertificateId")

	if id.UserPFXCertificateId != "userPFXCertificateId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserPFXCertificateId'", id.UserPFXCertificateId, "userPFXCertificateId")
	}
}

func TestFormatDeviceManagementUserPfxCertificateID(t *testing.T) {
	actual := NewDeviceManagementUserPfxCertificateID("userPFXCertificateId").ID()
	expected := "/deviceManagement/userPfxCertificates/userPFXCertificateId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserPfxCertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserPfxCertificateId
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
			Input: "/deviceManagement/userPfxCertificates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userPfxCertificates/userPFXCertificateId",
			Expected: &DeviceManagementUserPfxCertificateId{
				UserPFXCertificateId: "userPFXCertificateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userPfxCertificates/userPFXCertificateId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserPfxCertificateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserPFXCertificateId != v.Expected.UserPFXCertificateId {
			t.Fatalf("Expected %q but got %q for UserPFXCertificateId", v.Expected.UserPFXCertificateId, actual.UserPFXCertificateId)
		}

	}
}

func TestParseDeviceManagementUserPfxCertificateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserPfxCertificateId
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
			Input: "/deviceManagement/userPfxCertificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeRpFxCeRtIfIcAtEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userPfxCertificates/userPFXCertificateId",
			Expected: &DeviceManagementUserPfxCertificateId{
				UserPFXCertificateId: "userPFXCertificateId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userPfxCertificates/userPFXCertificateId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeRpFxCeRtIfIcAtEs/uSeRpFxCeRtIfIcAtEiD",
			Expected: &DeviceManagementUserPfxCertificateId{
				UserPFXCertificateId: "uSeRpFxCeRtIfIcAtEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeRpFxCeRtIfIcAtEs/uSeRpFxCeRtIfIcAtEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserPfxCertificateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserPFXCertificateId != v.Expected.UserPFXCertificateId {
			t.Fatalf("Expected %q but got %q for UserPFXCertificateId", v.Expected.UserPFXCertificateId, actual.UserPFXCertificateId)
		}

	}
}

func TestSegmentsForDeviceManagementUserPfxCertificateId(t *testing.T) {
	segments := DeviceManagementUserPfxCertificateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserPfxCertificateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
