package identityb2xuserflowuserattributeassignmentuserattribute

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowUserAttributeAssignmentId{}

// IdentityB2xUserFlowUserAttributeAssignmentId is a struct representing the Resource ID for a Identity B 2x User Flow User Attribute Assignment
type IdentityB2xUserFlowUserAttributeAssignmentId struct {
	B2xIdentityUserFlowId                 string
	IdentityUserFlowAttributeAssignmentId string
}

// NewIdentityB2xUserFlowUserAttributeAssignmentID returns a new IdentityB2xUserFlowUserAttributeAssignmentId struct
func NewIdentityB2xUserFlowUserAttributeAssignmentID(b2xIdentityUserFlowId string, identityUserFlowAttributeAssignmentId string) IdentityB2xUserFlowUserAttributeAssignmentId {
	return IdentityB2xUserFlowUserAttributeAssignmentId{
		B2xIdentityUserFlowId:                 b2xIdentityUserFlowId,
		IdentityUserFlowAttributeAssignmentId: identityUserFlowAttributeAssignmentId,
	}
}

// ParseIdentityB2xUserFlowUserAttributeAssignmentID parses 'input' into a IdentityB2xUserFlowUserAttributeAssignmentId
func ParseIdentityB2xUserFlowUserAttributeAssignmentID(input string) (*IdentityB2xUserFlowUserAttributeAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowUserAttributeAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowUserAttributeAssignmentId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.IdentityUserFlowAttributeAssignmentId, ok = parsed.Parsed["identityUserFlowAttributeAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowAttributeAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowUserAttributeAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowUserAttributeAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowUserAttributeAssignmentIDInsensitively(input string) (*IdentityB2xUserFlowUserAttributeAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowUserAttributeAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowUserAttributeAssignmentId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.IdentityUserFlowAttributeAssignmentId, ok = parsed.Parsed["identityUserFlowAttributeAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowAttributeAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityB2xUserFlowUserAttributeAssignmentID checks that 'input' can be parsed as a Identity B 2x User Flow User Attribute Assignment ID
func ValidateIdentityB2xUserFlowUserAttributeAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowUserAttributeAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow User Attribute Assignment ID
func (id IdentityB2xUserFlowUserAttributeAssignmentId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/userAttributeAssignments/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.IdentityUserFlowAttributeAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow User Attribute Assignment ID
func (id IdentityB2xUserFlowUserAttributeAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticB2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowIdValue"),
		resourceids.StaticSegment("staticUserAttributeAssignments", "userAttributeAssignments", "userAttributeAssignments"),
		resourceids.UserSpecifiedSegment("identityUserFlowAttributeAssignmentId", "identityUserFlowAttributeAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow User Attribute Assignment ID
func (id IdentityB2xUserFlowUserAttributeAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("Identity User Flow Attribute Assignment: %q", id.IdentityUserFlowAttributeAssignmentId),
	}
	return fmt.Sprintf("Identity B 2x User Flow User Attribute Assignment (%s)", strings.Join(components, "\n"))
}
