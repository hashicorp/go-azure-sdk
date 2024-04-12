package guestconfigurationassignmenthcrpreports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ReportId{})
}

var _ resourceids.ResourceId = &ReportId{}

// ReportId is a struct representing the Resource ID for a Report
type ReportId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	MachineName                      string
	GuestConfigurationAssignmentName string
	ReportId                         string
}

// NewReportID returns a new ReportId struct
func NewReportID(subscriptionId string, resourceGroupName string, machineName string, guestConfigurationAssignmentName string, reportId string) ReportId {
	return ReportId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		MachineName:                      machineName,
		GuestConfigurationAssignmentName: guestConfigurationAssignmentName,
		ReportId:                         reportId,
	}
}

// ParseReportID parses 'input' into a ReportId
func ParseReportID(input string) (*ReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseReportIDInsensitively parses 'input' case-insensitively into a ReportId
// note: this method should only be used for API response data and not user input
func ParseReportIDInsensitively(input string) (*ReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ReportId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.MachineName, ok = input.Parsed["machineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "machineName", input)
	}

	if id.GuestConfigurationAssignmentName, ok = input.Parsed["guestConfigurationAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", input)
	}

	if id.ReportId, ok = input.Parsed["reportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "reportId", input)
	}

	return nil
}

// ValidateReportID checks that 'input' can be parsed as a Report ID
func ValidateReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Report ID
func (id ReportId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HybridCompute/machines/%s/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MachineName, id.GuestConfigurationAssignmentName, id.ReportId)
}

// Segments returns a slice of Resource ID Segments which comprise this Report ID
func (id ReportId) Segments() []resourceids.Segment {
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

// String returns a human-readable description of this Report ID
func (id ReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Machine Name: %q", id.MachineName),
		fmt.Sprintf("Guest Configuration Assignment Name: %q", id.GuestConfigurationAssignmentName),
		fmt.Sprintf("Report: %q", id.ReportId),
	}
	return fmt.Sprintf("Report (%s)", strings.Join(components, "\n"))
}
