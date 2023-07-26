package machinepools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MachinePoolId{}

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
	parser := resourceids.NewParserFromResourceIdType(MachinePoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MachinePoolId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.OpenShiftClusterName, ok = parsed.Parsed["openShiftClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftClusterName", *parsed)
	}

	if id.MachinePoolName, ok = parsed.Parsed["machinePoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "machinePoolName", *parsed)
	}

	return &id, nil
}

// ParseMachinePoolIDInsensitively parses 'input' case-insensitively into a MachinePoolId
// note: this method should only be used for API response data and not user input
func ParseMachinePoolIDInsensitively(input string) (*MachinePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(MachinePoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MachinePoolId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.OpenShiftClusterName, ok = parsed.Parsed["openShiftClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftClusterName", *parsed)
	}

	if id.MachinePoolName, ok = parsed.Parsed["machinePoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "machinePoolName", *parsed)
	}

	return &id, nil
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
		resourceids.UserSpecifiedSegment("openShiftClusterName", "openShiftClusterValue"),
		resourceids.StaticSegment("staticMachinePool", "machinePool", "machinePool"),
		resourceids.UserSpecifiedSegment("machinePoolName", "machinePoolValue"),
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
