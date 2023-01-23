package cognitiveservicesaccounts

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DeletedAccountId{}

func TestNewDeletedAccountID(t *testing.T) {
	id := NewDeletedAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "example-resource-group", "deletedAccountValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.LocationName != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'LocationName'", id.LocationName, "locationValue")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.DeletedAccountName != "deletedAccountValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DeletedAccountName'", id.DeletedAccountName, "deletedAccountValue")
	}
}

func TestFormatDeletedAccountID(t *testing.T) {
	actual := NewDeletedAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "example-resource-group", "deletedAccountValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts/deletedAccountValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDeletedAccountID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedAccountId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts/deletedAccountValue",
			Expected: &DeletedAccountId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "locationValue",
				ResourceGroupName:  "example-resource-group",
				DeletedAccountName: "deletedAccountValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts/deletedAccountValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedAccountID(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.DeletedAccountName != v.Expected.DeletedAccountName {
			t.Fatalf("Expected %q but got %q for DeletedAccountName", v.Expected.DeletedAccountName, actual.DeletedAccountName)
		}

	}
}

func TestParseDeletedAccountIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DeletedAccountId
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
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs/lOcAtIoNvAlUe/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs/lOcAtIoNvAlUe/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs/lOcAtIoNvAlUe/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/dElEtEdAcCoUnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts/deletedAccountValue",
			Expected: &DeletedAccountId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "locationValue",
				ResourceGroupName:  "example-resource-group",
				DeletedAccountName: "deletedAccountValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.CognitiveServices/locations/locationValue/resourceGroups/example-resource-group/deletedAccounts/deletedAccountValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs/lOcAtIoNvAlUe/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/dElEtEdAcCoUnTs/dElEtEdAcCoUnTvAlUe",
			Expected: &DeletedAccountId{
				SubscriptionId:     "12345678-1234-9876-4563-123456789012",
				LocationName:       "lOcAtIoNvAlUe",
				ResourceGroupName:  "eXaMpLe-rEsOuRcE-GrOuP",
				DeletedAccountName: "dElEtEdAcCoUnTvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOgNiTiVeSeRvIcEs/lOcAtIoNs/lOcAtIoNvAlUe/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/dElEtEdAcCoUnTs/dElEtEdAcCoUnTvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDeletedAccountIDInsensitively(v.Input)
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

		if actual.LocationName != v.Expected.LocationName {
			t.Fatalf("Expected %q but got %q for LocationName", v.Expected.LocationName, actual.LocationName)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.DeletedAccountName != v.Expected.DeletedAccountName {
			t.Fatalf("Expected %q but got %q for DeletedAccountName", v.Expected.DeletedAccountName, actual.DeletedAccountName)
		}

	}
}

func TestSegmentsForDeletedAccountId(t *testing.T) {
	segments := DeletedAccountId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DeletedAccountId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
