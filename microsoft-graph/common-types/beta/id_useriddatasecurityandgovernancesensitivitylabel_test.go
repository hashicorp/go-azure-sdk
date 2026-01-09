package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdDataSecurityAndGovernanceSensitivityLabelId{}

func TestNewUserIdDataSecurityAndGovernanceSensitivityLabelID(t *testing.T) {
	id := NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}
}

func TestFormatUserIdDataSecurityAndGovernanceSensitivityLabelID(t *testing.T) {
	actual := NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId").ID()
	expected := "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdDataSecurityAndGovernanceSensitivityLabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDataSecurityAndGovernanceSensitivityLabelId
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
			// Valid URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Expected: &UserIdDataSecurityAndGovernanceSensitivityLabelId{
				UserId:             "userId",
				SensitivityLabelId: "sensitivityLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDataSecurityAndGovernanceSensitivityLabelID(v.Input)
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

	}
}

func TestParseUserIdDataSecurityAndGovernanceSensitivityLabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdDataSecurityAndGovernanceSensitivityLabelId
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
			// Valid URI
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Expected: &UserIdDataSecurityAndGovernanceSensitivityLabelId{
				UserId:             "userId",
				SensitivityLabelId: "sensitivityLabelId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Expected: &UserIdDataSecurityAndGovernanceSensitivityLabelId{
				UserId:             "uSeRiD",
				SensitivityLabelId: "sEnSiTiViTyLaBeLiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdDataSecurityAndGovernanceSensitivityLabelIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForUserIdDataSecurityAndGovernanceSensitivityLabelId(t *testing.T) {
	segments := UserIdDataSecurityAndGovernanceSensitivityLabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdDataSecurityAndGovernanceSensitivityLabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
