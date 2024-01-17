package hcireports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ReportId{}

// ReportId is a struct representing the Resource ID for a Report
type ReportId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	ClusterName                        string
	ConfigurationProfileAssignmentName string
	ReportName                         string
}

// NewReportID returns a new ReportId struct
func NewReportID(subscriptionId string, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, reportName string) ReportId {
	return ReportId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		ClusterName:                        clusterName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
		ReportName:                         reportName,
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

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.ConfigurationProfileAssignmentName, ok = input.Parsed["configurationProfileAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", input)
	}

	if id.ReportName, ok = input.Parsed["reportName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "reportName", input)
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AzureStackHCI/clusters/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ConfigurationProfileAssignmentName, id.ReportName)
}

// Segments returns a slice of Resource ID Segments which comprise this Report ID
func (id ReportId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAzureStackHCI", "Microsoft.AzureStackHCI", "Microsoft.AzureStackHCI"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticConfigurationProfileAssignments", "configurationProfileAssignments", "configurationProfileAssignments"),
		resourceids.UserSpecifiedSegment("configurationProfileAssignmentName", "configurationProfileAssignmentValue"),
		resourceids.StaticSegment("staticReports", "reports", "reports"),
		resourceids.UserSpecifiedSegment("reportName", "reportValue"),
	}
}

// String returns a human-readable description of this Report ID
func (id ReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
		fmt.Sprintf("Report Name: %q", id.ReportName),
	}
	return fmt.Sprintf("Report (%s)", strings.Join(components, "\n"))
}
