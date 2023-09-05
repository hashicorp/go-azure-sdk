package hcrpreports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2ConfigurationProfileAssignmentId{}

// Providers2ConfigurationProfileAssignmentId is a struct representing the Resource ID for a Providers 2 Configuration Profile Assignment
type Providers2ConfigurationProfileAssignmentId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	MachineName                        string
	ConfigurationProfileAssignmentName string
}

// NewProviders2ConfigurationProfileAssignmentID returns a new Providers2ConfigurationProfileAssignmentId struct
func NewProviders2ConfigurationProfileAssignmentID(subscriptionId string, resourceGroupName string, machineName string, configurationProfileAssignmentName string) Providers2ConfigurationProfileAssignmentId {
	return Providers2ConfigurationProfileAssignmentId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		MachineName:                        machineName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
	}
}

// ParseProviders2ConfigurationProfileAssignmentID parses 'input' into a Providers2ConfigurationProfileAssignmentId
func ParseProviders2ConfigurationProfileAssignmentID(input string) (*Providers2ConfigurationProfileAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ConfigurationProfileAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ConfigurationProfileAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.MachineName, ok = parsed.Parsed["machineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "machineName", *parsed)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseProviders2ConfigurationProfileAssignmentIDInsensitively parses 'input' case-insensitively into a Providers2ConfigurationProfileAssignmentId
// note: this method should only be used for API response data and not user input
func ParseProviders2ConfigurationProfileAssignmentIDInsensitively(input string) (*Providers2ConfigurationProfileAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ConfigurationProfileAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ConfigurationProfileAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.MachineName, ok = parsed.Parsed["machineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "machineName", *parsed)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateProviders2ConfigurationProfileAssignmentID checks that 'input' can be parsed as a Providers 2 Configuration Profile Assignment ID
func ValidateProviders2ConfigurationProfileAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2ConfigurationProfileAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Configuration Profile Assignment ID
func (id Providers2ConfigurationProfileAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HybridCompute/machines/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MachineName, id.ConfigurationProfileAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Configuration Profile Assignment ID
func (id Providers2ConfigurationProfileAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHybridCompute", "Microsoft.HybridCompute", "Microsoft.HybridCompute"),
		resourceids.StaticSegment("staticMachines", "machines", "machines"),
		resourceids.UserSpecifiedSegment("machineName", "machineValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticConfigurationProfileAssignments", "configurationProfileAssignments", "configurationProfileAssignments"),
		resourceids.UserSpecifiedSegment("configurationProfileAssignmentName", "configurationProfileAssignmentValue"),
	}
}

// String returns a human-readable description of this Providers 2 Configuration Profile Assignment ID
func (id Providers2ConfigurationProfileAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Machine Name: %q", id.MachineName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
	}
	return fmt.Sprintf("Providers 2 Configuration Profile Assignment (%s)", strings.Join(components, "\n"))
}
