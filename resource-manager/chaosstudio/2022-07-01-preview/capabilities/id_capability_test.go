package capabilities

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CapabilityId{}

func TestNewCapabilityID(t *testing.T) {
	id := NewCapabilityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "parentProviderNamespaceValue", "parentResourceTypeValue", "parentResourceValue", "targetValue", "capabilityValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ParentProviderNamespace != "parentProviderNamespaceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ParentProviderNamespace'", id.ParentProviderNamespace, "parentProviderNamespaceValue")
	}

	if id.ParentResourceType != "parentResourceTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ParentResourceType'", id.ParentResourceType, "parentResourceTypeValue")
	}

	if id.ParentResourceName != "parentResourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ParentResourceName'", id.ParentResourceName, "parentResourceValue")
	}

	if id.TargetName != "targetValue" {
		t.Fatalf("Expected %q but got %q for Segment 'TargetName'", id.TargetName, "targetValue")
	}

	if id.CapabilityName != "capabilityValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CapabilityName'", id.CapabilityName, "capabilityValue")
	}
}

func TestFormatCapabilityID(t *testing.T) {
	actual := NewCapabilityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "parentProviderNamespaceValue", "parentResourceTypeValue", "parentResourceValue", "targetValue", "capabilityValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities/capabilityValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCapabilityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CapabilityId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities/capabilityValue",
			Expected: &CapabilityId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				ParentProviderNamespace: "parentProviderNamespaceValue",
				ParentResourceType:      "parentResourceTypeValue",
				ParentResourceName:      "parentResourceValue",
				TargetName:              "targetValue",
				CapabilityName:          "capabilityValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities/capabilityValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCapabilityID(v.Input)
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

		if actual.ParentProviderNamespace != v.Expected.ParentProviderNamespace {
			t.Fatalf("Expected %q but got %q for ParentProviderNamespace", v.Expected.ParentProviderNamespace, actual.ParentProviderNamespace)
		}

		if actual.ParentResourceType != v.Expected.ParentResourceType {
			t.Fatalf("Expected %q but got %q for ParentResourceType", v.Expected.ParentResourceType, actual.ParentResourceType)
		}

		if actual.ParentResourceName != v.Expected.ParentResourceName {
			t.Fatalf("Expected %q but got %q for ParentResourceName", v.Expected.ParentResourceName, actual.ParentResourceName)
		}

		if actual.TargetName != v.Expected.TargetName {
			t.Fatalf("Expected %q but got %q for TargetName", v.Expected.TargetName, actual.TargetName)
		}

		if actual.CapabilityName != v.Expected.CapabilityName {
			t.Fatalf("Expected %q but got %q for CapabilityName", v.Expected.CapabilityName, actual.CapabilityName)
		}

	}
}

func TestParseCapabilityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CapabilityId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.cHaOs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.cHaOs/tArGeTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.cHaOs/tArGeTs/tArGeTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.cHaOs/tArGeTs/tArGeTvAlUe/cApAbIlItIeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities/capabilityValue",
			Expected: &CapabilityId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "example-resource-group",
				ParentProviderNamespace: "parentProviderNamespaceValue",
				ParentResourceType:      "parentResourceTypeValue",
				ParentResourceName:      "parentResourceValue",
				TargetName:              "targetValue",
				CapabilityName:          "capabilityValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/parentProviderNamespaceValue/parentResourceTypeValue/parentResourceValue/providers/Microsoft.Chaos/targets/targetValue/capabilities/capabilityValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.cHaOs/tArGeTs/tArGeTvAlUe/cApAbIlItIeS/cApAbIlItYvAlUe",
			Expected: &CapabilityId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:       "eXaMpLe-rEsOuRcE-GrOuP",
				ParentProviderNamespace: "pArEnTpRoViDeRnAmEsPaCeVaLuE",
				ParentResourceType:      "pArEnTrEsOuRcEtYpEvAlUe",
				ParentResourceName:      "pArEnTrEsOuRcEvAlUe",
				TargetName:              "tArGeTvAlUe",
				CapabilityName:          "cApAbIlItYvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pArEnTpRoViDeRnAmEsPaCeVaLuE/pArEnTrEsOuRcEtYpEvAlUe/pArEnTrEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.cHaOs/tArGeTs/tArGeTvAlUe/cApAbIlItIeS/cApAbIlItYvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCapabilityIDInsensitively(v.Input)
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

		if actual.ParentProviderNamespace != v.Expected.ParentProviderNamespace {
			t.Fatalf("Expected %q but got %q for ParentProviderNamespace", v.Expected.ParentProviderNamespace, actual.ParentProviderNamespace)
		}

		if actual.ParentResourceType != v.Expected.ParentResourceType {
			t.Fatalf("Expected %q but got %q for ParentResourceType", v.Expected.ParentResourceType, actual.ParentResourceType)
		}

		if actual.ParentResourceName != v.Expected.ParentResourceName {
			t.Fatalf("Expected %q but got %q for ParentResourceName", v.Expected.ParentResourceName, actual.ParentResourceName)
		}

		if actual.TargetName != v.Expected.TargetName {
			t.Fatalf("Expected %q but got %q for TargetName", v.Expected.TargetName, actual.TargetName)
		}

		if actual.CapabilityName != v.Expected.CapabilityName {
			t.Fatalf("Expected %q but got %q for CapabilityName", v.Expected.CapabilityName, actual.CapabilityName)
		}

	}
}

func TestSegmentsForCapabilityId(t *testing.T) {
	segments := CapabilityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CapabilityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
