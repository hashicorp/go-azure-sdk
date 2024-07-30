package termsofuseagreementacceptance

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementAcceptanceId{}

func TestNewIdentityGovernanceTermsOfUseAgreementAcceptanceID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementAcceptanceID("agreementAcceptanceIdValue")

	if id.AgreementAcceptanceId != "agreementAcceptanceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementAcceptanceId'", id.AgreementAcceptanceId, "agreementAcceptanceIdValue")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementAcceptanceID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementAcceptanceID("agreementAcceptanceIdValue").ID()
	expected := "/identityGovernance/termsOfUse/agreementAcceptances/agreementAcceptanceIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementAcceptanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementAcceptanceId
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
			Input: "/identityGovernance/termsOfUse/agreementAcceptances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreementAcceptances/agreementAcceptanceIdValue",
			Expected: &IdentityGovernanceTermsOfUseAgreementAcceptanceId{
				AgreementAcceptanceId: "agreementAcceptanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreementAcceptances/agreementAcceptanceIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementAcceptanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementAcceptanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementAcceptanceId
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
			Input: "/identityGovernance/termsOfUse/agreementAcceptances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtAcCePtAnCeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreementAcceptances/agreementAcceptanceIdValue",
			Expected: &IdentityGovernanceTermsOfUseAgreementAcceptanceId{
				AgreementAcceptanceId: "agreementAcceptanceIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreementAcceptances/agreementAcceptanceIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeIdVaLuE",
			Expected: &IdentityGovernanceTermsOfUseAgreementAcceptanceId{
				AgreementAcceptanceId: "aGrEeMeNtAcCePtAnCeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtAcCePtAnCeS/aGrEeMeNtAcCePtAnCeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementAcceptanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.AgreementAcceptanceId != v.Expected.AgreementAcceptanceId {
			t.Fatalf("Expected %q but got %q for AgreementAcceptanceId", v.Expected.AgreementAcceptanceId, actual.AgreementAcceptanceId)
		}

	}
}

func TestSegmentsForIdentityGovernanceTermsOfUseAgreementAcceptanceId(t *testing.T) {
	segments := IdentityGovernanceTermsOfUseAgreementAcceptanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceTermsOfUseAgreementAcceptanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
