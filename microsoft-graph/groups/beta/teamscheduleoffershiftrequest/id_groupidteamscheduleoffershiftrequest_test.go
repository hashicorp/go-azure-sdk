package teamscheduleoffershiftrequest

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdTeamScheduleOfferShiftRequestId{}

func TestNewGroupIdTeamScheduleOfferShiftRequestID(t *testing.T) {
	id := NewGroupIdTeamScheduleOfferShiftRequestID("groupIdValue", "offerShiftRequestIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.OfferShiftRequestId != "offerShiftRequestIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OfferShiftRequestId'", id.OfferShiftRequestId, "offerShiftRequestIdValue")
	}
}

func TestFormatGroupIdTeamScheduleOfferShiftRequestID(t *testing.T) {
	actual := NewGroupIdTeamScheduleOfferShiftRequestID("groupIdValue", "offerShiftRequestIdValue").ID()
	expected := "/groups/groupIdValue/team/schedule/offerShiftRequests/offerShiftRequestIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupIdTeamScheduleOfferShiftRequestID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleOfferShiftRequestId
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
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule/offerShiftRequests",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/offerShiftRequests/offerShiftRequestIdValue",
			Expected: &GroupIdTeamScheduleOfferShiftRequestId{
				GroupId:             "groupIdValue",
				OfferShiftRequestId: "offerShiftRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/offerShiftRequests/offerShiftRequestIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleOfferShiftRequestID(v.Input)
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

		if actual.OfferShiftRequestId != v.Expected.OfferShiftRequestId {
			t.Fatalf("Expected %q but got %q for OfferShiftRequestId", v.Expected.OfferShiftRequestId, actual.OfferShiftRequestId)
		}

	}
}

func TestParseGroupIdTeamScheduleOfferShiftRequestIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupIdTeamScheduleOfferShiftRequestId
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
			Input: "/groups/groupIdValue/team",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/team/schedule/offerShiftRequests",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/oFfErShIfTrEqUeStS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/team/schedule/offerShiftRequests/offerShiftRequestIdValue",
			Expected: &GroupIdTeamScheduleOfferShiftRequestId{
				GroupId:             "groupIdValue",
				OfferShiftRequestId: "offerShiftRequestIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/team/schedule/offerShiftRequests/offerShiftRequestIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/oFfErShIfTrEqUeStS/oFfErShIfTrEqUeStIdVaLuE",
			Expected: &GroupIdTeamScheduleOfferShiftRequestId{
				GroupId:             "gRoUpIdVaLuE",
				OfferShiftRequestId: "oFfErShIfTrEqUeStIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/tEaM/sChEdUlE/oFfErShIfTrEqUeStS/oFfErShIfTrEqUeStIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupIdTeamScheduleOfferShiftRequestIDInsensitively(v.Input)
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

		if actual.OfferShiftRequestId != v.Expected.OfferShiftRequestId {
			t.Fatalf("Expected %q but got %q for OfferShiftRequestId", v.Expected.OfferShiftRequestId, actual.OfferShiftRequestId)
		}

	}
}

func TestSegmentsForGroupIdTeamScheduleOfferShiftRequestId(t *testing.T) {
	segments := GroupIdTeamScheduleOfferShiftRequestId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupIdTeamScheduleOfferShiftRequestId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
