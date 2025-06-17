package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}

func TestNewMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(t *testing.T) {
	id := NewMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1")

	if id.SensitivityLabelId != "sensitivityLabelId" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId'", id.SensitivityLabelId, "sensitivityLabelId")
	}

	if id.SensitivityLabelId1 != "sensitivityLabelId1" {
		t.Fatalf("Expected %q but got %q for Segment 'SensitivityLabelId1'", id.SensitivityLabelId1, "sensitivityLabelId1")
	}
}

func TestFormatMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(t *testing.T) {
	actual := NewMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1").ID()
	expected := "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId
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
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId
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
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1",
			Expected: &MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
				SensitivityLabelId:  "sensitivityLabelId",
				SensitivityLabelId1: "sensitivityLabelId1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/dataSecurityAndGovernance/sensitivityLabels/sensitivityLabelId/sublabels/sensitivityLabelId1/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1",
			Expected: &MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{
				SensitivityLabelId:  "sEnSiTiViTyLaBeLiD",
				SensitivityLabelId1: "sEnSiTiViTyLaBeLiD1",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/dAtAsEcUrItYaNdGoVeRnAnCe/sEnSiTiViTyLaBeLs/sEnSiTiViTyLaBeLiD/sUbLaBeLs/sEnSiTiViTyLaBeLiD1/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeDataSecurityAndGovernanceSensitivityLabelIdSublabelIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SensitivityLabelId != v.Expected.SensitivityLabelId {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId", v.Expected.SensitivityLabelId, actual.SensitivityLabelId)
		}

		if actual.SensitivityLabelId1 != v.Expected.SensitivityLabelId1 {
			t.Fatalf("Expected %q but got %q for SensitivityLabelId1", v.Expected.SensitivityLabelId1, actual.SensitivityLabelId1)
		}

	}
}

func TestSegmentsForMeDataSecurityAndGovernanceSensitivityLabelIdSublabelId(t *testing.T) {
	segments := MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeDataSecurityAndGovernanceSensitivityLabelIdSublabelId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
