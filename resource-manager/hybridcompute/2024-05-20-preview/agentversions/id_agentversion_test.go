package agentversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AgentVersionId{}

func TestNewAgentVersionID(t *testing.T) {
	id := NewAgentVersionID("osTypeName", "agentVersionName")

	if id.OsTypeName != "osTypeName" {
		t.Fatalf("Expected %q but got %q for Segment 'OsTypeName'", id.OsTypeName, "osTypeName")
	}

	if id.AgentVersionName != "agentVersionName" {
		t.Fatalf("Expected %q but got %q for Segment 'AgentVersionName'", id.AgentVersionName, "agentVersionName")
	}
}

func TestFormatAgentVersionID(t *testing.T) {
	actual := NewAgentVersionID("osTypeName", "agentVersionName").ID()
	expected := "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions/agentVersionName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseAgentVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AgentVersionId
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
			Input: "/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/osType",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions/agentVersionName",
			Expected: &AgentVersionId{
				OsTypeName:       "osTypeName",
				AgentVersionName: "agentVersionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions/agentVersionName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAgentVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OsTypeName != v.Expected.OsTypeName {
			t.Fatalf("Expected %q but got %q for OsTypeName", v.Expected.OsTypeName, actual.OsTypeName)
		}

		if actual.AgentVersionName != v.Expected.AgentVersionName {
			t.Fatalf("Expected %q but got %q for AgentVersionName", v.Expected.AgentVersionName, actual.AgentVersionName)
		}

	}
}

func TestParseAgentVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AgentVersionId
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
			Input: "/providers/Microsoft.HybridCompute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/osType",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE/oStYpEnAmE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE/oStYpEnAmE/aGeNtVeRsIoNs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions/agentVersionName",
			Expected: &AgentVersionId{
				OsTypeName:       "osTypeName",
				AgentVersionName: "agentVersionName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.HybridCompute/osType/osTypeName/agentVersions/agentVersionName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE/oStYpEnAmE/aGeNtVeRsIoNs/aGeNtVeRsIoNnAmE",
			Expected: &AgentVersionId{
				OsTypeName:       "oStYpEnAmE",
				AgentVersionName: "aGeNtVeRsIoNnAmE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.hYbRiDcOmPuTe/oStYpE/oStYpEnAmE/aGeNtVeRsIoNs/aGeNtVeRsIoNnAmE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAgentVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OsTypeName != v.Expected.OsTypeName {
			t.Fatalf("Expected %q but got %q for OsTypeName", v.Expected.OsTypeName, actual.OsTypeName)
		}

		if actual.AgentVersionName != v.Expected.AgentVersionName {
			t.Fatalf("Expected %q but got %q for AgentVersionName", v.Expected.AgentVersionName, actual.AgentVersionName)
		}

	}
}

func TestSegmentsForAgentVersionId(t *testing.T) {
	segments := AgentVersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("AgentVersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
