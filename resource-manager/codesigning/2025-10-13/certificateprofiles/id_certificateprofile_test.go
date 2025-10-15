package certificateprofiles

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CertificateProfileId{}

func TestNewCertificateProfileID(t *testing.T) {
	id := NewCertificateProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "codeSigningAccountName", "certificateProfileName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.CodeSigningAccountName != "codeSigningAccountName" {
		t.Fatalf("Expected %q but got %q for Segment 'CodeSigningAccountName'", id.CodeSigningAccountName, "codeSigningAccountName")
	}

	if id.CertificateProfileName != "certificateProfileName" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateProfileName'", id.CertificateProfileName, "certificateProfileName")
	}
}

func TestFormatCertificateProfileID(t *testing.T) {
	actual := NewCertificateProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "codeSigningAccountName", "certificateProfileName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles/certificateProfileName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCertificateProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CertificateProfileId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles/certificateProfileName",
			Expected: &CertificateProfileId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				CodeSigningAccountName: "codeSigningAccountName",
				CertificateProfileName: "certificateProfileName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles/certificateProfileName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCertificateProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.CodeSigningAccountName != v.Expected.CodeSigningAccountName {
			t.Fatalf("Expected %q but got %q for CodeSigningAccountName", v.Expected.CodeSigningAccountName, actual.CodeSigningAccountName)
		}

		if actual.CertificateProfileName != v.Expected.CertificateProfileName {
			t.Fatalf("Expected %q but got %q for CertificateProfileName", v.Expected.CertificateProfileName, actual.CertificateProfileName)
		}

	}
}

func TestParseCertificateProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CertificateProfileId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOdEsIgNiNg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOdEsIgNiNg/cOdEsIgNiNgAcCoUnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOdEsIgNiNg/cOdEsIgNiNgAcCoUnTs/cOdEsIgNiNgAcCoUnTnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOdEsIgNiNg/cOdEsIgNiNgAcCoUnTs/cOdEsIgNiNgAcCoUnTnAmE/cErTiFiCaTePrOfIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles/certificateProfileName",
			Expected: &CertificateProfileId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				CodeSigningAccountName: "codeSigningAccountName",
				CertificateProfileName: "certificateProfileName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CodeSigning/codeSigningAccounts/codeSigningAccountName/certificateProfiles/certificateProfileName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOdEsIgNiNg/cOdEsIgNiNgAcCoUnTs/cOdEsIgNiNgAcCoUnTnAmE/cErTiFiCaTePrOfIlEs/cErTiFiCaTePrOfIlEnAmE",
			Expected: &CertificateProfileId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "eXaMpLe-rEsOuRcE-GrOuP",
				CodeSigningAccountName: "cOdEsIgNiNgAcCoUnTnAmE",
				CertificateProfileName: "cErTiFiCaTePrOfIlEnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOdEsIgNiNg/cOdEsIgNiNgAcCoUnTs/cOdEsIgNiNgAcCoUnTnAmE/cErTiFiCaTePrOfIlEs/cErTiFiCaTePrOfIlEnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCertificateProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.CodeSigningAccountName != v.Expected.CodeSigningAccountName {
			t.Fatalf("Expected %q but got %q for CodeSigningAccountName", v.Expected.CodeSigningAccountName, actual.CodeSigningAccountName)
		}

		if actual.CertificateProfileName != v.Expected.CertificateProfileName {
			t.Fatalf("Expected %q but got %q for CertificateProfileName", v.Expected.CertificateProfileName, actual.CertificateProfileName)
		}

	}
}

func TestSegmentsForCertificateProfileId(t *testing.T) {
	segments := CertificateProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CertificateProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
