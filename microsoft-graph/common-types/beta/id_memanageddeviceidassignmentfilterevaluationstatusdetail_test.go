package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

func TestNewMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	id := NewMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID("managedDeviceId", "assignmentFilterEvaluationStatusDetailsId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.AssignmentFilterEvaluationStatusDetailsId != "assignmentFilterEvaluationStatusDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentFilterEvaluationStatusDetailsId'", id.AssignmentFilterEvaluationStatusDetailsId, "assignmentFilterEvaluationStatusDetailsId")
	}
}

func TestFormatMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	actual := NewMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID("managedDeviceId", "assignmentFilterEvaluationStatusDetailsId").ID()
	expected := "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
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
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId",
			Expected: &MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceId",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.AssignmentFilterEvaluationStatusDetailsId != v.Expected.AssignmentFilterEvaluationStatusDetailsId {
			t.Fatalf("Expected %q but got %q for AssignmentFilterEvaluationStatusDetailsId", v.Expected.AssignmentFilterEvaluationStatusDetailsId, actual.AssignmentFilterEvaluationStatusDetailsId)
		}

	}
}

func TestParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
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
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId",
			Expected: &MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceId",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId",
			Expected: &MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "mAnAgEdDeViCeId",
				AssignmentFilterEvaluationStatusDetailsId: "aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.AssignmentFilterEvaluationStatusDetailsId != v.Expected.AssignmentFilterEvaluationStatusDetailsId {
			t.Fatalf("Expected %q but got %q for AssignmentFilterEvaluationStatusDetailsId", v.Expected.AssignmentFilterEvaluationStatusDetailsId, actual.AssignmentFilterEvaluationStatusDetailsId)
		}

	}
}

func TestSegmentsForMeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId(t *testing.T) {
	segments := MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedDeviceIdAssignmentFilterEvaluationStatusDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
