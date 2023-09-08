package mependingaccessreviewinstancedecisioninstancecontactedreviewer

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}

func TestNewMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(t *testing.T) {
	id := NewMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewReviewerIdValue")

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewInstanceDecisionItemId != "accessReviewInstanceDecisionItemIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceDecisionItemId'", id.AccessReviewInstanceDecisionItemId, "accessReviewInstanceDecisionItemIdValue")
	}

	if id.AccessReviewReviewerId != "accessReviewReviewerIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewReviewerId'", id.AccessReviewReviewerId, "accessReviewReviewerIdValue")
	}
}

func TestFormatMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(t *testing.T) {
	actual := NewMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID("accessReviewInstanceIdValue", "accessReviewInstanceDecisionItemIdValue", "accessReviewReviewerIdValue").ID()
	expected := "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers/accessReviewReviewerIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers/accessReviewReviewerIdValue",
			Expected: &MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{
				AccessReviewInstanceId:             "accessReviewInstanceIdValue",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemIdValue",
				AccessReviewReviewerId:             "accessReviewReviewerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers/accessReviewReviewerIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerID(v.Input)
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

func TestParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/cOnTaCtEdReViEwErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers/accessReviewReviewerIdValue",
			Expected: &MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{
				AccessReviewInstanceId:             "accessReviewInstanceIdValue",
				AccessReviewInstanceDecisionItemId: "accessReviewInstanceDecisionItemIdValue",
				AccessReviewReviewerId:             "accessReviewReviewerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/decisions/accessReviewInstanceDecisionItemIdValue/instance/contactedReviewers/accessReviewReviewerIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiDvAlUe",
			Expected: &MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{
				AccessReviewInstanceId:             "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewInstanceDecisionItemId: "aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe",
				AccessReviewReviewerId:             "aCcEsSrEvIeWrEvIeWeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/dEcIsIoNs/aCcEsSrEvIeWiNsTaNcEdEcIsIoNiTeMiDvAlUe/iNsTaNcE/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceDecisionInstanceContactedReviewerIDInsensitively(v.Input)
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

func TestSegmentsForMePendingAccessReviewInstanceDecisionInstanceContactedReviewerId(t *testing.T) {
	segments := MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePendingAccessReviewInstanceDecisionInstanceContactedReviewerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
