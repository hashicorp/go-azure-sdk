package memanageddevicesecuritybaselinestate

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceSecurityBaselineStateId{}

func TestNewMeManagedDeviceSecurityBaselineStateID(t *testing.T) {
	id := NewMeManagedDeviceSecurityBaselineStateID("managedDeviceIdValue", "securityBaselineStateIdValue")

	if id.ManagedDeviceId != "managedDeviceIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ManagedDeviceId'", id.ManagedDeviceId, "managedDeviceIdValue")
	}

	if id.SecurityBaselineStateId != "securityBaselineStateIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'SecurityBaselineStateId'", id.SecurityBaselineStateId, "securityBaselineStateIdValue")
	}
}

func TestFormatMeManagedDeviceSecurityBaselineStateID(t *testing.T) {
	actual := NewMeManagedDeviceSecurityBaselineStateID("managedDeviceIdValue", "securityBaselineStateIdValue").ID()
	expected := "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeManagedDeviceSecurityBaselineStateID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceSecurityBaselineStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Expected: &MeManagedDeviceSecurityBaselineStateId{
				ManagedDeviceId:         "managedDeviceIdValue",
				SecurityBaselineStateId: "securityBaselineStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceSecurityBaselineStateID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

	}
}

func TestParseMeManagedDeviceSecurityBaselineStateIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeManagedDeviceSecurityBaselineStateId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue",
			Expected: &MeManagedDeviceSecurityBaselineStateId{
				ManagedDeviceId:         "managedDeviceIdValue",
				SecurityBaselineStateId: "securityBaselineStateIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/managedDevices/managedDeviceIdValue/securityBaselineStates/securityBaselineStateIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE",
			Expected: &MeManagedDeviceSecurityBaselineStateId{
				ManagedDeviceId:         "mAnAgEdDeViCeIdVaLuE",
				SecurityBaselineStateId: "sEcUrItYbAsElInEsTaTeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAnAgEdDeViCeS/mAnAgEdDeViCeIdVaLuE/sEcUrItYbAsElInEsTaTeS/sEcUrItYbAsElInEsTaTeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeManagedDeviceSecurityBaselineStateIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ManagedDeviceId != v.Expected.ManagedDeviceId {
			t.Fatalf("Expected %q but got %q for ManagedDeviceId", v.Expected.ManagedDeviceId, actual.ManagedDeviceId)
		}

		if actual.SecurityBaselineStateId != v.Expected.SecurityBaselineStateId {
			t.Fatalf("Expected %q but got %q for SecurityBaselineStateId", v.Expected.SecurityBaselineStateId, actual.SecurityBaselineStateId)
		}

	}
}

func TestSegmentsForMeManagedDeviceSecurityBaselineStateId(t *testing.T) {
	segments := MeManagedDeviceSecurityBaselineStateId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeManagedDeviceSecurityBaselineStateId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
