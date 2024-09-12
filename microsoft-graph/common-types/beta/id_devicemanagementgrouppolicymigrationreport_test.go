package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyMigrationReportId{}

func TestNewDeviceManagementGroupPolicyMigrationReportID(t *testing.T) {
	id := NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportIdValue")

	if id.GroupPolicyMigrationReportId != "groupPolicyMigrationReportIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupPolicyMigrationReportId'", id.GroupPolicyMigrationReportId, "groupPolicyMigrationReportIdValue")
	}
}

func TestFormatDeviceManagementGroupPolicyMigrationReportID(t *testing.T) {
	actual := NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportIdValue").ID()
	expected := "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementGroupPolicyMigrationReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyMigrationReportId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue",
			Expected: &DeviceManagementGroupPolicyMigrationReportId{
				GroupPolicyMigrationReportId: "groupPolicyMigrationReportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyMigrationReportID(v.Input)
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

	}
}

func TestParseDeviceManagementGroupPolicyMigrationReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementGroupPolicyMigrationReportId
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
			// Valid URI
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue",
			Expected: &DeviceManagementGroupPolicyMigrationReportId{
				GroupPolicyMigrationReportId: "groupPolicyMigrationReportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/groupPolicyMigrationReports/groupPolicyMigrationReportIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe",
			Expected: &DeviceManagementGroupPolicyMigrationReportId{
				GroupPolicyMigrationReportId: "gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/gRoUpPoLiCyMiGrAtIoNrEpOrTs/gRoUpPoLiCyMiGrAtIoNrEpOrTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementGroupPolicyMigrationReportIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForDeviceManagementGroupPolicyMigrationReportId(t *testing.T) {
	segments := DeviceManagementGroupPolicyMigrationReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementGroupPolicyMigrationReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
