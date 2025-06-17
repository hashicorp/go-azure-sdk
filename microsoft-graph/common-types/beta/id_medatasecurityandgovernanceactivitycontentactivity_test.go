package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDataSecurityAndGovernanceActivityContentActivityId{}

func TestNewMeDataSecurityAndGovernanceActivityContentActivityID(t *testing.T) {
	id := NewMeDataSecurityAndGovernanceActivityContentActivityID("contentActivityId")

	if id.ContentActivityId != "contentActivityId" {
		t.Fatalf("Expected %q but got %q for Segment 'ContentActivityId'", id.ContentActivityId, "contentActivityId")
	}
}

func TestFormatMeDataSecurityAndGovernanceActivityContentActivityID(t *testing.T) {
	actual := NewMeDataSecurityAndGovernanceActivityContentActivityID("contentActivityId").ID()
	expected := "/me/dataSecurityAndGovernance/activities/contentActivities/contentActivityId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDataSecurityAndGovernanceActivityContentActivityID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDataSecurityAndGovernanceActivityContentActivityId
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
			Input: "/me/dataSecurityAndGovernance",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/activities",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/activities/contentActivities",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/dataSecurityAndGovernance/activities/contentActivities/contentActivityId",
			Expected: &MeDataSecurityAndGovernanceActivityContentActivityId{
				ContentActivityId: "contentActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/dataSecurityAndGovernance/activities/contentActivities/contentActivityId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDataSecurityAndGovernanceActivityContentActivityID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ContentActivityId != v.Expected.ContentActivityId {
			t.Fatalf("Expected %q but got %q for ContentActivityId", v.Expected.ContentActivityId, actual.ContentActivityId)
		}

	}
}

func TestParseMeDataSecurityAndGovernanceActivityContentActivityIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDataSecurityAndGovernanceActivityContentActivityId
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
			Input: "/me/dataSecurityAndGovernance",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/activities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/activities/contentActivities",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS/cOnTeNtAcTiViTiEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/dataSecurityAndGovernance/activities/contentActivities/contentActivityId",
			Expected: &MeDataSecurityAndGovernanceActivityContentActivityId{
				ContentActivityId: "contentActivityId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/dataSecurityAndGovernance/activities/contentActivities/contentActivityId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS/cOnTeNtAcTiViTiEs/cOnTeNtAcTiViTyId",
			Expected: &MeDataSecurityAndGovernanceActivityContentActivityId{
				ContentActivityId: "cOnTeNtAcTiViTyId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/aCtIvItIeS/cOnTeNtAcTiViTiEs/cOnTeNtAcTiViTyId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDataSecurityAndGovernanceActivityContentActivityIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ContentActivityId != v.Expected.ContentActivityId {
			t.Fatalf("Expected %q but got %q for ContentActivityId", v.Expected.ContentActivityId, actual.ContentActivityId)
		}

	}
}

func TestSegmentsForMeDataSecurityAndGovernanceActivityContentActivityId(t *testing.T) {
	segments := MeDataSecurityAndGovernanceActivityContentActivityId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDataSecurityAndGovernanceActivityContentActivityId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
