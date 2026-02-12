package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &EnrollmentGroupId{}

// EnrollmentGroupId is a struct representing the Resource ID for a Enrollment Group
type EnrollmentGroupId struct {
	BaseURI           string
	EnrollmentGroupId string
}

// NewEnrollmentGroupID returns a new EnrollmentGroupId struct
func NewEnrollmentGroupID(baseURI string, enrollmentGroupId string) EnrollmentGroupId {
	return EnrollmentGroupId{
		BaseURI:           strings.TrimSuffix(baseURI, "/"),
		EnrollmentGroupId: enrollmentGroupId,
	}
}

// ParseEnrollmentGroupID parses 'input' into a EnrollmentGroupId
func ParseEnrollmentGroupID(input string) (*EnrollmentGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnrollmentGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnrollmentGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnrollmentGroupIDInsensitively parses 'input' case-insensitively into a EnrollmentGroupId
// note: this method should only be used for API response data and not user input
func ParseEnrollmentGroupIDInsensitively(input string) (*EnrollmentGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnrollmentGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnrollmentGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnrollmentGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.EnrollmentGroupId, ok = input.Parsed["enrollmentGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentGroupId", input)
	}

	return nil
}

// ValidateEnrollmentGroupID checks that 'input' can be parsed as a Enrollment Group ID
func ValidateEnrollmentGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnrollmentGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Enrollment Group ID
func (id EnrollmentGroupId) ID() string {
	fmtString := "%s/enrollmentGroups/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.EnrollmentGroupId)
}

// Path returns the formatted Enrollment Group ID without the BaseURI
func (id EnrollmentGroupId) Path() string {
	fmtString := "/enrollmentGroups/%s"
	return fmt.Sprintf(fmtString, id.EnrollmentGroupId)
}

// PathElements returns the values of Enrollment Group ID Segments without the BaseURI
func (id EnrollmentGroupId) PathElements() []any {
	return []any{id.EnrollmentGroupId}
}

// Segments returns a slice of Resource ID Segments which comprise this Enrollment Group ID
func (id EnrollmentGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticEnrollmentGroups", "enrollmentGroups", "enrollmentGroups"),
		resourceids.UserSpecifiedSegment("enrollmentGroupId", "enrollmentGroupId"),
	}
}

// String returns a human-readable description of this Enrollment Group ID
func (id EnrollmentGroupId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Enrollment Group: %q", id.EnrollmentGroupId),
	}
	return fmt.Sprintf("Enrollment Group (%s)", strings.Join(components, "\n"))
}
