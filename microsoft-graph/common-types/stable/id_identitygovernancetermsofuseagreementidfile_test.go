package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceTermsOfUseAgreementIdFileId{}

func TestNewIdentityGovernanceTermsOfUseAgreementIdFileID(t *testing.T) {
	id := NewIdentityGovernanceTermsOfUseAgreementIdFileID("agreementIdValue", "agreementFileLocalizationIdValue")

	if id.AgreementId != "agreementIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementId'", id.AgreementId, "agreementIdValue")
	}

	if id.AgreementFileLocalizationId != "agreementFileLocalizationIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AgreementFileLocalizationId'", id.AgreementFileLocalizationId, "agreementFileLocalizationIdValue")
	}
}

func TestFormatIdentityGovernanceTermsOfUseAgreementIdFileID(t *testing.T) {
	actual := NewIdentityGovernanceTermsOfUseAgreementIdFileID("agreementIdValue", "agreementFileLocalizationIdValue").ID()
	expected := "/identityGovernance/termsOfUse/agreements/agreementIdValue/files/agreementFileLocalizationIdValue"
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
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/files",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/files/agreementFileLocalizationIdValue",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileId{
				AgreementId:                 "agreementIdValue",
				AgreementFileLocalizationId: "agreementFileLocalizationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/files/agreementFileLocalizationIdValue/extra",
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
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/files",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/files/agreementFileLocalizationIdValue",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileId{
				AgreementId:                 "agreementIdValue",
				AgreementFileLocalizationId: "agreementFileLocalizationIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/identityGovernance/termsOfUse/agreements/agreementIdValue/files/agreementFileLocalizationIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE",
			Expected: &IdentityGovernanceTermsOfUseAgreementIdFileId{
				AgreementId:                 "aGrEeMeNtIdVaLuE",
				AgreementFileLocalizationId: "aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/iDeNtItYgOvErNaNcE/tErMsOfUsE/aGrEeMeNtS/aGrEeMeNtIdVaLuE/fIlEs/aGrEeMeNtFiLeLoCaLiZaTiOnIdVaLuE/extra",
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
