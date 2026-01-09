package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{}

func TestNewIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID("agreementId", "agreementFileLocalizationId", "agreementFileVersionId")

	if id.AgreementId != "agreementId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementId'", id.AgreementId, "agreementId")
	}

	if id.AgreementFileLocalizationId != "agreementFileLocalizationId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementFileLocalizationId'", id.AgreementFileLocalizationId, "agreementFileLocalizationId")
	}

	if id.AgreementFileVersionId != "agreementFileVersionId" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementFileVersionId'", id.AgreementFileVersionId, "agreementFileVersionId")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID("agreementId", "agreementFileLocalizationId", "agreementFileVersionId").ID()
	expected := "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions/agreementFileVersionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId
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
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions/agreementFileVersionId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{
				AgreementId:                 "agreementId",
				AgreementFileLocalizationId: "agreementFileLocalizationId",
				AgreementFileVersionId:      "agreementFileVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions/agreementFileVersionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionID(v.Input)
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

		if actual.AgreementFileVersionId != v.Expected.AgreementFileVersionId {
			t.Fatalf("Expected %q but got %q for AgreementFileVersionId", v.Expected.AgreementFileVersionId, actual.AgreementFileVersionId)
		}

	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId
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
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnId/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions/agreementFileVersionId",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{
				AgreementId:                 "agreementId",
				AgreementFileLocalizationId: "agreementFileLocalizationId",
				AgreementFileVersionId:      "agreementFileVersionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementId/files/agreementFileLocalizationId/versions/agreementFileVersionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnId/vErSiOnS/aGrEeMeNtFiLeVeRsIoNiD",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{
				AgreementId:                 "aGrEeMeNtId",
				AgreementFileLocalizationId: "aGrEeMeNtFiLeLoCaLiZaTiOnId",
				AgreementFileVersionId:      "aGrEeMeNtFiLeVeRsIoNiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtId/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnId/vErSiOnS/aGrEeMeNtFiLeVeRsIoNiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileIdVersionIDInsensitively(v.Input)
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

		if actual.AgreementFileVersionId != v.Expected.AgreementFileVersionId {
			t.Fatalf("Expected %q but got %q for AgreementFileVersionId", v.Expected.AgreementFileVersionId, actual.AgreementFileVersionId)
		}

	}
}

func TestSegmentsForIdentityGovernanceTermsOfUseAgreementIdFileIdVersionId(t *testing.T) {
	segments := IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
