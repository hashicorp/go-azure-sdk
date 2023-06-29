package denyassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DenyAssignmentIdId{}

// DenyAssignmentIdId is a struct representing the Resource ID for a Deny Assignment Id
type DenyAssignmentIdId struct {
	DenyAssignmentId string
}

// NewDenyAssignmentIdID returns a new DenyAssignmentIdId struct
func NewDenyAssignmentIdID(denyAssignmentId string) DenyAssignmentIdId {
	return DenyAssignmentIdId{
		DenyAssignmentId: denyAssignmentId,
	}
}

// ParseDenyAssignmentIdID parses 'input' into a DenyAssignmentIdId
func ParseDenyAssignmentIdID(input string) (*DenyAssignmentIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(DenyAssignmentIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DenyAssignmentIdId{}

	if id.DenyAssignmentId, ok = parsed.Parsed["denyAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "denyAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseDenyAssignmentIdIDInsensitively parses 'input' case-insensitively into a DenyAssignmentIdId
// note: this method should only be used for API response data and not user input
func ParseDenyAssignmentIdIDInsensitively(input string) (*DenyAssignmentIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(DenyAssignmentIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DenyAssignmentIdId{}

	if id.DenyAssignmentId, ok = parsed.Parsed["denyAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "denyAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateDenyAssignmentIdID checks that 'input' can be parsed as a Deny Assignment Id ID
func ValidateDenyAssignmentIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDenyAssignmentIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deny Assignment Id ID
func (id DenyAssignmentIdId) ID() string {
	fmtString := "/%s"
	return fmt.Sprintf(fmtString, id.DenyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Deny Assignment Id ID
func (id DenyAssignmentIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.UserSpecifiedSegment("denyAssignmentId", "denyAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Deny Assignment Id ID
func (id DenyAssignmentIdId) String() string {
	components := []string{
		fmt.Sprintf("Deny Assignment: %q", id.DenyAssignmentId),
	}
	return fmt.Sprintf("Deny Assignment Id (%s)", strings.Join(components, "\n"))
}
