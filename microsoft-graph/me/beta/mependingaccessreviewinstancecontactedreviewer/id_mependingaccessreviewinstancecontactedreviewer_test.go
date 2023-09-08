package mependingaccessreviewinstancecontactedreviewer

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MePendingAccessReviewInstanceContactedReviewerId{}

func TestNewMePendingAccessReviewInstanceContactedReviewerID(t *testing.T) {
	id := NewMePendingAccessReviewInstanceContactedReviewerID("accessReviewInstanceIdValue", "accessReviewReviewerIdValue")

	if id.AccessReviewInstanceId != "accessReviewInstanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewInstanceId'", id.AccessReviewInstanceId, "accessReviewInstanceIdValue")
	}

	if id.AccessReviewReviewerId != "accessReviewReviewerIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AccessReviewReviewerId'", id.AccessReviewReviewerId, "accessReviewReviewerIdValue")
	}
}

func TestFormatMePendingAccessReviewInstanceContactedReviewerID(t *testing.T) {
	actual := NewMePendingAccessReviewInstanceContactedReviewerID("accessReviewInstanceIdValue", "accessReviewReviewerIdValue").ID()
	expected := "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers/accessReviewReviewerIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMePendingAccessReviewInstanceContactedReviewerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceContactedReviewerId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers/accessReviewReviewerIdValue",
			Expected: &MePendingAccessReviewInstanceContactedReviewerId{
				AccessReviewInstanceId: "accessReviewInstanceIdValue",
				AccessReviewReviewerId: "accessReviewReviewerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers/accessReviewReviewerIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceContactedReviewerID(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestParseMePendingAccessReviewInstanceContactedReviewerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MePendingAccessReviewInstanceContactedReviewerId
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
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/cOnTaCtEdReViEwErS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers/accessReviewReviewerIdValue",
			Expected: &MePendingAccessReviewInstanceContactedReviewerId{
				AccessReviewInstanceId: "accessReviewInstanceIdValue",
				AccessReviewReviewerId: "accessReviewReviewerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/pendingAccessReviewInstances/accessReviewInstanceIdValue/contactedReviewers/accessReviewReviewerIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiDvAlUe",
			Expected: &MePendingAccessReviewInstanceContactedReviewerId{
				AccessReviewInstanceId: "aCcEsSrEvIeWiNsTaNcEiDvAlUe",
				AccessReviewReviewerId: "aCcEsSrEvIeWrEvIeWeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/pEnDiNgAcCeSsReViEwInStAnCeS/aCcEsSrEvIeWiNsTaNcEiDvAlUe/cOnTaCtEdReViEwErS/aCcEsSrEvIeWrEvIeWeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMePendingAccessReviewInstanceContactedReviewerIDInsensitively(v.Input)
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

		if actual.AccessReviewReviewerId != v.Expected.AccessReviewReviewerId {
			t.Fatalf("Expected %q but got %q for AccessReviewReviewerId", v.Expected.AccessReviewReviewerId, actual.AccessReviewReviewerId)
		}

	}
}

func TestSegmentsForMePendingAccessReviewInstanceContactedReviewerId(t *testing.T) {
	segments := MePendingAccessReviewInstanceContactedReviewerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MePendingAccessReviewInstanceContactedReviewerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
