package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCloudPCId{}

func TestNewUserIdCloudPCID(t *testing.T) {
	id := NewUserIdCloudPCID("userIdValue", "cloudPCIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.CloudPCId != "cloudPCIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'CloudPCId'", id.CloudPCId, "cloudPCIdValue")
	}
}

func TestFormatUserIdCloudPCID(t *testing.T) {
	actual := NewUserIdCloudPCID("userIdValue", "cloudPCIdValue").ID()
	expected := "/users/userIdValue/cloudPCs/cloudPCIdValue"
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/cloudPCs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/cloudPCs/cloudPCIdValue",
			Expected: &UserIdCloudPCId{
				UserId:    "userIdValue",
				CloudPCId: "cloudPCIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/cloudPCs/cloudPCIdValue/extra",
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/cloudPCs",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdPcS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/cloudPCs/cloudPCIdValue",
			Expected: &UserIdCloudPCId{
				UserId:    "userIdValue",
				CloudPCId: "cloudPCIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/cloudPCs/cloudPCIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdPcS/cLoUdPcIdVaLuE",
			Expected: &UserIdCloudPCId{
				UserId:    "uSeRiDvAlUe",
				CloudPCId: "cLoUdPcIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cLoUdPcS/cLoUdPcIdVaLuE/extra",
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
