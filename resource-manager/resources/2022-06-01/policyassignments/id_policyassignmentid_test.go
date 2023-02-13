package policyassignments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PolicyAssignmentIdId{}

func TestNewPolicyAssignmentIdID(t *testing.T) {
	id := NewPolicyAssignmentIdID("policyAssignmentIdValue")

	if id.PolicyAssignmentId != "policyAssignmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'PolicyAssignmentId'", id.PolicyAssignmentId, "policyAssignmentIdValue")
	}
}

func TestFormatPolicyAssignmentIdID(t *testing.T) {
	actual := NewPolicyAssignmentIdID("policyAssignmentIdValue").ID()
	expected := "/policyAssignmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParsePolicyAssignmentIdID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAssignmentIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policyAssignmentIdValue",
			Expected: &PolicyAssignmentIdId{
				PolicyAssignmentId: "policyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policyAssignmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAssignmentIdID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicyAssignmentId != v.Expected.PolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for PolicyAssignmentId", v.Expected.PolicyAssignmentId, actual.PolicyAssignmentId)
		}

	}
}

func TestParsePolicyAssignmentIdIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PolicyAssignmentIdId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Valid URI
			Input: "/policyAssignmentIdValue",
			Expected: &PolicyAssignmentIdId{
				PolicyAssignmentId: "policyAssignmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/policyAssignmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pOlIcYaSsIgNmEnTiDvAlUe",
			Expected: &PolicyAssignmentIdId{
				PolicyAssignmentId: "pOlIcYaSsIgNmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pOlIcYaSsIgNmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePolicyAssignmentIdIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.PolicyAssignmentId != v.Expected.PolicyAssignmentId {
			t.Fatalf("Expected %q but got %q for PolicyAssignmentId", v.Expected.PolicyAssignmentId, actual.PolicyAssignmentId)
		}

	}
}

func TestSegmentsForPolicyAssignmentIdId(t *testing.T) {
	segments := PolicyAssignmentIdId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("PolicyAssignmentIdId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
