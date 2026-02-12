package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EntryId{}

// EntryId is a struct representing the Resource ID for a Entry
type EntryId struct {
	BaseURI           string
	EnrollmentGroupId string
	Entry             CertificateEntry
}

// NewEntryID returns a new EntryId struct
func NewEntryID(baseURI string, enrollmentGroupId string, entry CertificateEntry) EntryId {
	return EntryId{
		BaseURI:           strings.TrimSuffix(baseURI, "/"),
		EnrollmentGroupId: enrollmentGroupId,
		Entry:             entry,
	}
}

// ParseEntryID parses 'input' into a EntryId
func ParseEntryID(input string) (*EntryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EntryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EntryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEntryIDInsensitively parses 'input' case-insensitively into a EntryId
// note: this method should only be used for API response data and not user input
func ParseEntryIDInsensitively(input string) (*EntryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EntryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EntryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EntryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.EnrollmentGroupId, ok = input.Parsed["enrollmentGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentGroupId", input)
	}

	if v, ok := input.Parsed["entry"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "entry", input)
		}

		entry, err := parseCertificateEntry(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.Entry = *entry
	}

	return nil
}

// ValidateEntryID checks that 'input' can be parsed as a Entry ID
func ValidateEntryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEntryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Entry ID
func (id EntryId) ID() string {
	fmtString := "%s/enrollmentGroups/%s/certificates/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.EnrollmentGroupId, string(id.Entry))
}

// Path returns the formatted Entry ID without the BaseURI
func (id EntryId) Path() string {
	fmtString := "/enrollmentGroups/%s/certificates/%s"
	return fmt.Sprintf(fmtString, id.EnrollmentGroupId, string(id.Entry))
}

// PathElements returns the values of Entry ID Segments without the BaseURI
func (id EntryId) PathElements() []any {
	return []any{id.EnrollmentGroupId, string(id.Entry)}
}

// Segments returns a slice of Resource ID Segments which comprise this Entry ID
func (id EntryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticEnrollmentGroups", "enrollmentGroups", "enrollmentGroups"),
		resourceids.UserSpecifiedSegment("enrollmentGroupId", "enrollmentGroupId"),
		resourceids.StaticSegment("staticCertificates", "certificates", "certificates"),
		resourceids.ConstantSegment("entry", PossibleValuesForCertificateEntry(), "primary"),
	}
}

// String returns a human-readable description of this Entry ID
func (id EntryId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Enrollment Group: %q", id.EnrollmentGroupId),
		fmt.Sprintf("Entry: %q", string(id.Entry)),
	}
	return fmt.Sprintf("Entry (%s)", strings.Join(components, "\n"))
}
