package managementconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ManagementConfigurationId{}

// ManagementConfigurationId is a struct representing the Resource ID for a Management Configuration
type ManagementConfigurationId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ManagementConfigurationName string
}

// NewManagementConfigurationID returns a new ManagementConfigurationId struct
func NewManagementConfigurationID(subscriptionId string, resourceGroupName string, managementConfigurationName string) ManagementConfigurationId {
	return ManagementConfigurationId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ManagementConfigurationName: managementConfigurationName,
	}
}

// ParseManagementConfigurationID parses 'input' into a ManagementConfigurationId
func ParseManagementConfigurationID(input string) (*ManagementConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagementConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagementConfigurationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagementConfigurationIDInsensitively parses 'input' case-insensitively into a ManagementConfigurationId
// note: this method should only be used for API response data and not user input
func ParseManagementConfigurationIDInsensitively(input string) (*ManagementConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagementConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagementConfigurationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagementConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagementConfigurationName, ok = input.Parsed["managementConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementConfigurationName", input)
	}

	return nil
}

// ValidateManagementConfigurationID checks that 'input' can be parsed as a Management Configuration ID
func ValidateManagementConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagementConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Management Configuration ID
func (id ManagementConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationsManagement/managementConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagementConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Management Configuration ID
func (id ManagementConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationsManagement", "Microsoft.OperationsManagement", "Microsoft.OperationsManagement"),
		resourceids.StaticSegment("staticManagementConfigurations", "managementConfigurations", "managementConfigurations"),
		resourceids.UserSpecifiedSegment("managementConfigurationName", "managementConfigurationValue"),
	}
}

// String returns a human-readable description of this Management Configuration ID
func (id ManagementConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Management Configuration Name: %q", id.ManagementConfigurationName),
	}
	return fmt.Sprintf("Management Configuration (%s)", strings.Join(components, "\n"))
}
