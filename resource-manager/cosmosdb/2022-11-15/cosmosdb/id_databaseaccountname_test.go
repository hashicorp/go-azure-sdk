package cosmosdb

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DatabaseAccountNameId{}

func TestNewDatabaseAccountNameID(t *testing.T) {
	id := NewDatabaseAccountNameID("databaseAccountNameValue")

	if id.DatabaseAccountNameName != "databaseAccountNameValue" {
		t.Fatalf("Expected %q but got %q for Segment 'DatabaseAccountNameName'", id.DatabaseAccountNameName, "databaseAccountNameValue")
	}
}

func TestFormatDatabaseAccountNameID(t *testing.T) {
	actual := NewDatabaseAccountNameID("databaseAccountNameValue").ID()
	expected := "/providers/Microsoft.DocumentDB/databaseAccountNames/databaseAccountNameValue"
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
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/databaseAccountNameValue",
			Expected: &DatabaseAccountNameId{
				DatabaseAccountNameName: "databaseAccountNameValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/databaseAccountNameValue/extra",
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

		if actual.DatabaseAccountNameName != v.Expected.DatabaseAccountNameName {
			t.Fatalf("Expected %q but got %q for DatabaseAccountNameName", v.Expected.DatabaseAccountNameName, actual.DatabaseAccountNameName)
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
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/databaseAccountNameValue",
			Expected: &DatabaseAccountNameId{
				DatabaseAccountNameName: "databaseAccountNameValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/providers/Microsoft.DocumentDB/databaseAccountNames/databaseAccountNameValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dOcUmEnTdB/dAtAbAsEaCcOuNtNaMeS/dAtAbAsEaCcOuNtNaMeVaLuE",
			Expected: &DatabaseAccountNameId{
				DatabaseAccountNameName: "dAtAbAsEaCcOuNtNaMeVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/pRoViDeRs/mIcRoSoFt.dOcUmEnTdB/dAtAbAsEaCcOuNtNaMeS/dAtAbAsEaCcOuNtNaMeVaLuE/extra",
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

		if actual.DatabaseAccountNameName != v.Expected.DatabaseAccountNameName {
			t.Fatalf("Expected %q but got %q for DatabaseAccountNameName", v.Expected.DatabaseAccountNameName, actual.DatabaseAccountNameName)
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
