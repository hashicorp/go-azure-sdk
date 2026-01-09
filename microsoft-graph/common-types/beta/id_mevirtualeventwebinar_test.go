package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeVirtualEventWebinarId{}

func TestNewMeVirtualEventWebinarID(t *testing.T) {
	id := NewMeVirtualEventWebinarID("virtualEventWebinarId")

	if id.VirtualEventWebinarId != "virtualEventWebinarId" {
		t.Fatalf("Expected %q but got %q for Segment 'VirtualEventWebinarId'", id.VirtualEventWebinarId, "virtualEventWebinarId")
	}
}

func TestFormatMeVirtualEventWebinarID(t *testing.T) {
	actual := NewMeVirtualEventWebinarID("virtualEventWebinarId").ID()
	expected := "/me/virtualEvents/webinars/virtualEventWebinarId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeVirtualEventWebinarID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeVirtualEventWebinarId
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
			Input: "/me/virtualEvents",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/virtualEvents/webinars",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/virtualEvents/webinars/virtualEventWebinarId",
			Expected: &MeVirtualEventWebinarId{
				VirtualEventWebinarId: "virtualEventWebinarId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/virtualEvents/webinars/virtualEventWebinarId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeVirtualEventWebinarID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.VirtualEventWebinarId != v.Expected.VirtualEventWebinarId {
			t.Fatalf("Expected %q but got %q for VirtualEventWebinarId", v.Expected.VirtualEventWebinarId, actual.VirtualEventWebinarId)
		}

	}
}

func TestParseMeVirtualEventWebinarIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeVirtualEventWebinarId
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
			Input: "/me/virtualEvents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/vIrTuAlEvEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/virtualEvents/webinars",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/vIrTuAlEvEnTs/wEbInArS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/virtualEvents/webinars/virtualEventWebinarId",
			Expected: &MeVirtualEventWebinarId{
				VirtualEventWebinarId: "virtualEventWebinarId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/virtualEvents/webinars/virtualEventWebinarId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/vIrTuAlEvEnTs/wEbInArS/vIrTuAlEvEnTwEbInArId",
			Expected: &MeVirtualEventWebinarId{
				VirtualEventWebinarId: "vIrTuAlEvEnTwEbInArId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/vIrTuAlEvEnTs/wEbInArS/vIrTuAlEvEnTwEbInArId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeVirtualEventWebinarIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.VirtualEventWebinarId != v.Expected.VirtualEventWebinarId {
			t.Fatalf("Expected %q but got %q for VirtualEventWebinarId", v.Expected.VirtualEventWebinarId, actual.VirtualEventWebinarId)
		}

	}
}

func TestSegmentsForMeVirtualEventWebinarId(t *testing.T) {
	segments := MeVirtualEventWebinarId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeVirtualEventWebinarId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
