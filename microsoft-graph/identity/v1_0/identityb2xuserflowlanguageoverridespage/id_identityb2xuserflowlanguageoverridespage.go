package identityb2xuserflowlanguageoverridespage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowLanguageOverridesPageId{}

// IdentityB2xUserFlowLanguageOverridesPageId is a struct representing the Resource ID for a Identity B 2x User Flow Language Overrides Page
type IdentityB2xUserFlowLanguageOverridesPageId struct {
	B2xIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
	UserFlowLanguagePageId          string
}

// NewIdentityB2xUserFlowLanguageOverridesPageID returns a new IdentityB2xUserFlowLanguageOverridesPageId struct
func NewIdentityB2xUserFlowLanguageOverridesPageID(b2xIdentityUserFlowId string, userFlowLanguageConfigurationId string, userFlowLanguagePageId string) IdentityB2xUserFlowLanguageOverridesPageId {
	return IdentityB2xUserFlowLanguageOverridesPageId{
		B2xIdentityUserFlowId:           b2xIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
		UserFlowLanguagePageId:          userFlowLanguagePageId,
	}
}

// ParseIdentityB2xUserFlowLanguageOverridesPageID parses 'input' into a IdentityB2xUserFlowLanguageOverridesPageId
func ParseIdentityB2xUserFlowLanguageOverridesPageID(input string) (*IdentityB2xUserFlowLanguageOverridesPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowLanguageOverridesPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowLanguageOverridesPageId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.UserFlowLanguageConfigurationId, ok = parsed.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", *parsed)
	}

	if id.UserFlowLanguagePageId, ok = parsed.Parsed["userFlowLanguagePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguagePageId", *parsed)
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowLanguageOverridesPageIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowLanguageOverridesPageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowLanguageOverridesPageIDInsensitively(input string) (*IdentityB2xUserFlowLanguageOverridesPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowLanguageOverridesPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowLanguageOverridesPageId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.UserFlowLanguageConfigurationId, ok = parsed.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", *parsed)
	}

	if id.UserFlowLanguagePageId, ok = parsed.Parsed["userFlowLanguagePageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguagePageId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityB2xUserFlowLanguageOverridesPageID checks that 'input' can be parsed as a Identity B 2x User Flow Language Overrides Page ID
func ValidateIdentityB2xUserFlowLanguageOverridesPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowLanguageOverridesPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow Language Overrides Page ID
func (id IdentityB2xUserFlowLanguageOverridesPageId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/languages/%s/overridesPages/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.UserFlowLanguageConfigurationId, id.UserFlowLanguagePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow Language Overrides Page ID
func (id IdentityB2xUserFlowLanguageOverridesPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticB2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowIdValue"),
		resourceids.StaticSegment("staticLanguages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationIdValue"),
		resourceids.StaticSegment("staticOverridesPages", "overridesPages", "overridesPages"),
		resourceids.UserSpecifiedSegment("userFlowLanguagePageId", "userFlowLanguagePageIdValue"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow Language Overrides Page ID
func (id IdentityB2xUserFlowLanguageOverridesPageId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
		fmt.Sprintf("User Flow Language Page: %q", id.UserFlowLanguagePageId),
	}
	return fmt.Sprintf("Identity B 2x User Flow Language Overrides Page (%s)", strings.Join(components, "\n"))
}
