package devops

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SecurityConnectorId{}

// SecurityConnectorId is a struct representing the Resource ID for a Security Connector
type SecurityConnectorId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SecurityConnectorName string
}

// NewSecurityConnectorID returns a new SecurityConnectorId struct
func NewSecurityConnectorID(subscriptionId string, resourceGroupName string, securityConnectorName string) SecurityConnectorId {
	return SecurityConnectorId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SecurityConnectorName: securityConnectorName,
	}
}

// ParseSecurityConnectorID parses 'input' into a SecurityConnectorId
func ParseSecurityConnectorID(input string) (*SecurityConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecurityConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecurityConnectorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSecurityConnectorIDInsensitively parses 'input' case-insensitively into a SecurityConnectorId
// note: this method should only be used for API response data and not user input
func ParseSecurityConnectorIDInsensitively(input string) (*SecurityConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecurityConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecurityConnectorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SecurityConnectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SecurityConnectorName, ok = input.Parsed["securityConnectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityConnectorName", input)
	}

	return nil
}

// ValidateSecurityConnectorID checks that 'input' can be parsed as a Security Connector ID
func ValidateSecurityConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSecurityConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Security Connector ID
func (id SecurityConnectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/securityConnectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SecurityConnectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Security Connector ID
func (id SecurityConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecurityConnectors", "securityConnectors", "securityConnectors"),
		resourceids.UserSpecifiedSegment("securityConnectorName", "securityConnectorValue"),
	}
}

// String returns a human-readable description of this Security Connector ID
func (id SecurityConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Security Connector Name: %q", id.SecurityConnectorName),
	}
	return fmt.Sprintf("Security Connector (%s)", strings.Join(components, "\n"))
}
