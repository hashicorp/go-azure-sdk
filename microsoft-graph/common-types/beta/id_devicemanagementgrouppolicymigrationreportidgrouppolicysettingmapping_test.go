package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{}

func TestNewDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID("groupPolicyMigrationReportId", "groupPolicySettingMappingId")

	if id.GroupPolicyMigrationReportId != "groupPolicyMigrationReportId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyMigrationReportId'", id.GroupPolicyMigrationReportId, "groupPolicyMigrationReportId")
	}

	if id.GroupPolicySettingMappingId != "groupPolicySettingMappingId" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicySettingMappingId'", id.GroupPolicySettingMappingId, "groupPolicySettingMappingId")
	}
}

func TestFormatDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID("groupPolicyMigrationReportId", "groupPolicySettingMappingId").ID()
	expected := "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings/groupPolicySettingMappingId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId
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
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings/groupPolicySettingMappingId",
			Expected: &DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{
				GroupPolicyMigrationReportId: "groupPolicyMigrationReportId",
				GroupPolicySettingMappingId:  "groupPolicySettingMappingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings/groupPolicySettingMappingId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingID(v.Input)
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

		if actual.GroupPolicySettingMappingId != v.Expected.GroupPolicySettingMappingId {
			t.Fatalf("Expected %q but got %q for GroupPolicySettingMappingId", v.Expected.GroupPolicySettingMappingId, actual.GroupPolicySettingMappingId)
		}

	}
}

func TestParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId
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
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiD/gRoUpPoLiCySeTtInGmApPiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings/groupPolicySettingMappingId",
			Expected: &DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{
				GroupPolicyMigrationReportId: "groupPolicyMigrationReportId",
				GroupPolicySettingMappingId:  "groupPolicySettingMappingId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportId/groupPolicySettingMappings/groupPolicySettingMappingId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiD/gRoUpPoLiCySeTtInGmApPiNgS/gRoUpPoLiCySeTtInGmApPiNgId",
			Expected: &DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{
				GroupPolicyMigrationReportId: "gRoUpPoLiCyMiGrAtIoNrEpOrTiD",
				GroupPolicySettingMappingId:  "gRoUpPoLiCySeTtInGmApPiNgId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiD/gRoUpPoLiCySeTtInGmApPiNgS/gRoUpPoLiCySeTtInGmApPiNgId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingIDInsensitively(v.Input)
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

		if actual.GroupPolicySettingMappingId != v.Expected.GroupPolicySettingMappingId {
			t.Fatalf("Expected %q but got %q for GroupPolicySettingMappingId", v.Expected.GroupPolicySettingMappingId, actual.GroupPolicySettingMappingId)
		}

	}
}

func TestSegmentsForDeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId(t *testing.T) {
	segments := DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
