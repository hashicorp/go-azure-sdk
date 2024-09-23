package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

func TestNewDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	id := NewDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID("managedDeviceId", "assignmentFilterEvaluationStatusDetailsId")

	if id.ManagedDeviceId != "managedDeviceId" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceId")
	}

	if id.AssignmentFilterEvaluationStatusDetailsId != "assignmentFilterEvaluationStatusDetailsId" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentFilterEvaluationStatusDetailsId'", id.AssignmentFilterEvaluationStatusDetailsId, "assignmentFilterEvaluationStatusDetailsId")
	}
}

func TestFormatDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	actual := NewDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID("managedDeviceId", "assignmentFilterEvaluationStatusDetailsId").ID()
	expected := "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId",
			Expected: &DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceId",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v.Input)
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

func TestParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId",
			Expected: &DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceId",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/comanagedDevices/managedDeviceId/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId",
			Expected: &DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "mAnAgEdDeViCeId",
				AssignmentFilterEvaluationStatusDetailsId: "aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/cOmAnAgEdDeViCeS/mAnAgEdDeViCeId/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId(t *testing.T) {
	segments := DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementComanagedDeviceIdAssignmentFilterEvaluationStatusDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
