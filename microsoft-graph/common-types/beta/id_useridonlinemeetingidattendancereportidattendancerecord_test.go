package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}

func TestNewUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID("userId", "onlineMeetingId", "meetingAttendanceReportId", "attendanceRecordId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

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

func TestFormatUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID("userId", "onlineMeetingId", "meetingAttendanceReportId", "attendanceRecordId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId",
			Expected: &UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
				UserId:                    "userId",
				OnlineMeetingId:           "onlineMeetingId",
				MeetingAttendanceReportId: "meetingAttendanceReportId",
				AttendanceRecordId:        "attendanceRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
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

func TestParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId/aTtEnDaNcErEcOrDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId",
			Expected: &UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
				UserId:                    "userId",
				OnlineMeetingId:           "onlineMeetingId",
				MeetingAttendanceReportId: "meetingAttendanceReportId",
				AttendanceRecordId:        "attendanceRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/attendanceReports/meetingAttendanceReportId/attendanceRecords/attendanceRecordId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiD",
			Expected: &UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{
				UserId:                    "uSeRiD",
				OnlineMeetingId:           "oNlInEmEeTiNgId",
				MeetingAttendanceReportId: "mEeTiNgAtTeNdAnCeRePoRtId",
				AttendanceRecordId:        "aTtEnDaNcErEcOrDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtId/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
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

func TestSegmentsForUserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId(t *testing.T) {
	segments := UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdAttendanceReportIdAttendanceRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
