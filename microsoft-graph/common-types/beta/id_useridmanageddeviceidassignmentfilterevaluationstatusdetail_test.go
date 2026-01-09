package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

func TestNewUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	id := NewUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID("userId", "managedDeviceId", "assignmentFilterEvaluationStatusDetailsId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.AssignmentFilterEvaluationStatusDetailsId != "assignmentFilterEvaluationStatusDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentFilterEvaluationStatusDetailsId'", id.AssignmentFilterEvaluationStatusDetailsId, "assignmentFilterEvaluationStatusDetailsId")
	}
}

func TestFormatUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	actual := NewUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID("userId", "managedDeviceId", "assignmentFilterEvaluationStatusDetailsId").ID()
	expected := "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
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
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId",
			Expected: &UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				UserId:          "userId",
				ManagedDeviceId: "managedDeviceId",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v.Input)
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

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.AssignmentFilterEvaluationStatusDetailsId != v.Expected.AssignmentFilterEvaluationStatusDetailsId {
			t.Fatalf("Expected %q but got %q for AssignmentFilterEvaluationStatusDetailsId", v.Expected.AssignmentFilterEvaluationStatusDetailsId, actual.AssignmentFilterEvaluationStatusDetailsId)
		}

	}
}

func TestParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
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
			Input: "/users/userId/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId",
			Expected: &UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				UserId:          "userId",
				ManagedDeviceId: "managedDeviceId",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId",
			Expected: &UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				UserId:          "uSeRiD",
				ManagedDeviceId: "mAnAgEdDeViCeId",
				AssignmentFilterEvaluationStatusDetailsId: "aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(v.Input)
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

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.AssignmentFilterEvaluationStatusDetailsId != v.Expected.AssignmentFilterEvaluationStatusDetailsId {
			t.Fatalf("Expected %q but got %q for AssignmentFilterEvaluationStatusDetailsId", v.Expected.AssignmentFilterEvaluationStatusDetailsId, actual.AssignmentFilterEvaluationStatusDetailsId)
		}

	}
}

func TestSegmentsForUserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId(t *testing.T) {
	segments := UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdManagedDeviceIdAssignmentFilterEvaluationStatusDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
