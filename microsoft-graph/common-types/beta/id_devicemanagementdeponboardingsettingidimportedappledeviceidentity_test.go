package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{}

func TestNewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(t *testing.T) {
	id := NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingIdValue", "importedAppleDeviceIdentityIdValue")

	if id.DepOnboardingSettingId != "depOnboardingSettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DepOnboardingSettingId'", id.DepOnboardingSettingId, "depOnboardingSettingIdValue")
	}

	if id.ImportedAppleDeviceIdentityId != "importedAppleDeviceIdentityIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ImportedAppleDeviceIdentityId'", id.ImportedAppleDeviceIdentityId, "importedAppleDeviceIdentityIdValue")
	}
}

func TestFormatDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID(t *testing.T) {
	actual := NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingIdValue", "importedAppleDeviceIdentityIdValue").ID()
	expected := "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities/importedAppleDeviceIdentityIdValue"
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
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities/importedAppleDeviceIdentityIdValue",
			Expected: &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
				DepOnboardingSettingId:        "depOnboardingSettingIdValue",
				ImportedAppleDeviceIdentityId: "importedAppleDeviceIdentityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities/importedAppleDeviceIdentityIdValue/extra",
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
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/iMpOrTeDaPpLeDeViCeIdEnTiTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities/importedAppleDeviceIdentityIdValue",
			Expected: &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
				DepOnboardingSettingId:        "depOnboardingSettingIdValue",
				ImportedAppleDeviceIdentityId: "importedAppleDeviceIdentityIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/depOnboardingSettings/depOnboardingSettingIdValue/importedAppleDeviceIdentities/importedAppleDeviceIdentityIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/iMpOrTeDaPpLeDeViCeIdEnTiTiEs/iMpOrTeDaPpLeDeViCeIdEnTiTyIdVaLuE",
			Expected: &DeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityId{
				DepOnboardingSettingId:        "dEpOnBoArDiNgSeTtInGiDvAlUe",
				ImportedAppleDeviceIdentityId: "iMpOrTeDaPpLeDeViCeIdEnTiTyIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/dEpOnBoArDiNgSeTtInGs/dEpOnBoArDiNgSeTtInGiDvAlUe/iMpOrTeDaPpLeDeViCeIdEnTiTiEs/iMpOrTeDaPpLeDeViCeIdEnTiTyIdVaLuE/extra",
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
