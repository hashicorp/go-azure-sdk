package identityb2xuserflowlanguagedefaultpage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowLanguageId{}

// IdentityB2xUserFlowLanguageId is a struct representing the Resource ID for a Identity B 2x User Flow Language
type IdentityB2xUserFlowLanguageId struct {
	B2xIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
}

// NewIdentityB2xUserFlowLanguageID returns a new IdentityB2xUserFlowLanguageId struct
func NewIdentityB2xUserFlowLanguageID(b2xIdentityUserFlowId string, userFlowLanguageConfigurationId string) IdentityB2xUserFlowLanguageId {
	return IdentityB2xUserFlowLanguageId{
		B2xIdentityUserFlowId:           b2xIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
	}
}

// ParseIdentityB2xUserFlowLanguageID parses 'input' into a IdentityB2xUserFlowLanguageId
func ParseIdentityB2xUserFlowLanguageID(input string) (*IdentityB2xUserFlowLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowLanguageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowLanguageId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.UserFlowLanguageConfigurationId, ok = parsed.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", *parsed)
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowLanguageIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowLanguageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowLanguageIDInsensitively(input string) (*IdentityB2xUserFlowLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowLanguageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowLanguageId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.UserFlowLanguageConfigurationId, ok = parsed.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityB2xUserFlowLanguageID checks that 'input' can be parsed as a Identity B 2x User Flow Language ID
func ValidateIdentityB2xUserFlowLanguageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowLanguageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow Language ID
func (id IdentityB2xUserFlowLanguageId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/languages/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.UserFlowLanguageConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow Language ID
func (id IdentityB2xUserFlowLanguageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticB2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowIdValue"),
		resourceids.StaticSegment("staticLanguages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationIdValue"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow Language ID
func (id IdentityB2xUserFlowLanguageId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
	}
	return fmt.Sprintf("Identity B 2x User Flow Language (%s)", strings.Join(components, "\n"))
}
