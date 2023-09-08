package applicationtokenlifetimepolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationTokenLifetimePolicyId{}

// ApplicationTokenLifetimePolicyId is a struct representing the Resource ID for a Application Token Lifetime Policy
type ApplicationTokenLifetimePolicyId struct {
	ApplicationId         string
	TokenLifetimePolicyId string
}

// NewApplicationTokenLifetimePolicyID returns a new ApplicationTokenLifetimePolicyId struct
func NewApplicationTokenLifetimePolicyID(applicationId string, tokenLifetimePolicyId string) ApplicationTokenLifetimePolicyId {
	return ApplicationTokenLifetimePolicyId{
		ApplicationId:         applicationId,
		TokenLifetimePolicyId: tokenLifetimePolicyId,
	}
}

// ParseApplicationTokenLifetimePolicyID parses 'input' into a ApplicationTokenLifetimePolicyId
func ParseApplicationTokenLifetimePolicyID(input string) (*ApplicationTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationTokenLifetimePolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.TokenLifetimePolicyId, ok = parsed.Parsed["tokenLifetimePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", *parsed)
	}

	return &id, nil
}

// ParseApplicationTokenLifetimePolicyIDInsensitively parses 'input' case-insensitively into a ApplicationTokenLifetimePolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationTokenLifetimePolicyIDInsensitively(input string) (*ApplicationTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationTokenLifetimePolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.TokenLifetimePolicyId, ok = parsed.Parsed["tokenLifetimePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationTokenLifetimePolicyID checks that 'input' can be parsed as a Application Token Lifetime Policy ID
func ValidateApplicationTokenLifetimePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationTokenLifetimePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Token Lifetime Policy ID
func (id ApplicationTokenLifetimePolicyId) ID() string {
	fmtString := "/applications/%s/tokenLifetimePolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.TokenLifetimePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Token Lifetime Policy ID
func (id ApplicationTokenLifetimePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticTokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyIdValue"),
	}
}

// String returns a human-readable description of this Application Token Lifetime Policy ID
func (id ApplicationTokenLifetimePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
	}
	return fmt.Sprintf("Application Token Lifetime Policy (%s)", strings.Join(components, "\n"))
}
