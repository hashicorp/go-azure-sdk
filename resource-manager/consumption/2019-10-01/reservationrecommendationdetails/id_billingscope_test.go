package reservationrecommendationdetails

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = BillingScopeId{}

func TestNewBillingScopeID(t *testing.T) {
	id := NewBillingScopeID("billingScopeValue")

	if id.BillingScope != "billingScopeValue" {
		t.Fatalf("Expected %q but got %q for Segment 'BillingScope'", id.BillingScope, "billingScopeValue")
	}
}

func TestFormatBillingScopeID(t *testing.T) {
	actual := NewBillingScopeID("billingScopeValue").ID()
	expected := "/billingScopeValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseBillingScopeID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingScopeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/billingScopeValue",
			Expected: &BillingScopeId{
				BillingScope: "billingScopeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/billingScopeValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingScopeID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingScope != v.Expected.BillingScope {
			t.Fatalf("Expected %q but got %q for BillingScope", v.Expected.BillingScope, actual.BillingScope)
		}

	}
}

func TestParseBillingScopeIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *BillingScopeId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/billingScopeValue",
			Expected: &BillingScopeId{
				BillingScope: "billingScopeValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/billingScopeValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/bIlLiNgScOpEvAlUe",
			Expected: &BillingScopeId{
				BillingScope: "bIlLiNgScOpEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/bIlLiNgScOpEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseBillingScopeIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.BillingScope != v.Expected.BillingScope {
			t.Fatalf("Expected %q but got %q for BillingScope", v.Expected.BillingScope, actual.BillingScope)
		}

	}
}

func TestSegmentsForBillingScopeId(t *testing.T) {
	segments := BillingScopeId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("BillingScopeId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
