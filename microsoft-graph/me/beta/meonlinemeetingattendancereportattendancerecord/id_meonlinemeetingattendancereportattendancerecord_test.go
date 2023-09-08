package meonlinemeetingattendancereportattendancerecord

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingAttendanceReportAttendanceRecordId{}

func TestNewMeOnlineMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	id := NewMeOnlineMeetingAttendanceReportAttendanceRecordID("onlineMeetingIdValue", "meetingAttendanceReportIdValue", "attendanceRecordIdValue")

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

func TestFormatMeOnlineMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	actual := NewMeOnlineMeetingAttendanceReportAttendanceRecordID("onlineMeetingIdValue", "meetingAttendanceReportIdValue", "attendanceRecordIdValue").ID()
	expected := "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOnlineMeetingAttendanceReportAttendanceRecordID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingAttendanceReportAttendanceRecordId
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
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue",
			Expected: &MeOnlineMeetingAttendanceReportAttendanceRecordId{
				OnlineMeetingId:           "onlineMeetingIdValue",
				MeetingAttendanceReportId: "meetingAttendanceReportIdValue",
				AttendanceRecordId:        "attendanceRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingAttendanceReportAttendanceRecordID(v.Input)
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

func TestParseMeOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOnlineMeetingAttendanceReportAttendanceRecordId
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
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/aTtEnDaNcErEcOrDs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue",
			Expected: &MeOnlineMeetingAttendanceReportAttendanceRecordId{
				OnlineMeetingId:           "onlineMeetingIdValue",
				MeetingAttendanceReportId: "meetingAttendanceReportIdValue",
				AttendanceRecordId:        "attendanceRecordIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/onlineMeetings/onlineMeetingIdValue/attendanceReports/meetingAttendanceReportIdValue/attendanceRecords/attendanceRecordIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiDvAlUe",
			Expected: &MeOnlineMeetingAttendanceReportAttendanceRecordId{
				OnlineMeetingId:           "oNlInEmEeTiNgIdVaLuE",
				MeetingAttendanceReportId: "mEeTiNgAtTeNdAnCeRePoRtIdVaLuE",
				AttendanceRecordId:        "aTtEnDaNcErEcOrDiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oNlInEmEeTiNgS/oNlInEmEeTiNgIdVaLuE/aTtEnDaNcErEpOrTs/mEeTiNgAtTeNdAnCeRePoRtIdVaLuE/aTtEnDaNcErEcOrDs/aTtEnDaNcErEcOrDiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOnlineMeetingAttendanceReportAttendanceRecordIDInsensitively(v.Input)
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

func TestSegmentsForMeOnlineMeetingAttendanceReportAttendanceRecordId(t *testing.T) {
	segments := MeOnlineMeetingAttendanceReportAttendanceRecordId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOnlineMeetingAttendanceReportAttendanceRecordId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
