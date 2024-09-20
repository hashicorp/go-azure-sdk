package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdAgreementAcceptanceId{}

func TestNewUserIdAgreementAcceptanceID(t *testing.T) {
	id := NewUserIdAgreementAcceptanceID("userId", "agreementAcceptanceId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.AgreementAcceptanceId != "agreementAcceptanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementAcceptanceId'", id.AgreementAcceptanceId, "agreementAcceptanceId")
	}
}

func TestFormatUserIdAgreementAcceptanceID(t *testing.T) {
	actual := NewUserIdAgreementAcceptanceID("userId", "agreementAcceptanceId").ID()
	expected := "/users/userId/agreementAcceptances/agreementAcceptanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdAgreementAcceptanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAgreementAcceptanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/agreementAcceptances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/agreementAcceptances/agreementAcceptanceId",
			Expected: &UserIdAgreementAcceptanceId{
				UserId:                "userId",
				AgreementAcceptanceId: "agreementAcceptanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/agreementAcceptances/agreementAcceptanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAgreementAcceptanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestParseUserIdAgreementAcceptanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdAgreementAcceptanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/agreementAcceptances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aGrEeMeNtAcCePtAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/agreementAcceptances/agreementAcceptanceId",
			Expected: &UserIdAgreementAcceptanceId{
				UserId:                "userId",
				AgreementAcceptanceId: "agreementAcceptanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/agreementAcceptances/agreementAcceptanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeId",
			Expected: &UserIdAgreementAcceptanceId{
				UserId:                "uSeRiD",
				AgreementAcceptanceId: "aGrEeMeNtAcCePtAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdAgreementAcceptanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestSegmentsForUserIdAgreementAcceptanceId(t *testing.T) {
	segments := UserIdAgreementAcceptanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdAgreementAcceptanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
