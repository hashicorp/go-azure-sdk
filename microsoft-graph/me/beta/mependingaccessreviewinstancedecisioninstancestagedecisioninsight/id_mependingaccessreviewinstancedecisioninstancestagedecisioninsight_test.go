package mependingaccessreviewinstancedecisioninstancestagedecisioninsight

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}

func TestNewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(t *testing.T) {
	id := NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value", "governanceInsightIdValue")

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemIdValue")
	}

	if id.AccessReviewStageId != "accessReviewStageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewStageId'", id.AccessReviewStageId, "accessReviewStageIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId1 != "accessReviewInstanceDecisionItemId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId1'", id.AccessReviewInstanceDecisionItemId1, "accessReviewInstanceDecisionItemId1Value")
	}

	if id.GovernanceInsightId != "governanceInsightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GovernanceInsightId'", id.GovernanceInsightId, "governanceInsightIdValue")
	}
}

func TestFormatMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(t *testing.T) {
	actual := NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemId1Value", "governanceInsightIdValue").ID()
	expected := "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId
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
			Input: "/me/pendingAccessReviewInstances",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue",
			Expected: &MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{
				AccessReviewInstanceId:              "accessReviewInstanceIdValue",
				AccessReviewInstanceDecisionItemId:  "accessReviewInstanceDecisionItemIdValue",
				AccessReviewStageId:                 "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId1: "accessReviewInstanceDecisionItemId1Value",
				GovernanceInsightId:                 "governanceInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessReviewInstanceId != v.Expected.AccessReviewInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceId", v.Expected.AccessReviewInstanceId, actual.AccessReviewInstanceId)
		}

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId1 != v.Expected.AccessReviewInstanceDecisionItemId1 {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId1", v.Expected.AccessReviewInstanceDecisionItemId1, actual.AccessReviewInstanceDecisionItemId1)
		}

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId
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
			Input: "/me/pendingAccessReviewInstances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/iNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue",
			Expected: &MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{
				AccessReviewInstanceId:              "accessReviewInstanceIdValue",
				AccessReviewInstanceDecisionItemId:  "accessReviewInstanceDecisionItemIdValue",
				AccessReviewStageId:                 "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId1: "accessReviewInstanceDecisionItemId1Value",
				GovernanceInsightId:                 "governanceInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemId1Value/insights/governanceInsightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/iNsIgHtS/gOvErNaNcEiNsIgHtIdVaLuE",
			Expected: &MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{
				AccessReviewInstanceId:              "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewInstanceDecisionItemId:  "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
				AccessReviewStageId:                 "aCcEsSrEvIeWsTaGeIdVaLuE",
				AccessReviewInstanceDecisionItemId1: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE",
				GovernanceInsightId:                 "gOvErNaNcEiNsIgHtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD1VaLuE/iNsIgHtS/gOvErNaNcEiNsIgHtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AccessReviewInstanceId != v.Expected.AccessReviewInstanceId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceId", v.Expected.AccessReviewInstanceId, actual.AccessReviewInstanceId)
		}

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId1 != v.Expected.AccessReviewInstanceDecisionItemId1 {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId1", v.Expected.AccessReviewInstanceDecisionItemId1, actual.AccessReviewInstanceDecisionItemId1)
		}

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestSegmentsForMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId(t *testing.T) {
	segments := MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
