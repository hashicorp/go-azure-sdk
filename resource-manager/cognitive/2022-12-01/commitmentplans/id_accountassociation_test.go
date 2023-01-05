package commitmentplans

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = AccountAssociationId{}

func TestNewAccountAssociationID(t *testing.T) {
	id := NewAccountAssociationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "commitmentPlanValue", "commitmentPlanAssociationValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.CommitmentPlanName != "commitmentPlanValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CommitmentPlanName'", id.CommitmentPlanName, "commitmentPlanValue")
	}

	if id.CommitmentPlanAssociationName != "commitmentPlanAssociationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CommitmentPlanAssociationName'", id.CommitmentPlanAssociationName, "commitmentPlanAssociationValue")
	}
}

func TestFormatAccountAssociationID(t *testing.T) {
	actual := NewAccountAssociationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "commitmentPlanValue", "commitmentPlanAssociationValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations/commitmentPlanAssociationValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAccountAssociationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AccountAssociationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations/commitmentPlanAssociationValue",
			Expected: &AccountAssociationId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				CommitmentPlanName:            "commitmentPlanValue",
				CommitmentPlanAssociationName: "commitmentPlanAssociationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations/commitmentPlanAssociationValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAccountAssociationID(v.Input)
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

		if actual.CommitmentPlanName != v.Expected.CommitmentPlanName {
			t.Fatalf("Expected %q but got %q for CommitmentPlanName", v.Expected.CommitmentPlanName, actual.CommitmentPlanName)
		}

		if actual.CommitmentPlanAssociationName != v.Expected.CommitmentPlanAssociationName {
			t.Fatalf("Expected %q but got %q for CommitmentPlanAssociationName", v.Expected.CommitmentPlanAssociationName, actual.CommitmentPlanAssociationName)
		}

	}
}

func TestParseAccountAssociationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AccountAssociationId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/cOmMiTmEnTpLaNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/cOmMiTmEnTpLaNs/cOmMiTmEnTpLaNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/cOmMiTmEnTpLaNs/cOmMiTmEnTpLaNvAlUe/aCcOuNtAsSoCiAtIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations/commitmentPlanAssociationValue",
			Expected: &AccountAssociationId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "example-resource-group",
				CommitmentPlanName:            "commitmentPlanValue",
				CommitmentPlanAssociationName: "commitmentPlanAssociationValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanValue/accountAssociations/commitmentPlanAssociationValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/cOmMiTmEnTpLaNs/cOmMiTmEnTpLaNvAlUe/aCcOuNtAsSoCiAtIoNs/cOmMiTmEnTpLaNaSsOcIaTiOnVaLuE",
			Expected: &AccountAssociationId{
				SubscriptionId:                "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:             "eXaMpLe-rEsOuRcE-GrOuP",
				CommitmentPlanName:            "cOmMiTmEnTpLaNvAlUe",
				CommitmentPlanAssociationName: "cOmMiTmEnTpLaNaSsOcIaTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/cOmMiTmEnTpLaNs/cOmMiTmEnTpLaNvAlUe/aCcOuNtAsSoCiAtIoNs/cOmMiTmEnTpLaNaSsOcIaTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAccountAssociationIDInsensitively(v.Input)
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

		if actual.CommitmentPlanName != v.Expected.CommitmentPlanName {
			t.Fatalf("Expected %q but got %q for CommitmentPlanName", v.Expected.CommitmentPlanName, actual.CommitmentPlanName)
		}

		if actual.CommitmentPlanAssociationName != v.Expected.CommitmentPlanAssociationName {
			t.Fatalf("Expected %q but got %q for CommitmentPlanAssociationName", v.Expected.CommitmentPlanAssociationName, actual.CommitmentPlanAssociationName)
		}

	}
}

func TestSegmentsForAccountAssociationId(t *testing.T) {
	segments := AccountAssociationId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AccountAssociationId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
