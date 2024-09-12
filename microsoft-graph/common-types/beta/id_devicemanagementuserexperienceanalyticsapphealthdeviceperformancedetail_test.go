package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{}

func TestNewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(t *testing.T) {
	id := NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID("userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue")

	if id.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId != "userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId'", id.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId, "userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue")
	}
}

func TestFormatDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(t *testing.T) {
	actual := NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID("userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue").ID()
	expected := "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails/userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails/userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{
				UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId: "userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails/userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId != v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId", v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId, actual.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId)
		}

	}
}

func TestParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId
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
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeDeTaIlS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails/userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{
				UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId: "userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/userExperienceAnalyticsAppHealthDevicePerformanceDetails/userExperienceAnalyticsAppHealthDevicePerformanceDetailsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeDeTaIlS/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeDeTaIlSiDvAlUe",
			Expected: &DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{
				UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId: "uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeDeTaIlSiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeDeTaIlS/uSeReXpErIeNcEaNaLyTiCsApPhEaLtHdEvIcEpErFoRmAnCeDeTaIlSiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId != v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId {
			t.Fatalf("Expected %q but got %q for UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId", v.Expected.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId, actual.UserExperienceAnalyticsAppHealthDevicePerformanceDetailsId)
		}

	}
}

func TestSegmentsForDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId(t *testing.T) {
	segments := DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
