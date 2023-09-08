package applicationtokenissuancepolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationTokenIssuancePolicyId{}

// ApplicationTokenIssuancePolicyId is a struct representing the Resource ID for a Application Token Issuance Policy
type ApplicationTokenIssuancePolicyId struct {
	ApplicationId         string
	TokenIssuancePolicyId string
}

// NewApplicationTokenIssuancePolicyID returns a new ApplicationTokenIssuancePolicyId struct
func NewApplicationTokenIssuancePolicyID(applicationId string, tokenIssuancePolicyId string) ApplicationTokenIssuancePolicyId {
	return ApplicationTokenIssuancePolicyId{
		ApplicationId:         applicationId,
		TokenIssuancePolicyId: tokenIssuancePolicyId,
	}
}

// ParseApplicationTokenIssuancePolicyID parses 'input' into a ApplicationTokenIssuancePolicyId
func ParseApplicationTokenIssuancePolicyID(input string) (*ApplicationTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationTokenIssuancePolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	return &id, nil
}

// ParseApplicationTokenIssuancePolicyIDInsensitively parses 'input' case-insensitively into a ApplicationTokenIssuancePolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationTokenIssuancePolicyIDInsensitively(input string) (*ApplicationTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationTokenIssuancePolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationTokenIssuancePolicyID checks that 'input' can be parsed as a Application Token Issuance Policy ID
func ValidateApplicationTokenIssuancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationTokenIssuancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Token Issuance Policy ID
func (id ApplicationTokenIssuancePolicyId) ID() string {
	fmtString := "/applications/%s/tokenIssuancePolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.TokenIssuancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Token Issuance Policy ID
func (id ApplicationTokenIssuancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticTokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyIdValue"),
	}
}

// String returns a human-readable description of this Application Token Issuance Policy ID
func (id ApplicationTokenIssuancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
	}
	return fmt.Sprintf("Application Token Issuance Policy (%s)", strings.Join(components, "\n"))
}
