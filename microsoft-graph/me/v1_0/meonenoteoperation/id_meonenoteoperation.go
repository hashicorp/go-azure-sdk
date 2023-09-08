package meonenoteoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnenoteOperationId{}

// MeOnenoteOperationId is a struct representing the Resource ID for a Me Onenote Operation
type MeOnenoteOperationId struct {
	OnenoteOperationId string
}

// NewMeOnenoteOperationID returns a new MeOnenoteOperationId struct
func NewMeOnenoteOperationID(onenoteOperationId string) MeOnenoteOperationId {
	return MeOnenoteOperationId{
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseMeOnenoteOperationID parses 'input' into a MeOnenoteOperationId
func ParseMeOnenoteOperationID(input string) (*MeOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteOperationId{}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ParseMeOnenoteOperationIDInsensitively parses 'input' case-insensitively into a MeOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseMeOnenoteOperationIDInsensitively(input string) (*MeOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnenoteOperationId{}

	if id.OnenoteOperationId, ok = parsed.Parsed["onenoteOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnenoteOperationID checks that 'input' can be parsed as a Me Onenote Operation ID
func ValidateMeOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Onenote Operation ID
func (id MeOnenoteOperationId) ID() string {
	fmtString := "/me/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Onenote Operation ID
func (id MeOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnenote", "onenote", "onenote"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationIdValue"),
	}
}

// String returns a human-readable description of this Me Onenote Operation ID
func (id MeOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("Me Onenote Operation (%s)", strings.Join(components, "\n"))
}
