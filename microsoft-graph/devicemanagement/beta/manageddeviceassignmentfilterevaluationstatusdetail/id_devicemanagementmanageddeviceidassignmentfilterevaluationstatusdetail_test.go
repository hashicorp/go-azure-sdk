package manageddeviceassignmentfilterevaluationstatusdetail

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}

func TestNewDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	id := NewDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID("managedDeviceIdValue", "assignmentFilterEvaluationStatusDetailsIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.AssignmentFilterEvaluationStatusDetailsId != "assignmentFilterEvaluationStatusDetailsIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AssignmentFilterEvaluationStatusDetailsId'", id.AssignmentFilterEvaluationStatusDetailsId, "assignmentFilterEvaluationStatusDetailsIdValue")
	}
}

func TestFormatDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	actual := NewDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID("managedDeviceIdValue", "assignmentFilterEvaluationStatusDetailsIdValue").ID()
	expected := "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue",
			Expected: &DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceIdValue",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailID(v.Input)
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

func TestParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId
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
			Input: "/deviceManagement/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue",
			Expected: &DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "managedDeviceIdValue",
				AssignmentFilterEvaluationStatusDetailsId: "assignmentFilterEvaluationStatusDetailsIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/deviceManagement/managedDevices/managedDeviceIdValue/assignmentFilterEvaluationStatusDetails/assignmentFilterEvaluationStatusDetailsIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsIdVaLuE",
			Expected: &DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{
				ManagedDeviceId: "mAnAgEdDeViCeIdVaLuE",
				AssignmentFilterEvaluationStatusDetailsId: "aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/dEvIcEmAnAgEmEnT/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLs/aSsIgNmEnTfIlTeReVaLuAtIoNsTaTuSdEtAiLsIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailIDInsensitively(v.Input)
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

func TestSegmentsForDeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId(t *testing.T) {
	segments := DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
