package iotcentrals

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DashboardId{}

func TestNewDashboardID(t *testing.T) {
	id := NewDashboardID("https://endpoint-url.example.com", "dashboardId")

	if id.BaseURI != "https://endpoint-url.example.com" {
		t.Fatalf("Expected %q but got %q for Segment 'BaseURI'", id.BaseURI, "https://endpoint-url.example.com")
	}

	if id.DashboardId != "dashboardId" {
		t.Fatalf("Expected %q but got %q for Segment 'DashboardId'", id.DashboardId, "dashboardId")
	}
}

func TestFormatDashboardID(t *testing.T) {
	actual := NewDashboardID("https://endpoint-url.example.com", "dashboardId").ID()
	expected := "https://endpoint-url.example.com/dashboards/dashboardId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDashboardID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DashboardId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/dashboards",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dashboards/dashboardId",
			Expected: &DashboardId{
				BaseURI:     "https://endpoint-url.example.com",
				DashboardId: "dashboardId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dashboards/dashboardId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDashboardID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DashboardId != v.Expected.DashboardId {
			t.Fatalf("Expected %q but got %q for DashboardId", v.Expected.DashboardId, actual.DashboardId)
		}

	}
}

func TestParseDashboardIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DashboardId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "https://endpoint-url.example.com/dashboards",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAsHbOaRdS",
			Error: true,
		},
		{
			// Valid URI
			Input: "https://endpoint-url.example.com/dashboards/dashboardId",
			Expected: &DashboardId{
				BaseURI:     "https://endpoint-url.example.com",
				DashboardId: "dashboardId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "https://endpoint-url.example.com/dashboards/dashboardId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAsHbOaRdS/dAsHbOaRdId",
			Expected: &DashboardId{
				BaseURI:     "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM",
				DashboardId: "dAsHbOaRdId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "hTtPs://eNdPoInT-UrL.ExAmPlE.CoM/dAsHbOaRdS/dAsHbOaRdId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDashboardIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BaseURI != v.Expected.BaseURI {
			t.Fatalf("Expected %q but got %q for BaseURI", v.Expected.BaseURI, actual.BaseURI)
		}

		if actual.DashboardId != v.Expected.DashboardId {
			t.Fatalf("Expected %q but got %q for DashboardId", v.Expected.DashboardId, actual.DashboardId)
		}

	}
}

func TestSegmentsForDashboardId(t *testing.T) {
	segments := DashboardId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DashboardId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
