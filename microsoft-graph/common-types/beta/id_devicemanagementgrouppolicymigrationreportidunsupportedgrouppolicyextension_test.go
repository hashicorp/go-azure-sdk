package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{}

func TestNewDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID("groupPolicyMigrationReportIdValue", "unsupportedGroupPolicyExtensionIdValue")

	if id.GroupPolicyMigrationReportId != "groupPolicyMigrationReportIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyMigrationReportId'", id.GroupPolicyMigrationReportId, "groupPolicyMigrationReportIdValue")
	}

	if id.UnsupportedGroupPolicyExtensionId != "unsupportedGroupPolicyExtensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UnsupportedGroupPolicyExtensionId'", id.UnsupportedGroupPolicyExtensionId, "unsupportedGroupPolicyExtensionIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID("groupPolicyMigrationReportIdValue", "unsupportedGroupPolicyExtensionIdValue").ID()
	expected := "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions/unsupportedGroupPolicyExtensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId
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
			Input: "/deviceManagement/groupPolicyMigrationReports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions/unsupportedGroupPolicyExtensionIdValue",
			Expected: &DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{
				GroupPolicyMigrationReportId:      "groupPolicyMigrationReportIdValue",
				UnsupportedGroupPolicyExtensionId: "unsupportedGroupPolicyExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions/unsupportedGroupPolicyExtensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyMigrationReportId != v.Expected.GroupPolicyMigrationReportId {
			t.Fatalf("Expected %q but got %q for GroupPolicyMigrationReportId", v.Expected.GroupPolicyMigrationReportId, actual.GroupPolicyMigrationReportId)
		}

		if actual.UnsupportedGroupPolicyExtensionId != v.Expected.UnsupportedGroupPolicyExtensionId {
			t.Fatalf("Expected %q but got %q for UnsupportedGroupPolicyExtensionId", v.Expected.UnsupportedGroupPolicyExtensionId, actual.UnsupportedGroupPolicyExtensionId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId
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
			Input: "/deviceManagement/groupPolicyMigrationReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe/uNsUpPoRtEdGrOuPpOlIcYeXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions/unsupportedGroupPolicyExtensionIdValue",
			Expected: &DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{
				GroupPolicyMigrationReportId:      "groupPolicyMigrationReportIdValue",
				UnsupportedGroupPolicyExtensionId: "unsupportedGroupPolicyExtensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/unsupportedGroupPolicyExtensions/unsupportedGroupPolicyExtensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe/uNsUpPoRtEdGrOuPpOlIcYeXtEnSiOnS/uNsUpPoRtEdGrOuPpOlIcYeXtEnSiOnIdVaLuE",
			Expected: &DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{
				GroupPolicyMigrationReportId:      "gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe",
				UnsupportedGroupPolicyExtensionId: "uNsUpPoRtEdGrOuPpOlIcYeXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe/uNsUpPoRtEdGrOuPpOlIcYeXtEnSiOnS/uNsUpPoRtEdGrOuPpOlIcYeXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupPolicyMigrationReportId != v.Expected.GroupPolicyMigrationReportId {
			t.Fatalf("Expected %q but got %q for GroupPolicyMigrationReportId", v.Expected.GroupPolicyMigrationReportId, actual.GroupPolicyMigrationReportId)
		}

		if actual.UnsupportedGroupPolicyExtensionId != v.Expected.UnsupportedGroupPolicyExtensionId {
			t.Fatalf("Expected %q but got %q for UnsupportedGroupPolicyExtensionId", v.Expected.UnsupportedGroupPolicyExtensionId, actual.UnsupportedGroupPolicyExtensionId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId(t *testing.T) {
	segments := DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyMigrationReportIdUnsupportedGroupPolicyExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
