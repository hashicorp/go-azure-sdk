package cosmosdb

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DatabaseAccountNameId{}

func TestNewDatabaseAccountNameID(t *testing.T) {
	id := NewDatabaseAccountNameID("accountName")

	if id.DatabaseAccountName != "accountName" {
		t.Fatalf("Expected %q but got %q for Segment 'DatabaseAccountName'", id.DatabaseAccountName, "accountName")
	}
}

func TestFormatDatabaseAccountNameID(t *testing.T) {
	actual := NewDatabaseAccountNameID("accountName").ID()
	expected := "/providers/Microsoft.DocumentDB/databaseAccountNames/accountName"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseDatabaseAccountNameID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DatabaseAccountNameId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DocumentDB",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/accountName",
			Expected: &DatabaseAccountNameId{
				DatabaseAccountName: "accountName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/accountName/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDatabaseAccountNameID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DatabaseAccountName != v.Expected.DatabaseAccountName {
			t.Fatalf("Expected %q but got %q for DatabaseAccountName", v.Expected.DatabaseAccountName, actual.DatabaseAccountName)
		}

	}
}

func TestParseDatabaseAccountNameIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *DatabaseAccountNameId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DocumentDB",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dOcUmEnTdB",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dOcUmEnTdB/dAtAbAsEaCcOuNtNaMeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/accountName",
			Expected: &DatabaseAccountNameId{
				DatabaseAccountName: "accountName",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/accountName/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dOcUmEnTdB/dAtAbAsEaCcOuNtNaMeS/aCcOuNtNaMe",
			Expected: &DatabaseAccountNameId{
				DatabaseAccountName: "aCcOuNtNaMe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dOcUmEnTdB/dAtAbAsEaCcOuNtNaMeS/aCcOuNtNaMe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseDatabaseAccountNameIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.DatabaseAccountName != v.Expected.DatabaseAccountName {
			t.Fatalf("Expected %q but got %q for DatabaseAccountName", v.Expected.DatabaseAccountName, actual.DatabaseAccountName)
		}

	}
}

func TestSegmentsForDatabaseAccountNameId(t *testing.T) {
	segments := DatabaseAccountNameId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("DatabaseAccountNameId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
