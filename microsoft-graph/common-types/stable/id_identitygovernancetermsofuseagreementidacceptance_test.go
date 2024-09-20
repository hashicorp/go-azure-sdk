package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{}

func TestNewIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementIdAcceptanceID("agreementId", "agreementAcceptanceId")

	if id.AgreementId != "agreementId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementId'", id.AgreementId, "agreementId")
	}

	if id.AgreementAcceptanceId != "agreementAcceptanceId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementAcceptanceId'", id.AgreementAcceptanceId, "agreementAcceptanceId")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementIdAcceptanceID("agreementId", "agreementAcceptanceId").ID()
	expected := "/identityGovernance/termsOfUse/agreements/agreementId/acceptances/agreementAcceptanceId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdAcceptanceId
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
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/acceptances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/acceptances/agreementAcceptanceId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{
				AgreementId:           "agreementId",
				AgreementAcceptanceId: "agreementAcceptanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/acceptances/agreementAcceptanceId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceID(v.Input)
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

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdAcceptanceId
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
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/acceptances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/aCcEpTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/acceptances/agreementAcceptanceId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{
				AgreementId:           "agreementId",
				AgreementAcceptanceId: "agreementAcceptanceId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/acceptances/agreementAcceptanceId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/aCcEpTaNcEs/aGrEeMeNtAcCePtAnCeId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{
				AgreementId:           "aGrEeMeNtId",
				AgreementAcceptanceId: "aGrEeMeNtAcCePtAnCeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/aCcEpTaNcEs/aGrEeMeNtAcCePtAnCeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdAcceptanceIDInsensitively(v.Input)
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

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestSegmentsForIdentityGovernanceTermsOfUseAgreementIdAcceptanceId(t *testing.T) {
	segments := IdentityGovernanceTermsOfUseAgreementIdAcceptanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceTermsOfUseAgreementIdAcceptanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
