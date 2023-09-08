package identityb2xuserflowlanguagedefaultpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowLanguageDefaultPageId{}

// IdentityB2xUserFlowLanguageDefaultPageId is a struct representing the Resource ID for a Identity B 2x User Flow Language Default Page
type IdentityB2xUserFlowLanguageDefaultPageId struct {
	B2xIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
	UserFlowLanguagePageId          string
}

// NewIdentityB2xUserFlowLanguageDefaultPageID returns a new IdentityB2xUserFlowLanguageDefaultPageId struct
func NewIdentityB2xUserFlowLanguageDefaultPageID(b2xIdentityUserFlowId string, userFlowLanguageConfigurationId string, userFlowLanguagePageId string) IdentityB2xUserFlowLanguageDefaultPageId {
	return IdentityB2xUserFlowLanguageDefaultPageId{
		B2xIdentityUserFlowId:           b2xIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
		UserFlowLanguagePageId:          userFlowLanguagePageId,
	}
}

// ParseIdentityB2xUserFlowLanguageDefaultPageID parses 'input' into a IdentityB2xUserFlowLanguageDefaultPageId
func ParseIdentityB2xUserFlowLanguageDefaultPageID(input string) (*IdentityB2xUserFlowLanguageDefaultPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowLanguageDefaultPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowLanguageDefaultPageId{}

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

// ParseIdentityB2xUserFlowLanguageDefaultPageIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowLanguageDefaultPageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowLanguageDefaultPageIDInsensitively(input string) (*IdentityB2xUserFlowLanguageDefaultPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowLanguageDefaultPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowLanguageDefaultPageId{}

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

// ValidateIdentityB2xUserFlowLanguageDefaultPageID checks that 'input' can be parsed as a Identity B 2x User Flow Language Default Page ID
func ValidateIdentityB2xUserFlowLanguageDefaultPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowLanguageDefaultPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow Language Default Page ID
func (id IdentityB2xUserFlowLanguageDefaultPageId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/languages/%s/defaultPages/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.UserFlowLanguageConfigurationId, id.UserFlowLanguagePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow Language Default Page ID
func (id IdentityB2xUserFlowLanguageDefaultPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticB2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowIdValue"),
		resourceids.StaticSegment("staticLanguages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationIdValue"),
		resourceids.StaticSegment("staticDefaultPages", "defaultPages", "defaultPages"),
		resourceids.UserSpecifiedSegment("userFlowLanguagePageId", "userFlowLanguagePageIdValue"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow Language Default Page ID
func (id IdentityB2xUserFlowLanguageDefaultPageId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
		fmt.Sprintf("User Flow Language Page: %q", id.UserFlowLanguagePageId),
	}
	return fmt.Sprintf("Identity B 2x User Flow Language Default Page (%s)", strings.Join(components, "\n"))
}
