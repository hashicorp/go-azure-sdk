package virtualmachinescalesetvmruncommands

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VirtualMachineScaleSetVirtualMachineRunCommandId{})
}

var _ resourceids.ResourceId = &VirtualMachineScaleSetVirtualMachineRunCommandId{}

// VirtualMachineScaleSetVirtualMachineRunCommandId is a struct representing the Resource ID for a Virtual Machine Scale Set Virtual Machine Run Command
type VirtualMachineScaleSetVirtualMachineRunCommandId struct {
	SubscriptionId             string
	ResourceGroupName          string
	VirtualMachineScaleSetName string
	InstanceId                 string
	RunCommandName             string
}

// NewVirtualMachineScaleSetVirtualMachineRunCommandID returns a new VirtualMachineScaleSetVirtualMachineRunCommandId struct
func NewVirtualMachineScaleSetVirtualMachineRunCommandID(subscriptionId string, resourceGroupName string, virtualMachineScaleSetName string, instanceId string, runCommandName string) VirtualMachineScaleSetVirtualMachineRunCommandId {
	return VirtualMachineScaleSetVirtualMachineRunCommandId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		VirtualMachineScaleSetName: virtualMachineScaleSetName,
		InstanceId:                 instanceId,
		RunCommandName:             runCommandName,
	}
}

// ParseVirtualMachineScaleSetVirtualMachineRunCommandID parses 'input' into a VirtualMachineScaleSetVirtualMachineRunCommandId
func ParseVirtualMachineScaleSetVirtualMachineRunCommandID(input string) (*VirtualMachineScaleSetVirtualMachineRunCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineScaleSetVirtualMachineRunCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineScaleSetVirtualMachineRunCommandId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVirtualMachineScaleSetVirtualMachineRunCommandIDInsensitively parses 'input' case-insensitively into a VirtualMachineScaleSetVirtualMachineRunCommandId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScaleSetVirtualMachineRunCommandIDInsensitively(input string) (*VirtualMachineScaleSetVirtualMachineRunCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineScaleSetVirtualMachineRunCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineScaleSetVirtualMachineRunCommandId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VirtualMachineScaleSetVirtualMachineRunCommandId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VirtualMachineScaleSetName, ok = input.Parsed["virtualMachineScaleSetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineScaleSetName", input)
	}

	if id.InstanceId, ok = input.Parsed["instanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "instanceId", input)
	}

	if id.RunCommandName, ok = input.Parsed["runCommandName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runCommandName", input)
	}

	return nil
}

// ValidateVirtualMachineScaleSetVirtualMachineRunCommandID checks that 'input' can be parsed as a Virtual Machine Scale Set Virtual Machine Run Command ID
func ValidateVirtualMachineScaleSetVirtualMachineRunCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineScaleSetVirtualMachineRunCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Scale Set Virtual Machine Run Command ID
func (id VirtualMachineScaleSetVirtualMachineRunCommandId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/virtualMachines/%s/runCommands/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineScaleSetName, id.InstanceId, id.RunCommandName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Scale Set Virtual Machine Run Command ID
func (id VirtualMachineScaleSetVirtualMachineRunCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachineScaleSets", "virtualMachineScaleSets", "virtualMachineScaleSets"),
		resourceids.UserSpecifiedSegment("virtualMachineScaleSetName", "virtualMachineScaleSetValue"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("instanceId", "instanceIdValue"),
		resourceids.StaticSegment("staticRunCommands", "runCommands", "runCommands"),
		resourceids.UserSpecifiedSegment("runCommandName", "runCommandValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Scale Set Virtual Machine Run Command ID
func (id VirtualMachineScaleSetVirtualMachineRunCommandId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Scale Set Name: %q", id.VirtualMachineScaleSetName),
		fmt.Sprintf("Instance: %q", id.InstanceId),
		fmt.Sprintf("Run Command Name: %q", id.RunCommandName),
	}
	return fmt.Sprintf("Virtual Machine Scale Set Virtual Machine Run Command (%s)", strings.Join(components, "\n"))
}
