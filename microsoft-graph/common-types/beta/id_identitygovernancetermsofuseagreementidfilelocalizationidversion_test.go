package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{}

func TestNewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID("agreementIdValue", "agreementFileLocalizationIdValue", "agreementFileVersionIdValue")

	if id.AgreementId != "agreementIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementId'", id.AgreementId, "agreementIdValue")
	}

	if id.AgreementFileLocalizationId != "agreementFileLocalizationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementFileLocalizationId'", id.AgreementFileLocalizationId, "agreementFileLocalizationIdValue")
	}

	if id.AgreementFileVersionId != "agreementFileVersionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementFileVersionId'", id.AgreementFileVersionId, "agreementFileVersionIdValue")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID("agreementIdValue", "agreementFileLocalizationIdValue", "agreementFileVersionIdValue").ID()
	expected := "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions/agreementFileVersionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId
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
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions/agreementFileVersionIdValue",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{
				AgreementId:                 "agreementIdValue",
				AgreementFileLocalizationId: "agreementFileLocalizationIdValue",
				AgreementFileVersionId:      "agreementFileVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions/agreementFileVersionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionID(v.Input)
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

func TestParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId
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
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlE/lOcAlIzAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlE/lOcAlIzAtIoNs/aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlE/lOcAlIzAtIoNs/aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions/agreementFileVersionIdValue",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{
				AgreementId:                 "agreementIdValue",
				AgreementFileLocalizationId: "agreementFileLocalizationIdValue",
				AgreementFileVersionId:      "agreementFileVersionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/file/localizations/agreementFileLocalizationIdValue/versions/agreementFileVersionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlE/lOcAlIzAtIoNs/aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE/vErSiOnS/aGrEeMeNtFiLeVeRsIoNiDvAlUe",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{
				AgreementId:                 "aGrEeMeNtIdVaLuE",
				AgreementFileLocalizationId: "aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE",
				AgreementFileVersionId:      "aGrEeMeNtFiLeVeRsIoNiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlE/lOcAlIzAtIoNs/aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE/vErSiOnS/aGrEeMeNtFiLeVeRsIoNiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionIDInsensitively(v.Input)
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

func TestSegmentsForIdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId(t *testing.T) {
	segments := IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("IdentityGovernanceTermsOfUseAgreementIdFileLocalizationIdVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
