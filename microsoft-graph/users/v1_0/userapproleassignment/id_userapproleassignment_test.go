package userapproleassignment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppRoleAssignmentId{}

func TestNewUserAppRoleAssignmentID(t *testing.T) {
	id := NewUserAppRoleAssignmentID("userIdValue", "appRoleAssignmentIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.AppRoleAssignmentId != "appRoleAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AppRoleAssignmentId'", id.AppRoleAssignmentId, "appRoleAssignmentIdValue")
	}
}

func TestFormatUserAppRoleAssignmentID(t *testing.T) {
	actual := NewUserAppRoleAssignmentID("userIdValue", "appRoleAssignmentIdValue").ID()
	expected := "/users/userIdValue/appRoleAssignments/appRoleAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserAppRoleAssignmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppRoleAssignmentId
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
			Input: "/users/userIdValue/appRoleAssignments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/appRoleAssignments/appRoleAssignmentIdValue",
			Expected: &UserAppRoleAssignmentId{
				UserId:              "userIdValue",
				AppRoleAssignmentId: "appRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appRoleAssignments/appRoleAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppRoleAssignmentID(v.Input)
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

		if actual.AppRoleAssignmentId != v.Expected.AppRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for AppRoleAssignmentId", v.Expected.AppRoleAssignmentId, actual.AppRoleAssignmentId)
		}

	}
}

func TestParseUserAppRoleAssignmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserAppRoleAssignmentId
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
			Input: "/users/userIdValue/appRoleAssignments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoLeAsSiGnMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/appRoleAssignments/appRoleAssignmentIdValue",
			Expected: &UserAppRoleAssignmentId{
				UserId:              "userIdValue",
				AppRoleAssignmentId: "appRoleAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/appRoleAssignments/appRoleAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoLeAsSiGnMeNtS/aPpRoLeAsSiGnMeNtIdVaLuE",
			Expected: &UserAppRoleAssignmentId{
				UserId:              "uSeRiDvAlUe",
				AppRoleAssignmentId: "aPpRoLeAsSiGnMeNtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoLeAsSiGnMeNtS/aPpRoLeAsSiGnMeNtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserAppRoleAssignmentIDInsensitively(v.Input)
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

		if actual.AppRoleAssignmentId != v.Expected.AppRoleAssignmentId {
			t.Fatalf("Expected %q but got %q for AppRoleAssignmentId", v.Expected.AppRoleAssignmentId, actual.AppRoleAssignmentId)
		}

	}
}

func TestSegmentsForUserAppRoleAssignmentId(t *testing.T) {
	segments := UserAppRoleAssignmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserAppRoleAssignmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
