package memanageddeviceassignmentfilterevaluationstatusdetail

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{}

func TestNewMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	id := NewMeManagedDeviceAssignmentFilterEvaluationStatusDetailID("managedDeviceIdValue", "assignmentFilterEvaluationStatusDetailsIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.AssignmentFilterEvaluationStatusDetailsId != "assignmentFilterEvaluationStatusDetailsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentFilterEvaluationStatusDetailsId'", id.AssignmentFilterEvaluationStatusDetailsId, "assignmentFilterEvaluationStatusDetailsIdValue")
	}
}

func TestFormatMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	actual := NewMeManagedDeviceAssignmentFilterEvaluationStatusDetailID("managedDeviceIdValue", "assignmentFilterEvaluationStatusDetailsIdValue").ID()
	expected := "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceAssignmentFilterEvaluationStatusDetailId
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
			Input: "/me/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue",
			Expected: &MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceIdValue",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailID(v.Input)
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

func TestParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceAssignmentFilterEvaluationStatusDetailId
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
			Input: "/me/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue",
			Expected: &MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceIdValue",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsIdVaLuE",
			Expected: &MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "mAnAgEdDeViCeIdVaLuE",
				AssignmentFilterEvaluationStatusDetailsId: "aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceAssignmentFilterEvaluationStatusDetailIDInsensitively(v.Input)
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

func TestSegmentsForMeManagedDeviceAssignmentFilterEvaluationStatusDetailId(t *testing.T) {
	segments := MeManagedDeviceAssignmentFilterEvaluationStatusDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedDeviceAssignmentFilterEvaluationStatusDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
