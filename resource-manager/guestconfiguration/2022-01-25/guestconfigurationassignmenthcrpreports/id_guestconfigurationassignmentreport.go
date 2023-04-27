package guestconfigurationassignmenthcrpreports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GuestConfigurationAssignmentReportId{}

// GuestConfigurationAssignmentReportId is a struct representing the Resource ID for a Guest Configuration Assignment Report
type GuestConfigurationAssignmentReportId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	MachineName                      string
	GuestConfigurationAssignmentName string
	ReportId                         string
}

// NewGuestConfigurationAssignmentReportID returns a new GuestConfigurationAssignmentReportId struct
func NewGuestConfigurationAssignmentReportID(subscriptionId string, resourceGroupName string, machineName string, guestConfigurationAssignmentName string, reportId string) GuestConfigurationAssignmentReportId {
	return GuestConfigurationAssignmentReportId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		MachineName:                      machineName,
		GuestConfigurationAssignmentName: guestConfigurationAssignmentName,
		ReportId:                         reportId,
	}
}

// ParseGuestConfigurationAssignmentReportID parses 'input' into a GuestConfigurationAssignmentReportId
func ParseGuestConfigurationAssignmentReportID(input string) (*GuestConfigurationAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(GuestConfigurationAssignmentReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GuestConfigurationAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.MachineName, ok = parsed.Parsed["machineName"]; !ok {
		return nil, fmt.Errorf("the segment 'machineName' was not found in the resource id %q", input)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'guestConfigurationAssignmentName' was not found in the resource id %q", input)
	}

	if id.ReportId, ok = parsed.Parsed["reportId"]; !ok {
		return nil, fmt.Errorf("the segment 'reportId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseGuestConfigurationAssignmentReportIDInsensitively parses 'input' case-insensitively into a GuestConfigurationAssignmentReportId
// note: this method should only be used for API response data and not user input
func ParseGuestConfigurationAssignmentReportIDInsensitively(input string) (*GuestConfigurationAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(GuestConfigurationAssignmentReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GuestConfigurationAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.MachineName, ok = parsed.Parsed["machineName"]; !ok {
		return nil, fmt.Errorf("the segment 'machineName' was not found in the resource id %q", input)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'guestConfigurationAssignmentName' was not found in the resource id %q", input)
	}

	if id.ReportId, ok = parsed.Parsed["reportId"]; !ok {
		return nil, fmt.Errorf("the segment 'reportId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateGuestConfigurationAssignmentReportID checks that 'input' can be parsed as a Guest Configuration Assignment Report ID
func ValidateGuestConfigurationAssignmentReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGuestConfigurationAssignmentReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Guest Configuration Assignment Report ID
func (id GuestConfigurationAssignmentReportId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HybridCompute/machines/%s/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MachineName, id.GuestConfigurationAssignmentName, id.ReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Guest Configuration Assignment Report ID
func (id GuestConfigurationAssignmentReportId) Segments() []resourceids.Segment {
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
		resourceids.ResourceProviderSegment("staticMicrosoftGuestConfiguration", "Microsoft.GuestConfiguration", "Microsoft.GuestConfiguration"),
		resourceids.StaticSegment("staticGuestConfigurationAssignments", "guestConfigurationAssignments", "guestConfigurationAssignments"),
		resourceids.UserSpecifiedSegment("guestConfigurationAssignmentName", "guestConfigurationAssignmentValue"),
		resourceids.StaticSegment("staticReports", "reports", "reports"),
		resourceids.UserSpecifiedSegment("reportId", "reportIdValue"),
	}
}

// String returns a human-readable description of this Guest Configuration Assignment Report ID
func (id GuestConfigurationAssignmentReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Machine Name: %q", id.MachineName),
		fmt.Sprintf("Guest Configuration Assignment Name: %q", id.GuestConfigurationAssignmentName),
		fmt.Sprintf("Report: %q", id.ReportId),
	}
	return fmt.Sprintf("Guest Configuration Assignment Report (%s)", strings.Join(components, "\n"))
}
