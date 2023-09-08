package applicationhomerealmdiscoverypolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationHomeRealmDiscoveryPolicyId{}

// ApplicationHomeRealmDiscoveryPolicyId is a struct representing the Resource ID for a Application Home Realm Discovery Policy
type ApplicationHomeRealmDiscoveryPolicyId struct {
	ApplicationId              string
	HomeRealmDiscoveryPolicyId string
}

// NewApplicationHomeRealmDiscoveryPolicyID returns a new ApplicationHomeRealmDiscoveryPolicyId struct
func NewApplicationHomeRealmDiscoveryPolicyID(applicationId string, homeRealmDiscoveryPolicyId string) ApplicationHomeRealmDiscoveryPolicyId {
	return ApplicationHomeRealmDiscoveryPolicyId{
		ApplicationId:              applicationId,
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// ParseApplicationHomeRealmDiscoveryPolicyID parses 'input' into a ApplicationHomeRealmDiscoveryPolicyId
func ParseApplicationHomeRealmDiscoveryPolicyID(input string) (*ApplicationHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationHomeRealmDiscoveryPolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.HomeRealmDiscoveryPolicyId, ok = parsed.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", *parsed)
	}

	return &id, nil
}

// ParseApplicationHomeRealmDiscoveryPolicyIDInsensitively parses 'input' case-insensitively into a ApplicationHomeRealmDiscoveryPolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationHomeRealmDiscoveryPolicyIDInsensitively(input string) (*ApplicationHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationHomeRealmDiscoveryPolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.HomeRealmDiscoveryPolicyId, ok = parsed.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationHomeRealmDiscoveryPolicyID checks that 'input' can be parsed as a Application Home Realm Discovery Policy ID
func ValidateApplicationHomeRealmDiscoveryPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationHomeRealmDiscoveryPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Home Realm Discovery Policy ID
func (id ApplicationHomeRealmDiscoveryPolicyId) ID() string {
	fmtString := "/applications/%s/homeRealmDiscoveryPolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.HomeRealmDiscoveryPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Home Realm Discovery Policy ID
func (id ApplicationHomeRealmDiscoveryPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticHomeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyIdValue"),
	}
}

// String returns a human-readable description of this Application Home Realm Discovery Policy ID
func (id ApplicationHomeRealmDiscoveryPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
	}
	return fmt.Sprintf("Application Home Realm Discovery Policy (%s)", strings.Join(components, "\n"))
}
