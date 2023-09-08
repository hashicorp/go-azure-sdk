package applicationappmanagementpolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationAppManagementPolicyId{}

// ApplicationAppManagementPolicyId is a struct representing the Resource ID for a Application App Management Policy
type ApplicationAppManagementPolicyId struct {
	ApplicationId         string
	AppManagementPolicyId string
}

// NewApplicationAppManagementPolicyID returns a new ApplicationAppManagementPolicyId struct
func NewApplicationAppManagementPolicyID(applicationId string, appManagementPolicyId string) ApplicationAppManagementPolicyId {
	return ApplicationAppManagementPolicyId{
		ApplicationId:         applicationId,
		AppManagementPolicyId: appManagementPolicyId,
	}
}

// ParseApplicationAppManagementPolicyID parses 'input' into a ApplicationAppManagementPolicyId
func ParseApplicationAppManagementPolicyID(input string) (*ApplicationAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationAppManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationAppManagementPolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.AppManagementPolicyId, ok = parsed.Parsed["appManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", *parsed)
	}

	return &id, nil
}

// ParseApplicationAppManagementPolicyIDInsensitively parses 'input' case-insensitively into a ApplicationAppManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationAppManagementPolicyIDInsensitively(input string) (*ApplicationAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationAppManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationAppManagementPolicyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.AppManagementPolicyId, ok = parsed.Parsed["appManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationAppManagementPolicyID checks that 'input' can be parsed as a Application App Management Policy ID
func ValidateApplicationAppManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationAppManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application App Management Policy ID
func (id ApplicationAppManagementPolicyId) ID() string {
	fmtString := "/applications/%s/appManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.AppManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application App Management Policy ID
func (id ApplicationAppManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticAppManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyIdValue"),
	}
}

// String returns a human-readable description of this Application App Management Policy ID
func (id ApplicationAppManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
	}
	return fmt.Sprintf("Application App Management Policy (%s)", strings.Join(components, "\n"))
}
