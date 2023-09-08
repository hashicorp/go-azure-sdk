package userapprovalstep

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserApprovalStepId{}

func TestNewUserApprovalStepID(t *testing.T) {
	id := NewUserApprovalStepID("userIdValue", "approvalIdValue", "approvalStepIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ApprovalId != "approvalIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalId'", id.ApprovalId, "approvalIdValue")
	}

	if id.ApprovalStepId != "approvalStepIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApprovalStepId'", id.ApprovalStepId, "approvalStepIdValue")
	}
}

func TestFormatUserApprovalStepID(t *testing.T) {
	actual := NewUserApprovalStepID("userIdValue", "approvalIdValue", "approvalStepIdValue").ID()
	expected := "/users/userIdValue/approvals/approvalIdValue/steps/approvalStepIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserApprovalStepID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserApprovalStepId
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
			Input: "/users/userIdValue/approvals",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/approvals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/approvals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/approvals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &UserApprovalStepId{
				UserId:         "userIdValue",
				ApprovalId:     "approvalIdValue",
				ApprovalStepId: "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/approvals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserApprovalStepID(v.Input)
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

		if actual.ApprovalId != v.Expected.ApprovalId {
			t.Fatalf("Expected %q but got %q for ApprovalId", v.Expected.ApprovalId, actual.ApprovalId)
		}

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestParseUserApprovalStepIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserApprovalStepId
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
			Input: "/users/userIdValue/approvals",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoVaLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/approvals/approvalIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoVaLs/aPpRoVaLiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/approvals/approvalIdValue/steps",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoVaLs/aPpRoVaLiDvAlUe/sTePs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/approvals/approvalIdValue/steps/approvalStepIdValue",
			Expected: &UserApprovalStepId{
				UserId:         "userIdValue",
				ApprovalId:     "approvalIdValue",
				ApprovalStepId: "approvalStepIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/approvals/approvalIdValue/steps/approvalStepIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe",
			Expected: &UserApprovalStepId{
				UserId:         "uSeRiDvAlUe",
				ApprovalId:     "aPpRoVaLiDvAlUe",
				ApprovalStepId: "aPpRoVaLsTePiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/aPpRoVaLs/aPpRoVaLiDvAlUe/sTePs/aPpRoVaLsTePiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserApprovalStepIDInsensitively(v.Input)
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

		if actual.ApprovalId != v.Expected.ApprovalId {
			t.Fatalf("Expected %q but got %q for ApprovalId", v.Expected.ApprovalId, actual.ApprovalId)
		}

		if actual.ApprovalStepId != v.Expected.ApprovalStepId {
			t.Fatalf("Expected %q but got %q for ApprovalStepId", v.Expected.ApprovalStepId, actual.ApprovalStepId)
		}

	}
}

func TestSegmentsForUserApprovalStepId(t *testing.T) {
	segments := UserApprovalStepId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserApprovalStepId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
