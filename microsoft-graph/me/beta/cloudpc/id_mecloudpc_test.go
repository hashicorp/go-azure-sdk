package cloudpc

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCloudPCId{}

func TestNewMeCloudPCID(t *testing.T) {
	id := NewMeCloudPCID("cloudPCIdValue")

	if id.CloudPCId != "cloudPCIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCId'", id.CloudPCId, "cloudPCIdValue")
	}
}

func TestFormatMeCloudPCID(t *testing.T) {
	actual := NewMeCloudPCID("cloudPCIdValue").ID()
	expected := "/me/cloudPCs/cloudPCIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeCloudPCID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCloudPCId
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
			Input: "/me/cloudPCs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/cloudPCs/cloudPCIdValue",
			Expected: &MeCloudPCId{
				CloudPCId: "cloudPCIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/cloudPCs/cloudPCIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCloudPCID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCId != v.Expected.CloudPCId {
			t.Fatalf("Expected %q but got %q for CloudPCId", v.Expected.CloudPCId, actual.CloudPCId)
		}

	}
}

func TestParseMeCloudPCIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeCloudPCId
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
			Input: "/me/cloudPCs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdPcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/cloudPCs/cloudPCIdValue",
			Expected: &MeCloudPCId{
				CloudPCId: "cloudPCIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/cloudPCs/cloudPCIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdPcS/cLoUdPcIdVaLuE",
			Expected: &MeCloudPCId{
				CloudPCId: "cLoUdPcIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdPcS/cLoUdPcIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeCloudPCIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.CloudPCId != v.Expected.CloudPCId {
			t.Fatalf("Expected %q but got %q for CloudPCId", v.Expected.CloudPCId, actual.CloudPCId)
		}

	}
}

func TestSegmentsForMeCloudPCId(t *testing.T) {
	segments := MeCloudPCId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeCloudPCId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
