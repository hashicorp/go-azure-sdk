// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scheduledactions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScheduledActionId{}

func TestNewScheduledActionID(t *testing.T) {
	id := NewScheduledActionID("scheduledActionValue")

	if id.ScheduledActionName != "scheduledActionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ScheduledActionName'", id.ScheduledActionName, "scheduledActionValue")
	}
}

func TestFormatScheduledActionID(t *testing.T) {
	actual := NewScheduledActionID("scheduledActionValue").ID()
	expected := "/providers/Microsoft.CostManagement/scheduledActions/scheduledActionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseScheduledActionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScheduledActionId
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
			Input: "/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.CostManagement/scheduledActions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.CostManagement/scheduledActions/scheduledActionValue",
			Expected: &ScheduledActionId{
				ScheduledActionName: "scheduledActionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.CostManagement/scheduledActions/scheduledActionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScheduledActionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScheduledActionName != v.Expected.ScheduledActionName {
			t.Fatalf("Expected %q but got %q for ScheduledActionName", v.Expected.ScheduledActionName, actual.ScheduledActionName)
		}

	}
}

func TestParseScheduledActionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *ScheduledActionId
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
			Input: "/providers/Microsoft.CostManagement",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.CostManagement/scheduledActions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/sChEdUlEdAcTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.CostManagement/scheduledActions/scheduledActionValue",
			Expected: &ScheduledActionId{
				ScheduledActionName: "scheduledActionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.CostManagement/scheduledActions/scheduledActionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/sChEdUlEdAcTiOnS/sChEdUlEdAcTiOnVaLuE",
			Expected: &ScheduledActionId{
				ScheduledActionName: "sChEdUlEdAcTiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.cOsTmAnAgEmEnT/sChEdUlEdAcTiOnS/sChEdUlEdAcTiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseScheduledActionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ScheduledActionName != v.Expected.ScheduledActionName {
			t.Fatalf("Expected %q but got %q for ScheduledActionName", v.Expected.ScheduledActionName, actual.ScheduledActionName)
		}

	}
}

func TestSegmentsForScheduledActionId(t *testing.T) {
	segments := ScheduledActionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("ScheduledActionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
