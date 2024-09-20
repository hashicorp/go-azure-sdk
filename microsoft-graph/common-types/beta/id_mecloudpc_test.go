package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCloudPCId{}

func TestNewMeCloudPCID(t *testing.T) {
	id := NewMeCloudPCID("cloudPCId")

	if id.CloudPCId != "cloudPCId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCId'", id.CloudPCId, "cloudPCId")
	}
}

func TestFormatMeCloudPCID(t *testing.T) {
	actual := NewMeCloudPCID("cloudPCId").ID()
	expected := "/me/cloudPCs/cloudPCId"
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
			Input: "/me/cloudPCs/cloudPCId",
			Expected: &MeCloudPCId{
				CloudPCId: "cloudPCId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/cloudPCs/cloudPCId/extra",
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
			Input: "/me/cloudPCs/cloudPCId",
			Expected: &MeCloudPCId{
				CloudPCId: "cloudPCId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/cloudPCs/cloudPCId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdPcS/cLoUdPcId",
			Expected: &MeCloudPCId{
				CloudPCId: "cLoUdPcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cLoUdPcS/cLoUdPcId/extra",
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
