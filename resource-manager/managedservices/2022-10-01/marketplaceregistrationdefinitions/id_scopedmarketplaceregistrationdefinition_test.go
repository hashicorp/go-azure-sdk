package marketplaceregistrationdefinitions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedMarketplaceRegistrationDefinitionId{}

func TestNewScopedMarketplaceRegistrationDefinitionID(t *testing.T) {
	id := NewScopedMarketplaceRegistrationDefinitionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "marketplaceIdentifierValue")

	if id.Scope != "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'Scope'", id.Scope, "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")
	}

	if id.MarketplaceIdentifier != "marketplaceIdentifierValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MarketplaceIdentifier'", id.MarketplaceIdentifier, "marketplaceIdentifierValue")
	}
}

func TestFormatScopedMarketplaceRegistrationDefinitionID(t *testing.T) {
	actual := NewScopedMarketplaceRegistrationDefinitionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "marketplaceIdentifierValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifierValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScopedMarketplaceRegistrationDefinitionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedMarketplaceRegistrationDefinitionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifierValue",
			Expected: &ScopedMarketplaceRegistrationDefinitionId{
				Scope:                 "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				MarketplaceIdentifier: "marketplaceIdentifierValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifierValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedMarketplaceRegistrationDefinitionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.MarketplaceIdentifier != v.Expected.MarketplaceIdentifier {
			t.Fatalf("Expected %q but got %q for MarketplaceIdentifier", v.Expected.MarketplaceIdentifier, actual.MarketplaceIdentifier)
		}

	}
}

func TestParseScopedMarketplaceRegistrationDefinitionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScopedMarketplaceRegistrationDefinitionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/mArKeTpLaCeReGiStRaTiOnDeFiNiTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifierValue",
			Expected: &ScopedMarketplaceRegistrationDefinitionId{
				Scope:                 "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group",
				MarketplaceIdentifier: "marketplaceIdentifierValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplaceIdentifierValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/mArKeTpLaCeReGiStRaTiOnDeFiNiTiOnS/mArKeTpLaCeIdEnTiFiErVaLuE",
			Expected: &ScopedMarketplaceRegistrationDefinitionId{
				Scope:                 "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp",
				MarketplaceIdentifier: "mArKeTpLaCeIdEnTiFiErVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/sOmE-ReSoUrCe-gRoUp/pRoViDeRs/mIcRoSoFt.mAnAgEdSeRvIcEs/mArKeTpLaCeReGiStRaTiOnDeFiNiTiOnS/mArKeTpLaCeIdEnTiFiErVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScopedMarketplaceRegistrationDefinitionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.Scope != v.Expected.Scope {
			t.Fatalf("Expected %q but got %q for Scope", v.Expected.Scope, actual.Scope)
		}

		if actual.MarketplaceIdentifier != v.Expected.MarketplaceIdentifier {
			t.Fatalf("Expected %q but got %q for MarketplaceIdentifier", v.Expected.MarketplaceIdentifier, actual.MarketplaceIdentifier)
		}

	}
}

func TestSegmentsForScopedMarketplaceRegistrationDefinitionId(t *testing.T) {
	segments := ScopedMarketplaceRegistrationDefinitionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScopedMarketplaceRegistrationDefinitionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
