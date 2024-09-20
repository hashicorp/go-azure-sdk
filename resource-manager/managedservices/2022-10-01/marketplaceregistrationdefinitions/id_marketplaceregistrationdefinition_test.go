package marketplaceregistrationdefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MarketplaceRegistrationDefinitionId{}

func TestNewMarketplaceRegistrationDefinitionID(t *testing.T) {
	id := NewMarketplaceRegistrationDefinitionID("marketplaceIdentifier")

	if id.MarketplaceIdentifier != "marketplaceIdentifier" {
		t.Fatalf("Expected %q but got %q for Segment 'MarketplaceIdentifier'", id.MarketplaceIdentifier, "marketplaceIdentifier")
	}
}

func TestFormatMarketplaceRegistrationDefinitionID(t *testing.T) {
	actual := NewMarketplaceRegistrationDefinitionID("marketplaceIdentifier").ID()
	expected := "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifier"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMarketplaceRegistrationDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MarketplaceRegistrationDefinitionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.ManagedServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifier",
			Expected: &MarketplaceRegistrationDefinitionId{
				MarketplaceIdentifier: "marketplaceIdentifier",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifier/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMarketplaceRegistrationDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MarketplaceIdentifier != v.Expected.MarketplaceIdentifier {
			t.Fatalf("Expected %q but got %q for MarketplaceIdentifier", v.Expected.MarketplaceIdentifier, actual.MarketplaceIdentifier)
		}

	}
}

func TestParseMarketplaceRegistrationDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MarketplaceRegistrationDefinitionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.ManagedServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/mArKeTpLaCeReGiStRaTiOnDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifier",
			Expected: &MarketplaceRegistrationDefinitionId{
				MarketplaceIdentifier: "marketplaceIdentifier",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifier/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/mArKeTpLaCeReGiStRaTiOnDeFiNiTiOnS/mArKeTpLaCeIdEnTiFiEr",
			Expected: &MarketplaceRegistrationDefinitionId{
				MarketplaceIdentifier: "mArKeTpLaCeIdEnTiFiEr",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/mArKeTpLaCeReGiStRaTiOnDeFiNiTiOnS/mArKeTpLaCeIdEnTiFiEr/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMarketplaceRegistrationDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MarketplaceIdentifier != v.Expected.MarketplaceIdentifier {
			t.Fatalf("Expected %q but got %q for MarketplaceIdentifier", v.Expected.MarketplaceIdentifier, actual.MarketplaceIdentifier)
		}

	}
}

func TestSegmentsForMarketplaceRegistrationDefinitionId(t *testing.T) {
	segments := MarketplaceRegistrationDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MarketplaceRegistrationDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
