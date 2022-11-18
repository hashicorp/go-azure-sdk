package lots

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BillingProfileId{}

func TestNewBillingProfileID(t *testing.T) {
	id := NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

	if id.BillingAccountId != "billingAccountIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountId'", id.BillingAccountId, "billingAccountIdValue")
	}

	if id.BillingProfileId != "billingProfileIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingProfileId'", id.BillingProfileId, "billingProfileIdValue")
	}
}

func TestFormatBillingProfileID(t *testing.T) {
	actual := NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles/billingProfileIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingProfileID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileId
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
			Input: "/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles/billingProfileIdValue",
			Expected: &BillingProfileId{
				BillingAccountId: "billingAccountIdValue",
				BillingProfileId: "billingProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles/billingProfileIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountId != v.Expected.BillingAccountId {
			t.Fatalf("Expected %q but got %q for BillingAccountId", v.Expected.BillingAccountId, actual.BillingAccountId)
		}

		if actual.BillingProfileId != v.Expected.BillingProfileId {
			t.Fatalf("Expected %q but got %q for BillingProfileId", v.Expected.BillingProfileId, actual.BillingProfileId)
		}

	}
}

func TestParseBillingProfileIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingProfileId
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
			Input: "/providers/Microsoft.Billing",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/bIlLiNgPrOfIlEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles/billingProfileIdValue",
			Expected: &BillingProfileId{
				BillingAccountId: "billingAccountIdValue",
				BillingProfileId: "billingProfileIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/billingProfiles/billingProfileIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEiDvAlUe",
			Expected: &BillingProfileId{
				BillingAccountId: "bIlLiNgAcCoUnTiDvAlUe",
				BillingProfileId: "bIlLiNgPrOfIlEiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/bIlLiNgPrOfIlEs/bIlLiNgPrOfIlEiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingProfileIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingAccountId != v.Expected.BillingAccountId {
			t.Fatalf("Expected %q but got %q for BillingAccountId", v.Expected.BillingAccountId, actual.BillingAccountId)
		}

		if actual.BillingProfileId != v.Expected.BillingProfileId {
			t.Fatalf("Expected %q but got %q for BillingProfileId", v.Expected.BillingProfileId, actual.BillingProfileId)
		}

	}
}

func TestSegmentsForBillingProfileId(t *testing.T) {
	segments := BillingProfileId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingProfileId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
