package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCloudClipboardItemId{}

func TestNewMeCloudClipboardItemID(t *testing.T) {
	id := NewMeCloudClipboardItemID("cloudClipboardItemId")

	if id.CloudClipboardItemId != "cloudClipboardItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudClipboardItemId'", id.CloudClipboardItemId, "cloudClipboardItemId")
	}
}

func TestFormatMeCloudClipboardItemID(t *testing.T) {
	actual := NewMeCloudClipboardItemID("cloudClipboardItemId").ID()
	expected := "/me/cloudClipboard/items/cloudClipboardItemId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCloudClipboardItemID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCloudClipboardItemId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/cloudClipboard",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/cloudClipboard/items",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/cloudClipboard/items/cloudClipboardItemId",
			Expected: &MeCloudClipboardItemId{
				CloudClipboardItemId: "cloudClipboardItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/cloudClipboard/items/cloudClipboardItemId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCloudClipboardItemID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudClipboardItemId != v.Expected.CloudClipboardItemId {
			t.Fatalf("Expected %q but got %q for CloudClipboardItemId", v.Expected.CloudClipboardItemId, actual.CloudClipboardItemId)
		}

	}
}

func TestParseMeCloudClipboardItemIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCloudClipboardItemId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/cloudClipboard",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdClIpBoArD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/cloudClipboard/items",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdClIpBoArD/iTeMs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/cloudClipboard/items/cloudClipboardItemId",
			Expected: &MeCloudClipboardItemId{
				CloudClipboardItemId: "cloudClipboardItemId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/cloudClipboard/items/cloudClipboardItemId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdClIpBoArD/iTeMs/cLoUdClIpBoArDiTeMiD",
			Expected: &MeCloudClipboardItemId{
				CloudClipboardItemId: "cLoUdClIpBoArDiTeMiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdClIpBoArD/iTeMs/cLoUdClIpBoArDiTeMiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCloudClipboardItemIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudClipboardItemId != v.Expected.CloudClipboardItemId {
			t.Fatalf("Expected %q but got %q for CloudClipboardItemId", v.Expected.CloudClipboardItemId, actual.CloudClipboardItemId)
		}

	}
}

func TestSegmentsForMeCloudClipboardItemId(t *testing.T) {
	segments := MeCloudClipboardItemId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCloudClipboardItemId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
