package hcrpreports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ConfigurationProfileAssignmentReportId{})
}

var _ resourceids.ResourceId = &ConfigurationProfileAssignmentReportId{}

// ConfigurationProfileAssignmentReportId is a struct representing the Resource ID for a Configuration Profile Assignment Report
type ConfigurationProfileAssignmentReportId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	MachineName                        string
	ConfigurationProfileAssignmentName string
	ReportName                         string
}

// NewConfigurationProfileAssignmentReportID returns a new ConfigurationProfileAssignmentReportId struct
func NewConfigurationProfileAssignmentReportID(subscriptionId string, resourceGroupName string, machineName string, configurationProfileAssignmentName string, reportName string) ConfigurationProfileAssignmentReportId {
	return ConfigurationProfileAssignmentReportId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		MachineName:                        machineName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
		ReportName:                         reportName,
	}
}

// ParseConfigurationProfileAssignmentReportID parses 'input' into a ConfigurationProfileAssignmentReportId
func ParseConfigurationProfileAssignmentReportID(input string) (*ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConfigurationProfileAssignmentReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConfigurationProfileAssignmentReportIDInsensitively parses 'input' case-insensitively into a ConfigurationProfileAssignmentReportId
// note: this method should only be used for API response data and not user input
func ParseConfigurationProfileAssignmentReportIDInsensitively(input string) (*ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConfigurationProfileAssignmentReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConfigurationProfileAssignmentReportId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ConfigurationProfileAssignmentName, ok = input.Parsed["configurationProfileAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", input)
	}

	if id.ReportName, ok = input.Parsed["reportName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "reportName", input)
	}

	return nil
}

// ValidateConfigurationProfileAssignmentReportID checks that 'input' can be parsed as a Configuration Profile Assignment Report ID
func ValidateConfigurationProfileAssignmentReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConfigurationProfileAssignmentReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Configuration Profile Assignment Report ID
func (id ConfigurationProfileAssignmentReportId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HybridCompute/machines/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MachineName, id.ConfigurationProfileAssignmentName, id.ReportName)
}

// Segments returns a slice of Resource ID Segments which comprise this Configuration Profile Assignment Report ID
func (id ConfigurationProfileAssignmentReportId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReports", "reports", "reports"),
		resourceids.UserSpecifiedSegment("reportName", "reportValue"),
	}
}

// String returns a human-readable description of this Configuration Profile Assignment Report ID
func (id ConfigurationProfileAssignmentReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Machine Name: %q", id.MachineName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
		fmt.Sprintf("Report Name: %q", id.ReportName),
	}
	return fmt.Sprintf("Configuration Profile Assignment Report (%s)", strings.Join(components, "\n"))
}
