package lots

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CustomerId{}

func TestNewCustomerID(t *testing.T) {
	id := NewCustomerID("billingAccountIdValue", "customerIdValue")

	if id.BillingAccountId != "billingAccountIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingAccountId'", id.BillingAccountId, "billingAccountIdValue")
	}

	if id.CustomerId != "customerIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CustomerId'", id.CustomerId, "customerIdValue")
	}
}

func TestFormatCustomerID(t *testing.T) {
	actual := NewCustomerID("billingAccountIdValue", "customerIdValue").ID()
	expected := "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers/customerIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseCustomerID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CustomerId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers/customerIdValue",
			Expected: &CustomerId{
				BillingAccountId: "billingAccountIdValue",
				CustomerId:       "customerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers/customerIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCustomerID(v.Input)
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

		if actual.CustomerId != v.Expected.CustomerId {
			t.Fatalf("Expected %q but got %q for CustomerId", v.Expected.CustomerId, actual.CustomerId)
		}

	}
}

func TestParseCustomerIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *CustomerId
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
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/cUsToMeRs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers/customerIdValue",
			Expected: &CustomerId{
				BillingAccountId: "billingAccountIdValue",
				CustomerId:       "customerIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.Billing/billingAccounts/billingAccountIdValue/customers/customerIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/cUsToMeRs/cUsToMeRiDvAlUe",
			Expected: &CustomerId{
				BillingAccountId: "bIlLiNgAcCoUnTiDvAlUe",
				CustomerId:       "cUsToMeRiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.bIlLiNg/bIlLiNgAcCoUnTs/bIlLiNgAcCoUnTiDvAlUe/cUsToMeRs/cUsToMeRiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseCustomerIDInsensitively(v.Input)
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

		if actual.CustomerId != v.Expected.CustomerId {
			t.Fatalf("Expected %q but got %q for CustomerId", v.Expected.CustomerId, actual.CustomerId)
		}

	}
}

func TestSegmentsForCustomerId(t *testing.T) {
	segments := CustomerId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("CustomerId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
