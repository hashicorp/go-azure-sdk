package useronlinemeetingattendancereportattendancerecord

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingAttendanceReportAttendanceRecordId{}

func TestNewUserOnlineMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	id := NewUserOnlineMeetingAttendanceReportAttendanceRecordID("userIdValue", "onlineMeetingIdValue", "meetingAttendanceReportIdValue", "attendanceRecordIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.OnlineMeetingId != "onlineMeetingIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OnlineMeetingId'", id.OnlineMeetingId, "onlineMeetingIdValue")
	}

	if id.MeetingAttendanceReportId != "meetingAttendanceReportIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MeetingAttendanceReportId'", id.MeetingAttendanceReportId, "meetingAttendanceReportIdValue")
	}

	if id.AttendanceRecordId != "attendanceRecordIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttendanceRecordId'", id.AttendanceRecordId, "attendanceRecordIdValue")
	}
}

func TestFormatUserOnlineMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	actual := NewUserOnlineMeetingAttendanceReportAttendanceRecordID("userIdValue", "onlineMeetingIdValue", "meetingAttendanceReportIdValue", "attendanceRecordIdValue").ID()
	expected := "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserOnlineMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingAttendanceReportAttendanceRecordId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue",
			Expected: &UserOnlineMeetingAttendanceReportAttendanceRecordId{
				UserId:                    "userIdValue",
				OnlineMeetingId:           "onlineMeetingIdValue",
				MeetingAttendanceReportId: "meetingAttendanceReportIdValue",
				AttendanceRecordId:        "attendanceRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingAttendanceReportAttendanceRecordID(v.Input)
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

func TestParseUserOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserOnlineMeetingAttendanceReportAttendanceRecordId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/aTtEnDaNcErEcOrDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue",
			Expected: &UserOnlineMeetingAttendanceReportAttendanceRecordId{
				UserId:                    "userIdValue",
				OnlineMeetingId:           "onlineMeetingIdValue",
				MeetingAttendanceReportId: "meetingAttendanceReportIdValue",
				AttendanceRecordId:        "attendanceRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiDvAlUe",
			Expected: &UserOnlineMeetingAttendanceReportAttendanceRecordId{
				UserId:                    "uSeRiDvAlUe",
				OnlineMeetingId:           "oNlInEmEeTiNgIdVaLuE",
				MeetingAttendanceReportId: "mEeTiNgAtTeNdAnCeRePoRtIdVaLuE",
				AttendanceRecordId:        "aTtEnDaNcErEcOrDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively(v.Input)
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

func TestSegmentsForUserOnlineMeetingAttendanceReportAttendanceRecordId(t *testing.T) {
	segments := UserOnlineMeetingAttendanceReportAttendanceRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserOnlineMeetingAttendanceReportAttendanceRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
