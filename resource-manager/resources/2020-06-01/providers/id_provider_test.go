package providers

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ProviderId{}

func TestNewProviderID(t *testing.T) {
	id := NewProviderID("resourceProviderNamespaceValue")

	if id.ResourceProviderNamespace != "resourceProviderNamespaceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceProviderNamespace'", id.ResourceProviderNamespace, "resourceProviderNamespaceValue")
	}
}

func TestFormatProviderID(t *testing.T) {
	actual := NewProviderID("resourceProviderNamespaceValue").ID()
	expected := "/providers/resourceProviderNamespaceValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseProviderID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderId
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
			// Valid URI
			Input: "/providers/resourceProviderNamespaceValue",
			Expected: &ProviderId{
				ResourceProviderNamespace: "resourceProviderNamespaceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/resourceProviderNamespaceValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceProviderNamespace != v.Expected.ResourceProviderNamespace {
			t.Fatalf("Expected %q but got %q for ResourceProviderNamespace", v.Expected.ResourceProviderNamespace, actual.ResourceProviderNamespace)
		}

	}
}

func TestParseProviderIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ProviderId
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
			// Valid URI
			Input: "/providers/resourceProviderNamespaceValue",
			Expected: &ProviderId{
				ResourceProviderNamespace: "resourceProviderNamespaceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/resourceProviderNamespaceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/rEsOuRcEpRoViDeRnAmEsPaCeVaLuE",
			Expected: &ProviderId{
				ResourceProviderNamespace: "rEsOuRcEpRoViDeRnAmEsPaCeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/rEsOuRcEpRoViDeRnAmEsPaCeVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseProviderIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ResourceProviderNamespace != v.Expected.ResourceProviderNamespace {
			t.Fatalf("Expected %q but got %q for ResourceProviderNamespace", v.Expected.ResourceProviderNamespace, actual.ResourceProviderNamespace)
		}

	}
}

func TestSegmentsForProviderId(t *testing.T) {
	segments := ProviderId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ProviderId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
