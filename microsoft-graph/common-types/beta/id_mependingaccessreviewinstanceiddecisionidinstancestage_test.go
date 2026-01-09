package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{}

func TestNewMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(t *testing.T) {
	id := NewMePendingAccessReviewInstanceIdDecisionIdInstanceStageID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewStageId")

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemId")
	}

	if id.AccessReviewStageId != "accessReviewStageId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewStageId'", id.AccessReviewStageId, "accessReviewStageId")
	}
}

func TestFormatMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(t *testing.T) {
	actual := NewMePendingAccessReviewInstanceIdDecisionIdInstanceStageID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewStageId").ID()
	expected := "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages/accessReviewStageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceIdDecisionIdInstanceStageId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages/accessReviewStageId",
			Expected: &MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				AccessReviewStageId:                "accessReviewStageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages/accessReviewStageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(v.Input)
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

	}
}

func TestParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceIdDecisionIdInstanceStageId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/sTaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages/accessReviewStageId",
			Expected: &MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				AccessReviewStageId:                "accessReviewStageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/stages/accessReviewStageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeId",
			Expected: &MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiD",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
				AccessReviewStageId:                "aCcEsSrEvIeWsTaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/sTaGeS/aCcEsSrEvIeWsTaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMePendingAccessReviewInstanceIdDecisionIdInstanceStageId(t *testing.T) {
	segments := MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePendingAccessReviewInstanceIdDecisionIdInstanceStageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
