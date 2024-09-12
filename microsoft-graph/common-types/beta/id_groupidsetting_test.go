package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSettingId{}

func TestNewGroupIdSettingID(t *testing.T) {
	id := NewGroupIdSettingID("groupIdValue", "directorySettingIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.DirectorySettingId != "directorySettingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DirectorySettingId'", id.DirectorySettingId, "directorySettingIdValue")
	}
}

func TestFormatGroupIdSettingID(t *testing.T) {
	actual := NewGroupIdSettingID("groupIdValue", "directorySettingIdValue").ID()
	expected := "/groups/groupIdValue/settings/directorySettingIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdSettingID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/settings",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/settings/directorySettingIdValue",
			Expected: &GroupIdSettingId{
				GroupId:            "groupIdValue",
				DirectorySettingId: "directorySettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/settings/directorySettingIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSettingID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DirectorySettingId != v.Expected.DirectorySettingId {
			t.Fatalf("Expected %q but got %q for DirectorySettingId", v.Expected.DirectorySettingId, actual.DirectorySettingId)
		}

	}
}

func TestParseGroupIdSettingIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdSettingId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/settings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sEtTiNgS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/settings/directorySettingIdValue",
			Expected: &GroupIdSettingId{
				GroupId:            "groupIdValue",
				DirectorySettingId: "directorySettingIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/settings/directorySettingIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sEtTiNgS/dIrEcToRySeTtInGiDvAlUe",
			Expected: &GroupIdSettingId{
				GroupId:            "gRoUpIdVaLuE",
				DirectorySettingId: "dIrEcToRySeTtInGiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/sEtTiNgS/dIrEcToRySeTtInGiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdSettingIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.DirectorySettingId != v.Expected.DirectorySettingId {
			t.Fatalf("Expected %q but got %q for DirectorySettingId", v.Expected.DirectorySettingId, actual.DirectorySettingId)
		}

	}
}

func TestSegmentsForGroupIdSettingId(t *testing.T) {
	segments := GroupIdSettingId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdSettingId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
