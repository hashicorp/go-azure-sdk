package mescopedrolememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeScopedRoleMemberOfId{}

// MeScopedRoleMemberOfId is a struct representing the Resource ID for a Me Scoped Role Member Of
type MeScopedRoleMemberOfId struct {
	ScopedRoleMembershipId string
}

// NewMeScopedRoleMemberOfID returns a new MeScopedRoleMemberOfId struct
func NewMeScopedRoleMemberOfID(scopedRoleMembershipId string) MeScopedRoleMemberOfId {
	return MeScopedRoleMemberOfId{
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseMeScopedRoleMemberOfID parses 'input' into a MeScopedRoleMemberOfId
func ParseMeScopedRoleMemberOfID(input string) (*MeScopedRoleMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeScopedRoleMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeScopedRoleMemberOfId{}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ParseMeScopedRoleMemberOfIDInsensitively parses 'input' case-insensitively into a MeScopedRoleMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeScopedRoleMemberOfIDInsensitively(input string) (*MeScopedRoleMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeScopedRoleMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeScopedRoleMemberOfId{}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ValidateMeScopedRoleMemberOfID checks that 'input' can be parsed as a Me Scoped Role Member Of ID
func ValidateMeScopedRoleMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeScopedRoleMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Scoped Role Member Of ID
func (id MeScopedRoleMemberOfId) ID() string {
	fmtString := "/me/scopedRoleMemberOf/%s"
	return fmt.Sprintf(fmtString, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Scoped Role Member Of ID
func (id MeScopedRoleMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticScopedRoleMemberOf", "scopedRoleMemberOf", "scopedRoleMemberOf"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipIdValue"),
	}
}

// String returns a human-readable description of this Me Scoped Role Member Of ID
func (id MeScopedRoleMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("Me Scoped Role Member Of (%s)", strings.Join(components, "\n"))
}
