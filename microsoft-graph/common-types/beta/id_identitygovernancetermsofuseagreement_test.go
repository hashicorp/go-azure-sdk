package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementId{}

func TestNewIdentityGovernanceTermsOfUseAgreementID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementID("agreementId")

	if id.AgreementId != "agreementId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementId'", id.AgreementId, "agreementId")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementID("agreementId").ID()
	expected := "/identityGovernance/termsOfUse/agreements/agreementId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId",
			Expected: &IdentityGovernanceTermsOfUseAgreementId{
				AgreementId: "agreementId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AgreementId != v.Expected.AgreementId {
			t.Fatalf("Expected %q but got %q for AgreementId", v.Expected.AgreementId, actual.AgreementId)
		}

	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId",
			Expected: &IdentityGovernanceTermsOfUseAgreementId{
				AgreementId: "agreementId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId",
			Expected: &IdentityGovernanceTermsOfUseAgreementId{
				AgreementId: "aGrEeMeNtId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AgreementId != v.Expected.AgreementId {
			t.Fatalf("Expected %q but got %q for AgreementId", v.Expected.AgreementId, actual.AgreementId)
		}

	}
}

func TestSegmentsForIdentityGovernanceTermsOfUseAgreementId(t *testing.T) {
	segments := IdentityGovernanceTermsOfUseAgreementId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceTermsOfUseAgreementId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
