package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCloudPCId{}

func TestNewUserIdCloudPCID(t *testing.T) {
	id := NewUserIdCloudPCID("userId", "cloudPCId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.CloudPCId != "cloudPCId" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCId'", id.CloudPCId, "cloudPCId")
	}
}

func TestFormatUserIdCloudPCID(t *testing.T) {
	actual := NewUserIdCloudPCID("userId", "cloudPCId").ID()
	expected := "/users/userId/cloudPCs/cloudPCId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdCloudPCID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCloudPCId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/cloudPCs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/cloudPCs/cloudPCId",
			Expected: &UserIdCloudPCId{
				UserId:    "userId",
				CloudPCId: "cloudPCId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/cloudPCs/cloudPCId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCloudPCID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.CloudPCId != v.Expected.CloudPCId {
			t.Fatalf("Expected %q but got %q for CloudPCId", v.Expected.CloudPCId, actual.CloudPCId)
		}

	}
}

func TestParseUserIdCloudPCIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdCloudPCId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/cloudPCs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdPcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/cloudPCs/cloudPCId",
			Expected: &UserIdCloudPCId{
				UserId:    "userId",
				CloudPCId: "cloudPCId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/cloudPCs/cloudPCId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdPcS/cLoUdPcId",
			Expected: &UserIdCloudPCId{
				UserId:    "uSeRiD",
				CloudPCId: "cLoUdPcId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/cLoUdPcS/cLoUdPcId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdCloudPCIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.CloudPCId != v.Expected.CloudPCId {
			t.Fatalf("Expected %q but got %q for CloudPCId", v.Expected.CloudPCId, actual.CloudPCId)
		}

	}
}

func TestSegmentsForUserIdCloudPCId(t *testing.T) {
	segments := UserIdCloudPCId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdCloudPCId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
