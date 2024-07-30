package conditionalaccespolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccesPolicyId{}

// IdentityConditionalAccesPolicyId is a struct representing the Resource ID for a Identity Conditional Acces Policy
type IdentityConditionalAccesPolicyId struct {
	ConditionalAccessPolicyId string
}

// NewIdentityConditionalAccesPolicyID returns a new IdentityConditionalAccesPolicyId struct
func NewIdentityConditionalAccesPolicyID(conditionalAccessPolicyId string) IdentityConditionalAccesPolicyId {
	return IdentityConditionalAccesPolicyId{
		ConditionalAccessPolicyId: conditionalAccessPolicyId,
	}
}

// ParseIdentityConditionalAccesPolicyID parses 'input' into a IdentityConditionalAccesPolicyId
func ParseIdentityConditionalAccesPolicyID(input string) (*IdentityConditionalAccesPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccesPolicyIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesPolicyIDInsensitively(input string) (*IdentityConditionalAccesPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccesPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ConditionalAccessPolicyId, ok = input.Parsed["conditionalAccessPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conditionalAccessPolicyId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccesPolicyID checks that 'input' can be parsed as a Identity Conditional Acces Policy ID
func ValidateIdentityConditionalAccesPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Policy ID
func (id IdentityConditionalAccesPolicyId) ID() string {
	fmtString := "/identity/conditionalAccess/policies/%s"
	return fmt.Sprintf(fmtString, id.ConditionalAccessPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Policy ID
func (id IdentityConditionalAccesPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("conditionalAccessPolicyId", "conditionalAccessPolicyIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Policy ID
func (id IdentityConditionalAccesPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Conditional Access Policy: %q", id.ConditionalAccessPolicyId),
	}
	return fmt.Sprintf("Identity Conditional Acces Policy (%s)", strings.Join(components, "\n"))
}
