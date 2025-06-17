package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}

func TestNewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(t *testing.T) {
	id := NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}

	if id.SensitivityLabelId1 != "sensitivityLabelId1" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId1'", id.SensitivityLabelId1, "sensitivityLabelId1")
	}
}

func TestFormatUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(t *testing.T) {
	actual := NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1").ID()
	expected := "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId
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
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
				UserId:              "userId",
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID(v.Input)
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

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId
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
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
				UserId:              "userId",
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1",
			Expected: &UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
				UserId:              "uSeRiD",
				SensitivityLabelId:  "sEnSiTiViTyLaBeLiD",
				SensitivityLabelId1: "sEnSiTiViTyLaBeLiD1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively(v.Input)
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

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestSegmentsForUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId(t *testing.T) {
	segments := UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
