package mememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMemberOfId{}

// MeMemberOfId is a struct representing the Resource ID for a Me Member Of
type MeMemberOfId struct {
	DirectoryObjectId string
}

// NewMeMemberOfID returns a new MeMemberOfId struct
func NewMeMemberOfID(directoryObjectId string) MeMemberOfId {
	return MeMemberOfId{
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeMemberOfID parses 'input' into a MeMemberOfId
func ParseMeMemberOfID(input string) (*MeMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMemberOfId{}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseMeMemberOfIDInsensitively parses 'input' case-insensitively into a MeMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeMemberOfIDInsensitively(input string) (*MeMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMemberOfId{}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateMeMemberOfID checks that 'input' can be parsed as a Me Member Of ID
func ValidateMeMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Member Of ID
func (id MeMemberOfId) ID() string {
	fmtString := "/me/memberOf/%s"
	return fmt.Sprintf(fmtString, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Member Of ID
func (id MeMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMemberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Me Member Of ID
func (id MeMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Member Of (%s)", strings.Join(components, "\n"))
}
