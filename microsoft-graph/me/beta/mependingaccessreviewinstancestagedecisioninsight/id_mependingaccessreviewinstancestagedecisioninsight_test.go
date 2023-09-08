package mependingaccessreviewinstancestagedecisioninsight

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceStageDecisionInsightId{}

func TestNewMePendingAccessReviewInstanceStageDecisionInsightID(t *testing.T) {
	id := NewMePendingAccessReviewInstanceStageDecisionInsightID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "governanceInsightIdValue")

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewStageId != "accessReviewStageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewStageId'", id.AccessReviewStageId, "accessReviewStageIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemIdValue")
	}

	if id.GovernanceInsightId != "governanceInsightIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GovernanceInsightId'", id.GovernanceInsightId, "governanceInsightIdValue")
	}
}

func TestFormatMePendingAccessReviewInstanceStageDecisionInsightID(t *testing.T) {
	actual := NewMePendingAccessReviewInstanceStageDecisionInsightID("accessReviewInstanceIdValue", "accessReviewStageIdValue", "accessReviewInstanceDecisionItemIdValue", "governanceInsightIdValue").ID()
	expected := "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights/governanceInsightIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePendingAccessReviewInstanceStageDecisionInsightID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceStageDecisionInsightId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights/governanceInsightIdValue",
			Expected: &MePendingAccessReviewInstanceStageDecisionInsightId{
				AccessReviewInstanceId:             "accessReviewInstanceIdValue",
				AccessReviewStageId:                "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemIdValue",
				GovernanceInsightId:                "governanceInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights/governanceInsightIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceStageDecisionInsightID(v.Input)
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

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestParseMePendingAccessReviewInstanceStageDecisionInsightIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceStageDecisionInsightId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsIgHtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights/governanceInsightIdValue",
			Expected: &MePendingAccessReviewInstanceStageDecisionInsightId{
				AccessReviewInstanceId:             "accessReviewInstanceIdValue",
				AccessReviewStageId:                "accessReviewStageIdValue",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemIdValue",
				GovernanceInsightId:                "governanceInsightIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/stages/accessReviewStageIdValue/decisions/accessReviewInstanceDecisionItemIdValue/insights/governanceInsightIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsIgHtS/gOvErNaNcEiNsIgHtIdVaLuE",
			Expected: &MePendingAccessReviewInstanceStageDecisionInsightId{
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewStageId:                "aCcEsSrEvIeWsTaGeIdVaLuE",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
				GovernanceInsightId:                "gOvErNaNcEiNsIgHtIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/sTaGeS/aCcEsSrEvIeWsTaGeIdVaLuE/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsIgHtS/gOvErNaNcEiNsIgHtIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceStageDecisionInsightIDInsensitively(v.Input)
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

		if actual.AccessReviewStageId != v.Expected.AccessReviewStageId {
			t.Fatalf("Expected %q but got %q for AccessReviewStageId", v.Expected.AccessReviewStageId, actual.AccessReviewStageId)
		}

		if actual.AccessReviewInstanceDecisionItemId != v.Expected.AccessReviewInstanceDecisionItemId {
			t.Fatalf("Expected %q but got %q for AccessReviewInstanceDecisionItemId", v.Expected.AccessReviewInstanceDecisionItemId, actual.AccessReviewInstanceDecisionItemId)
		}

		if actual.GovernanceInsightId != v.Expected.GovernanceInsightId {
			t.Fatalf("Expected %q but got %q for GovernanceInsightId", v.Expected.GovernanceInsightId, actual.GovernanceInsightId)
		}

	}
}

func TestSegmentsForMePendingAccessReviewInstanceStageDecisionInsightId(t *testing.T) {
	segments := MePendingAccessReviewInstanceStageDecisionInsightId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePendingAccessReviewInstanceStageDecisionInsightId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
