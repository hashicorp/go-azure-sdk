package benefitutilizationsummaries

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SavingsPlanId{}

func TestNewSavingsPlanID(t *testing.T) {
	id := NewSavingsPlanID("savingsPlanOrderIdValue", "savingsPlanIdValue")

	if id.SavingsPlanOrderId != "savingsPlanOrderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SavingsPlanOrderId'", id.SavingsPlanOrderId, "savingsPlanOrderIdValue")
	}

	if id.SavingsPlanId != "savingsPlanIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SavingsPlanId'", id.SavingsPlanId, "savingsPlanIdValue")
	}
}

func TestFormatSavingsPlanID(t *testing.T) {
	actual := NewSavingsPlanID("savingsPlanOrderIdValue", "savingsPlanIdValue").ID()
	expected := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseSavingsPlanID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SavingsPlanId
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
			Input: "/providers/Microsoft.BillingBenefits",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue",
			Expected: &SavingsPlanId{
				SavingsPlanOrderId: "savingsPlanOrderIdValue",
				SavingsPlanId:      "savingsPlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSavingsPlanID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SavingsPlanOrderId != v.Expected.SavingsPlanOrderId {
			t.Fatalf("Expected %q but got %q for SavingsPlanOrderId", v.Expected.SavingsPlanOrderId, actual.SavingsPlanOrderId)
		}

		if actual.SavingsPlanId != v.Expected.SavingsPlanId {
			t.Fatalf("Expected %q but got %q for SavingsPlanId", v.Expected.SavingsPlanId, actual.SavingsPlanId)
		}

	}
}

func TestParseSavingsPlanIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SavingsPlanId
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
			Input: "/providers/Microsoft.BillingBenefits",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNgBeNeFiTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNgBeNeFiTs/sAvInGsPlAnOrDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNgBeNeFiTs/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNgBeNeFiTs/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/sAvInGsPlAnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue",
			Expected: &SavingsPlanId{
				SavingsPlanOrderId: "savingsPlanOrderIdValue",
				SavingsPlanId:      "savingsPlanIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.BillingBenefits/savingsPlanOrders/savingsPlanOrderIdValue/savingsPlans/savingsPlanIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNgBeNeFiTs/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/sAvInGsPlAnS/sAvInGsPlAnIdVaLuE",
			Expected: &SavingsPlanId{
				SavingsPlanOrderId: "sAvInGsPlAnOrDeRiDvAlUe",
				SavingsPlanId:      "sAvInGsPlAnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNgBeNeFiTs/sAvInGsPlAnOrDeRs/sAvInGsPlAnOrDeRiDvAlUe/sAvInGsPlAnS/sAvInGsPlAnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseSavingsPlanIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SavingsPlanOrderId != v.Expected.SavingsPlanOrderId {
			t.Fatalf("Expected %q but got %q for SavingsPlanOrderId", v.Expected.SavingsPlanOrderId, actual.SavingsPlanOrderId)
		}

		if actual.SavingsPlanId != v.Expected.SavingsPlanId {
			t.Fatalf("Expected %q but got %q for SavingsPlanId", v.Expected.SavingsPlanId, actual.SavingsPlanId)
		}

	}
}

func TestSegmentsForSavingsPlanId(t *testing.T) {
	segments := SavingsPlanId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("SavingsPlanId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
