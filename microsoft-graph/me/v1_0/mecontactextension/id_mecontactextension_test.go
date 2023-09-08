package mecontactextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeContactExtensionId{}

func TestNewMeContactExtensionID(t *testing.T) {
	id := NewMeContactExtensionID("contactIdValue", "extensionIdValue")

	if id.ContactId != "contactIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ContactId'", id.ContactId, "contactIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatMeContactExtensionID(t *testing.T) {
	actual := NewMeContactExtensionID("contactIdValue", "extensionIdValue").ID()
	expected := "/me/contacts/contactIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeContactExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactExtensionId
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
			Input: "/me/contacts",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contacts/contactIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contacts/contactIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contacts/contactIdValue/extensions/extensionIdValue",
			Expected: &MeContactExtensionId{
				ContactId:   "contactIdValue",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contacts/contactIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactExtensionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ContactId != v.Expected.ContactId {
			t.Fatalf("Expected %q but got %q for ContactId", v.Expected.ContactId, actual.ContactId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseMeContactExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeContactExtensionId
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
			Input: "/me/contacts",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contacts/contactIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtS/cOnTaCtIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/contacts/contactIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/contacts/contactIdValue/extensions/extensionIdValue",
			Expected: &MeContactExtensionId{
				ContactId:   "contactIdValue",
				ExtensionId: "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/contacts/contactIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &MeContactExtensionId{
				ContactId:   "cOnTaCtIdVaLuE",
				ExtensionId: "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/cOnTaCtS/cOnTaCtIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeContactExtensionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.ContactId != v.Expected.ContactId {
			t.Fatalf("Expected %q but got %q for ContactId", v.Expected.ContactId, actual.ContactId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForMeContactExtensionId(t *testing.T) {
	segments := MeContactExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeContactExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
