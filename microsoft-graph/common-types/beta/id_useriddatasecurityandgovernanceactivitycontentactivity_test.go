package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDataSecurityAndGovernanceActivityContentActivityId{}

func TestNewUserIdDataSecurityAndGovernanceActivityContentActivityID(t *testing.T) {
	id := NewUserIdDataSecurityAndGovernanceActivityContentActivityID("userId", "contentActivityId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.ContentActivityId != "contentActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentActivityId'", id.ContentActivityId, "contentActivityId")
	}
}

func TestFormatUserIdDataSecurityAndGovernanceActivityContentActivityID(t *testing.T) {
	actual := NewUserIdDataSecurityAndGovernanceActivityContentActivityID("userId", "contentActivityId").ID()
	expected := "/users/userId/dataSecurityAndGovernance/activities/contentActivities/contentActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDataSecurityAndGovernanceActivityContentActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDataSecurityAndGovernanceActivityContentActivityId
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
			Input: "/users/userId/dataSecurityAndGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/activities",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/activities/contentActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/dataSecurityAndGovernance/activities/contentActivities/contentActivityId",
			Expected: &UserIdDataSecurityAndGovernanceActivityContentActivityId{
				UserId:            "userId",
				ContentActivityId: "contentActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/dataSecurityAndGovernance/activities/contentActivities/contentActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDataSecurityAndGovernanceActivityContentActivityID(v.Input)
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

		if actual.ContentActivityId != v.Expected.ContentActivityId {
			t.Fatalf("Expected %q but got %q for ContentActivityId", v.Expected.ContentActivityId, actual.ContentActivityId)
		}

	}
}

func TestParseUserIdDataSecurityAndGovernanceActivityContentActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDataSecurityAndGovernanceActivityContentActivityId
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
			Input: "/users/userId/dataSecurityAndGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/activities/contentActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS/cOnTeNtAcTiViTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/dataSecurityAndGovernance/activities/contentActivities/contentActivityId",
			Expected: &UserIdDataSecurityAndGovernanceActivityContentActivityId{
				UserId:            "userId",
				ContentActivityId: "contentActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/dataSecurityAndGovernance/activities/contentActivities/contentActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS/cOnTeNtAcTiViTiEs/cOnTeNtAcTiViTyId",
			Expected: &UserIdDataSecurityAndGovernanceActivityContentActivityId{
				UserId:            "uSeRiD",
				ContentActivityId: "cOnTeNtAcTiViTyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS/cOnTeNtAcTiViTiEs/cOnTeNtAcTiViTyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDataSecurityAndGovernanceActivityContentActivityIDInsensitively(v.Input)
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

		if actual.ContentActivityId != v.Expected.ContentActivityId {
			t.Fatalf("Expected %q but got %q for ContentActivityId", v.Expected.ContentActivityId, actual.ContentActivityId)
		}

	}
}

func TestSegmentsForUserIdDataSecurityAndGovernanceActivityContentActivityId(t *testing.T) {
	segments := UserIdDataSecurityAndGovernanceActivityContentActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDataSecurityAndGovernanceActivityContentActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
