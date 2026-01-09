package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}

func TestNewMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(t *testing.T) {
	id := NewMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID("onlineMeetingId", "meetingAttendanceReportId", "attendanceRecordId")

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.MeetingAttendanceReportId != "meetingAttendanceReportId" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingAttendanceReportId'", id.MeetingAttendanceReportId, "meetingAttendanceReportId")
	}

	if id.AttendanceRecordId != "attendanceRecordId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttendanceRecordId'", id.AttendanceRecordId, "attendanceRecordId")
	}
}

func TestFormatMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(t *testing.T) {
	actual := NewMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID("onlineMeetingId", "meetingAttendanceReportId", "attendanceRecordId").ID()
	expected := "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId
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
			Input: "/me/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId",
			Expected: &MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
				OnlineMeetingId:           "onlineMeetingId",
				MeetingAttendanceReportId: "meetingAttendanceReportId",
				AttendanceRecordId:        "attendanceRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordID(v.Input)
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

		if actual.AttendanceRecordId != v.Expected.AttendanceRecordId {
			t.Fatalf("Expected %q but got %q for AttendanceRecordId", v.Expected.AttendanceRecordId, actual.AttendanceRecordId)
		}

	}
}

func TestParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId
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
			Input: "/me/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId/aTtEnDaNcErEcOrDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId",
			Expected: &MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
				OnlineMeetingId:           "onlineMeetingId",
				MeetingAttendanceReportId: "meetingAttendanceReportId",
				AttendanceRecordId:        "attendanceRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiD",
			Expected: &MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
				OnlineMeetingId:           "oNlInEmEeTiNgId",
				MeetingAttendanceReportId: "mEeTiNgAtTeNdAnCeRePoRtId",
				AttendanceRecordId:        "aTtEnDaNcErEcOrDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively(v.Input)
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

		if actual.AttendanceRecordId != v.Expected.AttendanceRecordId {
			t.Fatalf("Expected %q but got %q for AttendanceRecordId", v.Expected.AttendanceRecordId, actual.AttendanceRecordId)
		}

	}
}

func TestSegmentsForMeOnlineMeetingIdAttendanceReportIdAttendanceRecordId(t *testing.T) {
	segments := MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingIdAttendanceReportIdAttendanceRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
