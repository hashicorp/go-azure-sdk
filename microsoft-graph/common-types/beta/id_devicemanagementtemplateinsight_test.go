package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementTemplateInsightId{}

func TestNewDeviceManagementTemplateInsightID(t *testing.T) {
	id := NewDeviceManagementTemplateInsightID("deviceManagementTemplateInsightsDefinitionId")

	if id.DeviceManagementTemplateInsightsDefinitionId != "deviceManagementTemplateInsightsDefinitionId" {
		t.Fatalf("Expected %q but got %q for Segment 'DeviceManagementTemplateInsightsDefinitionId'", id.DeviceManagementTemplateInsightsDefinitionId, "deviceManagementTemplateInsightsDefinitionId")
	}
}

func TestFormatDeviceManagementTemplateInsightID(t *testing.T) {
	actual := NewDeviceManagementTemplateInsightID("deviceManagementTemplateInsightsDefinitionId").ID()
	expected := "/deviceManagement/templateInsights/deviceManagementTemplateInsightsDefinitionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementTemplateInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateInsightId
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
			Input: "/deviceManagement/templateInsights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templateInsights/deviceManagementTemplateInsightsDefinitionId",
			Expected: &DeviceManagementTemplateInsightId{
				DeviceManagementTemplateInsightsDefinitionId: "deviceManagementTemplateInsightsDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templateInsights/deviceManagementTemplateInsightsDefinitionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateInsightID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateInsightsDefinitionId != v.Expected.DeviceManagementTemplateInsightsDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateInsightsDefinitionId", v.Expected.DeviceManagementTemplateInsightsDefinitionId, actual.DeviceManagementTemplateInsightsDefinitionId)
		}

	}
}

func TestParseDeviceManagementTemplateInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementTemplateInsightId
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
			Input: "/deviceManagement/templateInsights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEiNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/templateInsights/deviceManagementTemplateInsightsDefinitionId",
			Expected: &DeviceManagementTemplateInsightId{
				DeviceManagementTemplateInsightsDefinitionId: "deviceManagementTemplateInsightsDefinitionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/templateInsights/deviceManagementTemplateInsightsDefinitionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEiNsIgHtS/dEvIcEmAnAgEmEnTtEmPlAtEiNsIgHtSdEfInItIoNiD",
			Expected: &DeviceManagementTemplateInsightId{
				DeviceManagementTemplateInsightsDefinitionId: "dEvIcEmAnAgEmEnTtEmPlAtEiNsIgHtSdEfInItIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/tEmPlAtEiNsIgHtS/dEvIcEmAnAgEmEnTtEmPlAtEiNsIgHtSdEfInItIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementTemplateInsightIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DeviceManagementTemplateInsightsDefinitionId != v.Expected.DeviceManagementTemplateInsightsDefinitionId {
			t.Fatalf("Expected %q but got %q for DeviceManagementTemplateInsightsDefinitionId", v.Expected.DeviceManagementTemplateInsightsDefinitionId, actual.DeviceManagementTemplateInsightsDefinitionId)
		}

	}
}

func TestSegmentsForDeviceManagementTemplateInsightId(t *testing.T) {
	segments := DeviceManagementTemplateInsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementTemplateInsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
