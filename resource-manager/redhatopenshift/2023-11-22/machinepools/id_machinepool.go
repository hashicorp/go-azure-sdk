package machinepools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MachinePoolId{})
}

var _ resourceids.ResourceId = &MachinePoolId{}

// MachinePoolId is a struct representing the Resource ID for a Machine Pool
type MachinePoolId struct {
	SubscriptionId       string
	ResourceGroupName    string
	OpenShiftClusterName string
	MachinePoolName      string
}

// NewMachinePoolID returns a new MachinePoolId struct
func NewMachinePoolID(subscriptionId string, resourceGroupName string, openShiftClusterName string, machinePoolName string) MachinePoolId {
	return MachinePoolId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		OpenShiftClusterName: openShiftClusterName,
		MachinePoolName:      machinePoolName,
	}
}

// ParseMachinePoolID parses 'input' into a MachinePoolId
func ParseMachinePoolID(input string) (*MachinePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MachinePoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MachinePoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMachinePoolIDInsensitively parses 'input' case-insensitively into a MachinePoolId
// note: this method should only be used for API response data and not user input
func ParseMachinePoolIDInsensitively(input string) (*MachinePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MachinePoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MachinePoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MachinePoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.OpenShiftClusterName, ok = input.Parsed["openShiftClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftClusterName", input)
	}

	if id.MachinePoolName, ok = input.Parsed["machinePoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "machinePoolName", input)
	}

	return nil
}

// ValidateMachinePoolID checks that 'input' can be parsed as a Machine Pool ID
func ValidateMachinePoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMachinePoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Machine Pool ID
func (id MachinePoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RedHatOpenShift/openShiftClusters/%s/machinePool/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.OpenShiftClusterName, id.MachinePoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Machine Pool ID
func (id MachinePoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticOpenShiftClusters", "openShiftClusters", "openShiftClusters"),
		resourceids.UserSpecifiedSegment("openShiftClusterName", "openShiftClusterName"),
		resourceids.StaticSegment("staticMachinePool", "machinePool", "machinePool"),
		resourceids.UserSpecifiedSegment("machinePoolName", "machinePoolName"),
	}
}

// String returns a human-readable description of this Machine Pool ID
func (id MachinePoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Open Shift Cluster Name: %q", id.OpenShiftClusterName),
		fmt.Sprintf("Machine Pool Name: %q", id.MachinePoolName),
	}
	return fmt.Sprintf("Machine Pool (%s)", strings.Join(components, "\n"))
}
