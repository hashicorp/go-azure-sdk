package securitysolutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SecuritySolutionId{})
}

var _ resourceids.ResourceId = &SecuritySolutionId{}

// SecuritySolutionId is a struct representing the Resource ID for a Security Solution
type SecuritySolutionId struct {
	SubscriptionId       string
	ResourceGroupName    string
	LocationName         string
	SecuritySolutionName string
}

// NewSecuritySolutionID returns a new SecuritySolutionId struct
func NewSecuritySolutionID(subscriptionId string, resourceGroupName string, locationName string, securitySolutionName string) SecuritySolutionId {
	return SecuritySolutionId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		LocationName:         locationName,
		SecuritySolutionName: securitySolutionName,
	}
}

// ParseSecuritySolutionID parses 'input' into a SecuritySolutionId
func ParseSecuritySolutionID(input string) (*SecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecuritySolutionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecuritySolutionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSecuritySolutionIDInsensitively parses 'input' case-insensitively into a SecuritySolutionId
// note: this method should only be used for API response data and not user input
func ParseSecuritySolutionIDInsensitively(input string) (*SecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecuritySolutionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecuritySolutionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SecuritySolutionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.SecuritySolutionName, ok = input.Parsed["securitySolutionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securitySolutionName", input)
	}

	return nil
}

// ValidateSecuritySolutionID checks that 'input' can be parsed as a Security Solution ID
func ValidateSecuritySolutionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSecuritySolutionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Security Solution ID
func (id SecuritySolutionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/securitySolutions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.SecuritySolutionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Security Solution ID
func (id SecuritySolutionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticSecuritySolutions", "securitySolutions", "securitySolutions"),
		resourceids.UserSpecifiedSegment("securitySolutionName", "securitySolutionValue"),
	}
}

// String returns a human-readable description of this Security Solution ID
func (id SecuritySolutionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Security Solution Name: %q", id.SecuritySolutionName),
	}
	return fmt.Sprintf("Security Solution (%s)", strings.Join(components, "\n"))
}
