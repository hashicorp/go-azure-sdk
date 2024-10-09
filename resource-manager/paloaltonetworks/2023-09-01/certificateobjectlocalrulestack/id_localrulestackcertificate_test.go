package certificateobjectlocalrulestack

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LocalRulestackCertificateId{}

func TestNewLocalRulestackCertificateID(t *testing.T) {
	id := NewLocalRulestackCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRulestackName", "certificateName")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.LocalRulestackName != "localRulestackName" {
		t.Fatalf("Expected %q but got %q for Segment 'LocalRulestackName'", id.LocalRulestackName, "localRulestackName")
	}

	if id.CertificateName != "certificateName" {
		t.Fatalf("Expected %q but got %q for Segment 'CertificateName'", id.CertificateName, "certificateName")
	}
}

func TestFormatLocalRulestackCertificateID(t *testing.T) {
	actual := NewLocalRulestackCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "localRulestackName", "certificateName").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates/certificateName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseLocalRulestackCertificateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocalRulestackCertificateId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates/certificateName",
			Expected: &LocalRulestackCertificateId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				LocalRulestackName: "localRulestackName",
				CertificateName:    "certificateName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates/certificateName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocalRulestackCertificateID(v.Input)
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

		if actual.LocalRulestackName != v.Expected.LocalRulestackName {
			t.Fatalf("Expected %q but got %q for LocalRulestackName", v.Expected.LocalRulestackName, actual.LocalRulestackName)
		}

		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}

	}
}

func TestParseLocalRulestackCertificateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *LocalRulestackCertificateId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE/cErTiFiCaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates/certificateName",
			Expected: &LocalRulestackCertificateId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				LocalRulestackName: "localRulestackName",
				CertificateName:    "certificateName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/localRulestackName/certificates/certificateName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE/cErTiFiCaTeS/cErTiFiCaTeNaMe",
			Expected: &LocalRulestackCertificateId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "eXaMpLe-rEsOuRcE-GrOuP",
				LocalRulestackName: "lOcAlRuLeStAcKnAmE",
				CertificateName:    "cErTiFiCaTeNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pAlOaLtOnEtWoRkS.ClOuDnGfW/lOcAlRuLeStAcKs/lOcAlRuLeStAcKnAmE/cErTiFiCaTeS/cErTiFiCaTeNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseLocalRulestackCertificateIDInsensitively(v.Input)
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

		if actual.LocalRulestackName != v.Expected.LocalRulestackName {
			t.Fatalf("Expected %q but got %q for LocalRulestackName", v.Expected.LocalRulestackName, actual.LocalRulestackName)
		}

		if actual.CertificateName != v.Expected.CertificateName {
			t.Fatalf("Expected %q but got %q for CertificateName", v.Expected.CertificateName, actual.CertificateName)
		}

	}
}

func TestSegmentsForLocalRulestackCertificateId(t *testing.T) {
	segments := LocalRulestackCertificateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("LocalRulestackCertificateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
