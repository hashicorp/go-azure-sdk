package applyupdates

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2ApplyUpdateId{}

func TestNewProviders2ApplyUpdateID(t *testing.T) {
	id := NewProviders2ApplyUpdateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceParentTypeValue", "resourceParentValue", "resourceTypeValue", "resourceValue", "applyUpdateValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.ProviderName != "providerValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ProviderName'", id.ProviderName, "providerValue")
	}

	if id.ResourceParentType != "resourceParentTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceParentType'", id.ResourceParentType, "resourceParentTypeValue")
	}

	if id.ResourceParentName != "resourceParentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceParentName'", id.ResourceParentName, "resourceParentValue")
	}

	if id.ResourceType != "resourceTypeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceType'", id.ResourceType, "resourceTypeValue")
	}

	if id.ResourceName != "resourceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceName'", id.ResourceName, "resourceValue")
	}

	if id.ApplyUpdateName != "applyUpdateValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ApplyUpdateName'", id.ApplyUpdateName, "applyUpdateValue")
	}
}

func TestFormatProviders2ApplyUpdateID(t *testing.T) {
	actual := NewProviders2ApplyUpdateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "resourceParentTypeValue", "resourceParentValue", "resourceTypeValue", "resourceValue", "applyUpdateValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates/applyUpdateValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviders2ApplyUpdateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2ApplyUpdateId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates/applyUpdateValue",
			Expected: &Providers2ApplyUpdateId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				ProviderName:       "providerValue",
				ResourceParentType: "resourceParentTypeValue",
				ResourceParentName: "resourceParentValue",
				ResourceType:       "resourceTypeValue",
				ResourceName:       "resourceValue",
				ApplyUpdateName:    "applyUpdateValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates/applyUpdateValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2ApplyUpdateID(v.Input)
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

		if actual.ProviderName != v.Expected.ProviderName {
			t.Fatalf("Expected %q but got %q for ProviderName", v.Expected.ProviderName, actual.ProviderName)
		}

		if actual.ResourceParentType != v.Expected.ResourceParentType {
			t.Fatalf("Expected %q but got %q for ResourceParentType", v.Expected.ResourceParentType, actual.ResourceParentType)
		}

		if actual.ResourceParentName != v.Expected.ResourceParentName {
			t.Fatalf("Expected %q but got %q for ResourceParentName", v.Expected.ResourceParentName, actual.ResourceParentName)
		}

		if actual.ResourceType != v.Expected.ResourceType {
			t.Fatalf("Expected %q but got %q for ResourceType", v.Expected.ResourceType, actual.ResourceType)
		}

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.ApplyUpdateName != v.Expected.ApplyUpdateName {
			t.Fatalf("Expected %q but got %q for ApplyUpdateName", v.Expected.ApplyUpdateName, actual.ApplyUpdateName)
		}

	}
}

func TestParseProviders2ApplyUpdateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *Providers2ApplyUpdateId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/aPpLyUpDaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates/applyUpdateValue",
			Expected: &Providers2ApplyUpdateId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "example-resource-group",
				ProviderName:       "providerValue",
				ResourceParentType: "resourceParentTypeValue",
				ResourceParentName: "resourceParentValue",
				ResourceType:       "resourceTypeValue",
				ResourceName:       "resourceValue",
				ApplyUpdateName:    "applyUpdateValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/providerValue/resourceParentTypeValue/resourceParentValue/resourceTypeValue/resourceValue/providers/Microsoft.Maintenance/applyUpdates/applyUpdateValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/aPpLyUpDaTeS/aPpLyUpDaTeVaLuE",
			Expected: &Providers2ApplyUpdateId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:  "eXaMpLe-rEsOuRcE-GrOuP",
				ProviderName:       "pRoViDeRvAlUe",
				ResourceParentType: "rEsOuRcEpArEnTtYpEvAlUe",
				ResourceParentName: "rEsOuRcEpArEnTvAlUe",
				ResourceType:       "rEsOuRcEtYpEvAlUe",
				ResourceName:       "rEsOuRcEvAlUe",
				ApplyUpdateName:    "aPpLyUpDaTeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/pRoViDeRvAlUe/rEsOuRcEpArEnTtYpEvAlUe/rEsOuRcEpArEnTvAlUe/rEsOuRcEtYpEvAlUe/rEsOuRcEvAlUe/pRoViDeRs/mIcRoSoFt.mAiNtEnAnCe/aPpLyUpDaTeS/aPpLyUpDaTeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviders2ApplyUpdateIDInsensitively(v.Input)
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

		if actual.ProviderName != v.Expected.ProviderName {
			t.Fatalf("Expected %q but got %q for ProviderName", v.Expected.ProviderName, actual.ProviderName)
		}

		if actual.ResourceParentType != v.Expected.ResourceParentType {
			t.Fatalf("Expected %q but got %q for ResourceParentType", v.Expected.ResourceParentType, actual.ResourceParentType)
		}

		if actual.ResourceParentName != v.Expected.ResourceParentName {
			t.Fatalf("Expected %q but got %q for ResourceParentName", v.Expected.ResourceParentName, actual.ResourceParentName)
		}

		if actual.ResourceType != v.Expected.ResourceType {
			t.Fatalf("Expected %q but got %q for ResourceType", v.Expected.ResourceType, actual.ResourceType)
		}

		if actual.ResourceName != v.Expected.ResourceName {
			t.Fatalf("Expected %q but got %q for ResourceName", v.Expected.ResourceName, actual.ResourceName)
		}

		if actual.ApplyUpdateName != v.Expected.ApplyUpdateName {
			t.Fatalf("Expected %q but got %q for ApplyUpdateName", v.Expected.ApplyUpdateName, actual.ApplyUpdateName)
		}

	}
}

func TestSegmentsForProviders2ApplyUpdateId(t *testing.T) {
	segments := Providers2ApplyUpdateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("Providers2ApplyUpdateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
