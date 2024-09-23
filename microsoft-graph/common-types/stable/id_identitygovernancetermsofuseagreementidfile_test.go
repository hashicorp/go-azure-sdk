package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileId{}

func TestNewIdentityGovernanceTermsOfUseAgreementIdFileID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementIdFileID("agreementId", "agreementFileLocalizationId")

	if id.AgreementId != "agreementId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementId'", id.AgreementId, "agreementId")
	}

	if id.AgreementFileLocalizationId != "agreementFileLocalizationId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementFileLocalizationId'", id.AgreementFileLocalizationId, "agreementFileLocalizationId")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementIdFileID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementIdFileID("agreementId", "agreementFileLocalizationId").ID()
	expected := "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdFileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdFileId
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
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileId{
				AgreementId:                 "agreementId",
				AgreementFileLocalizationId: "agreementFileLocalizationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileID(v.Input)
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

		if actual.AgreementFileLocalizationId != v.Expected.AgreementFileLocalizationId {
			t.Fatalf("Expected %q but got %q for AgreementFileLocalizationId", v.Expected.AgreementFileLocalizationId, actual.AgreementFileLocalizationId)
		}

	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdFileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdFileId
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
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileId{
				AgreementId:                 "agreementId",
				AgreementFileLocalizationId: "agreementFileLocalizationId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileId{
				AgreementId:                 "aGrEeMeNtId",
				AgreementFileLocalizationId: "aGrEeMeNtFiLeLoCaLiZaTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileIDInsensitively(v.Input)
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

		if actual.AgreementFileLocalizationId != v.Expected.AgreementFileLocalizationId {
			t.Fatalf("Expected %q but got %q for AgreementFileLocalizationId", v.Expected.AgreementFileLocalizationId, actual.AgreementFileLocalizationId)
		}

	}
}

func TestSegmentsForIdentityGovernanceTermsOfUseAgreementIdFileId(t *testing.T) {
	segments := IdentityGovernanceTermsOfUseAgreementIdFileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceTermsOfUseAgreementIdFileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
