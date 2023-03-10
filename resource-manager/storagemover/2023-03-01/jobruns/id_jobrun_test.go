package jobruns

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = JobRunId{}

func TestNewJobRunID(t *testing.T) {
	id := NewJobRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverValue", "projectValue", "jobDefinitionValue", "jobRunValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.StorageMoverName != "storageMoverValue" {
		t.Fatalf("Expected %q but got %q for Segment 'StorageMoverName'", id.StorageMoverName, "storageMoverValue")
	}

	if id.ProjectName != "projectValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProjectName'", id.ProjectName, "projectValue")
	}

	if id.JobDefinitionName != "jobDefinitionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'JobDefinitionName'", id.JobDefinitionName, "jobDefinitionValue")
	}

	if id.JobRunName != "jobRunValue" {
		t.Fatalf("Expected %q but got %q for Segment 'JobRunName'", id.JobRunName, "jobRunValue")
	}
}

func TestFormatJobRunID(t *testing.T) {
	actual := NewJobRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageMoverValue", "projectValue", "jobDefinitionValue", "jobRunValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns/jobRunValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseJobRunID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *JobRunId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns/jobRunValue",
			Expected: &JobRunId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				StorageMoverName:  "storageMoverValue",
				ProjectName:       "projectValue",
				JobDefinitionName: "jobDefinitionValue",
				JobRunName:        "jobRunValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns/jobRunValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseJobRunID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.StorageMoverName != v.Expected.StorageMoverName {
			t.Fatalf("Expected %q but got %q for StorageMoverName", v.Expected.StorageMoverName, actual.StorageMoverName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

		if actual.JobDefinitionName != v.Expected.JobDefinitionName {
			t.Fatalf("Expected %q but got %q for JobDefinitionName", v.Expected.JobDefinitionName, actual.JobDefinitionName)
		}

		if actual.JobRunName != v.Expected.JobRunName {
			t.Fatalf("Expected %q but got %q for JobRunName", v.Expected.JobRunName, actual.JobRunName)
		}

	}
}

func TestParseJobRunIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *JobRunId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS/pRoJeCtVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS/pRoJeCtVaLuE/jObDeFiNiTiOnS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS/pRoJeCtVaLuE/jObDeFiNiTiOnS/jObDeFiNiTiOnVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS/pRoJeCtVaLuE/jObDeFiNiTiOnS/jObDeFiNiTiOnVaLuE/jObRuNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns/jobRunValue",
			Expected: &JobRunId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "example-resource-group",
				StorageMoverName:  "storageMoverValue",
				ProjectName:       "projectValue",
				JobDefinitionName: "jobDefinitionValue",
				JobRunName:        "jobRunValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StorageMover/storageMovers/storageMoverValue/projects/projectValue/jobDefinitions/jobDefinitionValue/jobRuns/jobRunValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS/pRoJeCtVaLuE/jObDeFiNiTiOnS/jObDeFiNiTiOnVaLuE/jObRuNs/jObRuNvAlUe",
			Expected: &JobRunId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroupName: "eXaMpLe-rEsOuRcE-GrOuP",
				StorageMoverName:  "sToRaGeMoVeRvAlUe",
				ProjectName:       "pRoJeCtVaLuE",
				JobDefinitionName: "jObDeFiNiTiOnVaLuE",
				JobRunName:        "jObRuNvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.sToRaGeMoVeR/sToRaGeMoVeRs/sToRaGeMoVeRvAlUe/pRoJeCtS/pRoJeCtVaLuE/jObDeFiNiTiOnS/jObDeFiNiTiOnVaLuE/jObRuNs/jObRuNvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseJobRunIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.StorageMoverName != v.Expected.StorageMoverName {
			t.Fatalf("Expected %q but got %q for StorageMoverName", v.Expected.StorageMoverName, actual.StorageMoverName)
		}

		if actual.ProjectName != v.Expected.ProjectName {
			t.Fatalf("Expected %q but got %q for ProjectName", v.Expected.ProjectName, actual.ProjectName)
		}

		if actual.JobDefinitionName != v.Expected.JobDefinitionName {
			t.Fatalf("Expected %q but got %q for JobDefinitionName", v.Expected.JobDefinitionName, actual.JobDefinitionName)
		}

		if actual.JobRunName != v.Expected.JobRunName {
			t.Fatalf("Expected %q but got %q for JobRunName", v.Expected.JobRunName, actual.JobRunName)
		}

	}
}

func TestSegmentsForJobRunId(t *testing.T) {
	segments := JobRunId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("JobRunId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
