package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}

func TestNewMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(t *testing.T) {
	id := NewMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewReviewerId")

	if id.AccessReviewInstanceId != "accessReviewInstanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceId")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemId")
	}

	if id.AccessReviewReviewerId != "accessReviewReviewerId" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewReviewerId'", id.AccessReviewReviewerId, "accessReviewReviewerId")
	}
}

func TestFormatMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(t *testing.T) {
	actual := NewMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID("accessReviewInstanceId", "accessReviewInstanceDecisionItemId", "accessReviewReviewerId").ID()
	expected := "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId",
			Expected: &MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				AccessReviewReviewerId:             "accessReviewReviewerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/cOnTaCtEdReViEwErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId",
			Expected: &MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
				AccessReviewInstanceId:             "accessReviewInstanceId",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemId",
				AccessReviewReviewerId:             "accessReviewReviewerId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceId/decisions/accessReviewInstanceDecisionItemId/instance/contactedReviewers/accessReviewReviewerId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiD",
			Expected: &MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiD",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD",
				AccessReviewReviewerId:             "aCcEsSrEvIeWrEvIeWeRiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiD/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiD/iNsTaNcE/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestSegmentsForMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId(t *testing.T) {
	segments := MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
