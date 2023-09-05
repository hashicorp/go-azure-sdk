package configurationprofileassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VirtualMachineProviders2ConfigurationProfileAssignmentId{}

// VirtualMachineProviders2ConfigurationProfileAssignmentId is a struct representing the Resource ID for a Virtual Machine Providers 2 Configuration Profile Assignment
type VirtualMachineProviders2ConfigurationProfileAssignmentId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VirtualMachineName                 string
	ConfigurationProfileAssignmentName string
}

// NewVirtualMachineProviders2ConfigurationProfileAssignmentID returns a new VirtualMachineProviders2ConfigurationProfileAssignmentId struct
func NewVirtualMachineProviders2ConfigurationProfileAssignmentID(subscriptionId string, resourceGroupName string, virtualMachineName string, configurationProfileAssignmentName string) VirtualMachineProviders2ConfigurationProfileAssignmentId {
	return VirtualMachineProviders2ConfigurationProfileAssignmentId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VirtualMachineName:                 virtualMachineName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
	}
}

// ParseVirtualMachineProviders2ConfigurationProfileAssignmentID parses 'input' into a VirtualMachineProviders2ConfigurationProfileAssignmentId
func ParseVirtualMachineProviders2ConfigurationProfileAssignmentID(input string) (*VirtualMachineProviders2ConfigurationProfileAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineProviders2ConfigurationProfileAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineProviders2ConfigurationProfileAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseVirtualMachineProviders2ConfigurationProfileAssignmentIDInsensitively parses 'input' case-insensitively into a VirtualMachineProviders2ConfigurationProfileAssignmentId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineProviders2ConfigurationProfileAssignmentIDInsensitively(input string) (*VirtualMachineProviders2ConfigurationProfileAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineProviders2ConfigurationProfileAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineProviders2ConfigurationProfileAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateVirtualMachineProviders2ConfigurationProfileAssignmentID checks that 'input' can be parsed as a Virtual Machine Providers 2 Configuration Profile Assignment ID
func ValidateVirtualMachineProviders2ConfigurationProfileAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineProviders2ConfigurationProfileAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Providers 2 Configuration Profile Assignment ID
func (id VirtualMachineProviders2ConfigurationProfileAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.ConfigurationProfileAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Providers 2 Configuration Profile Assignment ID
func (id VirtualMachineProviders2ConfigurationProfileAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticConfigurationProfileAssignments", "configurationProfileAssignments", "configurationProfileAssignments"),
		resourceids.UserSpecifiedSegment("configurationProfileAssignmentName", "configurationProfileAssignmentValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Providers 2 Configuration Profile Assignment ID
func (id VirtualMachineProviders2ConfigurationProfileAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
	}
	return fmt.Sprintf("Virtual Machine Providers 2 Configuration Profile Assignment (%s)", strings.Join(components, "\n"))
}
