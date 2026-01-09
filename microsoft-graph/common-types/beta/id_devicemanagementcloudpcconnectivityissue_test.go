package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementCloudPCConnectivityIssueId{}

func TestNewDeviceManagementCloudPCConnectivityIssueID(t *testing.T) {
	id := NewDeviceManagementCloudPCConnectivityIssueID("cloudPCConnectivityIssueId")

	if id.CloudPCConnectivityIssueId != "cloudPCConnectivityIssueId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCConnectivityIssueId'", id.CloudPCConnectivityIssueId, "cloudPCConnectivityIssueId")
	}
}

func TestFormatDeviceManagementCloudPCConnectivityIssueID(t *testing.T) {
	actual := NewDeviceManagementCloudPCConnectivityIssueID("cloudPCConnectivityIssueId").ID()
	expected := "/deviceManagement/cloudPCConnectivityIssues/cloudPCConnectivityIssueId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementCloudPCConnectivityIssueID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudPCConnectivityIssueId
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
			Input: "/deviceManagement/cloudPCConnectivityIssues",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cloudPCConnectivityIssues/cloudPCConnectivityIssueId",
			Expected: &DeviceManagementCloudPCConnectivityIssueId{
				CloudPCConnectivityIssueId: "cloudPCConnectivityIssueId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudPCConnectivityIssues/cloudPCConnectivityIssueId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudPCConnectivityIssueID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCConnectivityIssueId != v.Expected.CloudPCConnectivityIssueId {
			t.Fatalf("Expected %q but got %q for CloudPCConnectivityIssueId", v.Expected.CloudPCConnectivityIssueId, actual.CloudPCConnectivityIssueId)
		}

	}
}

func TestParseDeviceManagementCloudPCConnectivityIssueIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementCloudPCConnectivityIssueId
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
			Input: "/deviceManagement/cloudPCConnectivityIssues",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdPcCoNnEcTiViTyIsSuEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/cloudPCConnectivityIssues/cloudPCConnectivityIssueId",
			Expected: &DeviceManagementCloudPCConnectivityIssueId{
				CloudPCConnectivityIssueId: "cloudPCConnectivityIssueId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/cloudPCConnectivityIssues/cloudPCConnectivityIssueId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdPcCoNnEcTiViTyIsSuEs/cLoUdPcCoNnEcTiViTyIsSuEiD",
			Expected: &DeviceManagementCloudPCConnectivityIssueId{
				CloudPCConnectivityIssueId: "cLoUdPcCoNnEcTiViTyIsSuEiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cLoUdPcCoNnEcTiViTyIsSuEs/cLoUdPcCoNnEcTiViTyIsSuEiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementCloudPCConnectivityIssueIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCConnectivityIssueId != v.Expected.CloudPCConnectivityIssueId {
			t.Fatalf("Expected %q but got %q for CloudPCConnectivityIssueId", v.Expected.CloudPCConnectivityIssueId, actual.CloudPCConnectivityIssueId)
		}

	}
}

func TestSegmentsForDeviceManagementCloudPCConnectivityIssueId(t *testing.T) {
	segments := DeviceManagementCloudPCConnectivityIssueId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementCloudPCConnectivityIssueId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
