package identityconditionalaccestemplate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesTemplateId{}

// IdentityConditionalAccesTemplateId is a struct representing the Resource ID for a Identity Conditional Acces Template
type IdentityConditionalAccesTemplateId struct {
	ConditionalAccessTemplateId string
}

// NewIdentityConditionalAccesTemplateID returns a new IdentityConditionalAccesTemplateId struct
func NewIdentityConditionalAccesTemplateID(conditionalAccessTemplateId string) IdentityConditionalAccesTemplateId {
	return IdentityConditionalAccesTemplateId{
		ConditionalAccessTemplateId: conditionalAccessTemplateId,
	}
}

// ParseIdentityConditionalAccesTemplateID parses 'input' into a IdentityConditionalAccesTemplateId
func ParseIdentityConditionalAccesTemplateID(input string) (*IdentityConditionalAccesTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesTemplateId{}

	if id.ConditionalAccessTemplateId, ok = parsed.Parsed["conditionalAccessTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conditionalAccessTemplateId", *parsed)
	}

	return &id, nil
}

// ParseIdentityConditionalAccesTemplateIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesTemplateId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesTemplateIDInsensitively(input string) (*IdentityConditionalAccesTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesTemplateId{}

	if id.ConditionalAccessTemplateId, ok = parsed.Parsed["conditionalAccessTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conditionalAccessTemplateId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityConditionalAccesTemplateID checks that 'input' can be parsed as a Identity Conditional Acces Template ID
func ValidateIdentityConditionalAccesTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Template ID
func (id IdentityConditionalAccesTemplateId) ID() string {
	fmtString := "/identity/conditionalAccess/templates/%s"
	return fmt.Sprintf(fmtString, id.ConditionalAccessTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Template ID
func (id IdentityConditionalAccesTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticConditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("staticTemplates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("conditionalAccessTemplateId", "conditionalAccessTemplateIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Template ID
func (id IdentityConditionalAccesTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Conditional Access Template: %q", id.ConditionalAccessTemplateId),
	}
	return fmt.Sprintf("Identity Conditional Acces Template (%s)", strings.Join(components, "\n"))
}
