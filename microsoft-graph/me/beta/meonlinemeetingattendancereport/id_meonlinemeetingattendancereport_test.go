package meonlinemeetingattendancereport

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingAttendanceReportId{}

func TestNewMeOnlineMeetingAttendanceReportID(t *testing.T) {
	id := NewMeOnlineMeetingAttendanceReportID("onlineMeetingIdValue", "meetingAttendanceReportIdValue")

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.MeetingAttendanceReportId != "meetingAttendanceReportIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingAttendanceReportId'", id.MeetingAttendanceReportId, "meetingAttendanceReportIdValue")
	}
}

func TestFormatMeOnlineMeetingAttendanceReportID(t *testing.T) {
	actual := NewMeOnlineMeetingAttendanceReportID("onlineMeetingIdValue", "meetingAttendanceReportIdValue").ID()
	expected := "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingAttendanceReportID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingAttendanceReportId
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
			Input: "/me/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue",
			Expected: &MeOnlineMeetingAttendanceReportId{
				OnlineMeetingId:           "onlineMeetingIdValue",
				MeetingAttendanceReportId: "meetingAttendanceReportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingAttendanceReportID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.MeetingAttendanceReportId != v.Expected.MeetingAttendanceReportId {
			t.Fatalf("Expected %q but got %q for MeetingAttendanceReportId", v.Expected.MeetingAttendanceReportId, actual.MeetingAttendanceReportId)
		}

	}
}

func TestParseMeOnlineMeetingAttendanceReportIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingAttendanceReportId
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
			Input: "/me/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue",
			Expected: &MeOnlineMeetingAttendanceReportId{
				OnlineMeetingId:           "onlineMeetingIdValue",
				MeetingAttendanceReportId: "meetingAttendanceReportIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE",
			Expected: &MeOnlineMeetingAttendanceReportId{
				OnlineMeetingId:           "oNlInEmEeTiNgIdVaLuE",
				MeetingAttendanceReportId: "mEeTiNgAtTeNdAnCeRePoRtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingAttendanceReportIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OnlineMeetingId != v.Expected.OnlineMeetingId {
			t.Fatalf("Expected %q but got %q for OnlineMeetingId", v.Expected.OnlineMeetingId, actual.OnlineMeetingId)
		}

		if actual.MeetingAttendanceReportId != v.Expected.MeetingAttendanceReportId {
			t.Fatalf("Expected %q but got %q for MeetingAttendanceReportId", v.Expected.MeetingAttendanceReportId, actual.MeetingAttendanceReportId)
		}

	}
}

func TestSegmentsForMeOnlineMeetingAttendanceReportId(t *testing.T) {
	segments := MeOnlineMeetingAttendanceReportId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingAttendanceReportId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
