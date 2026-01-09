package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}

func TestNewUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	id := NewUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID("userId", "onlineMeetingId", "attendanceRecordId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OnlineMeetingId != "onlineMeetingId" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingId")
	}

	if id.AttendanceRecordId != "attendanceRecordId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttendanceRecordId'", id.AttendanceRecordId, "attendanceRecordId")
	}
}

func TestFormatUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	actual := NewUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID("userId", "onlineMeetingId", "attendanceRecordId").ID()
	expected := "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords/attendanceRecordId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId
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
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords/attendanceRecordId",
			Expected: &UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{
				UserId:             "userId",
				OnlineMeetingId:    "onlineMeetingId",
				AttendanceRecordId: "attendanceRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords/attendanceRecordId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(v.Input)
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

		if actual.AttendanceRecordId != v.Expected.AttendanceRecordId {
			t.Fatalf("Expected %q but got %q for AttendanceRecordId", v.Expected.AttendanceRecordId, actual.AttendanceRecordId)
		}

	}
}

func TestParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId
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
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/mEeTiNgAtTeNdAnCeRePoRt",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/mEeTiNgAtTeNdAnCeRePoRt/aTtEnDaNcErEcOrDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords/attendanceRecordId",
			Expected: &UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{
				UserId:             "userId",
				OnlineMeetingId:    "onlineMeetingId",
				AttendanceRecordId: "attendanceRecordId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/onlineMeetings/onlineMeetingId/meetingAttendanceReport/attendanceRecords/attendanceRecordId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/mEeTiNgAtTeNdAnCeRePoRt/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiD",
			Expected: &UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{
				UserId:             "uSeRiD",
				OnlineMeetingId:    "oNlInEmEeTiNgId",
				AttendanceRecordId: "aTtEnDaNcErEcOrDiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oNlInEmEeTiNgS/oNlInEmEeTiNgId/mEeTiNgAtTeNdAnCeRePoRt/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordIDInsensitively(v.Input)
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

		if actual.AttendanceRecordId != v.Expected.AttendanceRecordId {
			t.Fatalf("Expected %q but got %q for AttendanceRecordId", v.Expected.AttendanceRecordId, actual.AttendanceRecordId)
		}

	}
}

func TestSegmentsForUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId(t *testing.T) {
	segments := UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
