package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{}

func TestNewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(t *testing.T) {
	id := NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingId", "importedAppleDeviceIdentityId")

	if id.DepOnboardingSettingId != "depOnboardingSettingId" {
		t.Fatalf("Expected %q but got %q for Segment 'DepOnboardingSettingId'", id.DepOnboardingSettingId, "depOnboardingSettingId")
	}

	if id.ImportedAppleDeviceIdentityId != "importedAppleDeviceIdentityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ImportedAppleDeviceIdentityId'", id.ImportedAppleDeviceIdentityId, "importedAppleDeviceIdentityId")
	}
}

func TestFormatDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(t *testing.T) {
	actual := NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingId", "importedAppleDeviceIdentityId").ID()
	expected := "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities/importedAppleDeviceIdentityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId
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
			Input: "/deviceManagement/depOnboardingSettings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities/importedAppleDeviceIdentityId",
			Expected: &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
				DepOnboardingSettingId:        "depOnboardingSettingId",
				ImportedAppleDeviceIdentityId: "importedAppleDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities/importedAppleDeviceIdentityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DepOnboardingSettingId != v.Expected.DepOnboardingSettingId {
			t.Fatalf("Expected %q but got %q for DepOnboardingSettingId", v.Expected.DepOnboardingSettingId, actual.DepOnboardingSettingId)
		}

		if actual.ImportedAppleDeviceIdentityId != v.Expected.ImportedAppleDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for ImportedAppleDeviceIdentityId", v.Expected.ImportedAppleDeviceIdentityId, actual.ImportedAppleDeviceIdentityId)
		}

	}
}

func TestParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId
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
			Input: "/deviceManagement/depOnboardingSettings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD/iMpOrTeDaPpLeDeViCeIdEnTiTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities/importedAppleDeviceIdentityId",
			Expected: &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
				DepOnboardingSettingId:        "depOnboardingSettingId",
				ImportedAppleDeviceIdentityId: "importedAppleDeviceIdentityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingId/importedAppleDeviceIdentities/importedAppleDeviceIdentityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD/iMpOrTeDaPpLeDeViCeIdEnTiTiEs/iMpOrTeDaPpLeDeViCeIdEnTiTyId",
			Expected: &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
				DepOnboardingSettingId:        "dEpOnBoArDiNgSeTtInGiD",
				ImportedAppleDeviceIdentityId: "iMpOrTeDaPpLeDeViCeIdEnTiTyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiD/iMpOrTeDaPpLeDeViCeIdEnTiTiEs/iMpOrTeDaPpLeDeViCeIdEnTiTyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DepOnboardingSettingId != v.Expected.DepOnboardingSettingId {
			t.Fatalf("Expected %q but got %q for DepOnboardingSettingId", v.Expected.DepOnboardingSettingId, actual.DepOnboardingSettingId)
		}

		if actual.ImportedAppleDeviceIdentityId != v.Expected.ImportedAppleDeviceIdentityId {
			t.Fatalf("Expected %q but got %q for ImportedAppleDeviceIdentityId", v.Expected.ImportedAppleDeviceIdentityId, actual.ImportedAppleDeviceIdentityId)
		}

	}
}

func TestSegmentsForDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId(t *testing.T) {
	segments := DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
