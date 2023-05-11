package guestconfigurationconnectedvmwarevsphereassignmentsreports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VirtualMachineProviders2GuestConfigurationAssignmentReportId{}

// VirtualMachineProviders2GuestConfigurationAssignmentReportId is a struct representing the Resource ID for a Virtual Machine Providers 2 Guest Configuration Assignment Report
type VirtualMachineProviders2GuestConfigurationAssignmentReportId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	VirtualMachineName               string
	GuestConfigurationAssignmentName string
	ReportId                         string
}

// NewVirtualMachineProviders2GuestConfigurationAssignmentReportID returns a new VirtualMachineProviders2GuestConfigurationAssignmentReportId struct
func NewVirtualMachineProviders2GuestConfigurationAssignmentReportID(subscriptionId string, resourceGroupName string, virtualMachineName string, guestConfigurationAssignmentName string, reportId string) VirtualMachineProviders2GuestConfigurationAssignmentReportId {
	return VirtualMachineProviders2GuestConfigurationAssignmentReportId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		VirtualMachineName:               virtualMachineName,
		GuestConfigurationAssignmentName: guestConfigurationAssignmentName,
		ReportId:                         reportId,
	}
}

// ParseVirtualMachineProviders2GuestConfigurationAssignmentReportID parses 'input' into a VirtualMachineProviders2GuestConfigurationAssignmentReportId
func ParseVirtualMachineProviders2GuestConfigurationAssignmentReportID(input string) (*VirtualMachineProviders2GuestConfigurationAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineProviders2GuestConfigurationAssignmentReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineProviders2GuestConfigurationAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", *parsed)
	}

	if id.ReportId, ok = parsed.Parsed["reportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "reportId", *parsed)
	}

	return &id, nil
}

// ParseVirtualMachineProviders2GuestConfigurationAssignmentReportIDInsensitively parses 'input' case-insensitively into a VirtualMachineProviders2GuestConfigurationAssignmentReportId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineProviders2GuestConfigurationAssignmentReportIDInsensitively(input string) (*VirtualMachineProviders2GuestConfigurationAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineProviders2GuestConfigurationAssignmentReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineProviders2GuestConfigurationAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", *parsed)
	}

	if id.ReportId, ok = parsed.Parsed["reportId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "reportId", *parsed)
	}

	return &id, nil
}

// ValidateVirtualMachineProviders2GuestConfigurationAssignmentReportID checks that 'input' can be parsed as a Virtual Machine Providers 2 Guest Configuration Assignment Report ID
func ValidateVirtualMachineProviders2GuestConfigurationAssignmentReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineProviders2GuestConfigurationAssignmentReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Providers 2 Guest Configuration Assignment Report ID
func (id VirtualMachineProviders2GuestConfigurationAssignmentReportId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/%s/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.GuestConfigurationAssignmentName, id.ReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Providers 2 Guest Configuration Assignment Report ID
func (id VirtualMachineProviders2GuestConfigurationAssignmentReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftGuestConfiguration", "Microsoft.GuestConfiguration", "Microsoft.GuestConfiguration"),
		resourceids.StaticSegment("staticGuestConfigurationAssignments", "guestConfigurationAssignments", "guestConfigurationAssignments"),
		resourceids.UserSpecifiedSegment("guestConfigurationAssignmentName", "guestConfigurationAssignmentValue"),
		resourceids.StaticSegment("staticReports", "reports", "reports"),
		resourceids.UserSpecifiedSegment("reportId", "reportIdValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Providers 2 Guest Configuration Assignment Report ID
func (id VirtualMachineProviders2GuestConfigurationAssignmentReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Guest Configuration Assignment Name: %q", id.GuestConfigurationAssignmentName),
		fmt.Sprintf("Report: %q", id.ReportId),
	}
	return fmt.Sprintf("Virtual Machine Providers 2 Guest Configuration Assignment Report (%s)", strings.Join(components, "\n"))
}
