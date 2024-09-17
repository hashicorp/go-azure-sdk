package reports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&Providers2ConfigurationProfileAssignmentReportId{})
}

var _ resourceids.ResourceId = &Providers2ConfigurationProfileAssignmentReportId{}

// Providers2ConfigurationProfileAssignmentReportId is a struct representing the Resource ID for a Providers 2 Configuration Profile Assignment Report
type Providers2ConfigurationProfileAssignmentReportId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VirtualMachineName                 string
	ConfigurationProfileAssignmentName string
	ReportName                         string
}

// NewProviders2ConfigurationProfileAssignmentReportID returns a new Providers2ConfigurationProfileAssignmentReportId struct
func NewProviders2ConfigurationProfileAssignmentReportID(subscriptionId string, resourceGroupName string, virtualMachineName string, configurationProfileAssignmentName string, reportName string) Providers2ConfigurationProfileAssignmentReportId {
	return Providers2ConfigurationProfileAssignmentReportId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VirtualMachineName:                 virtualMachineName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
		ReportName:                         reportName,
	}
}

// ParseProviders2ConfigurationProfileAssignmentReportID parses 'input' into a Providers2ConfigurationProfileAssignmentReportId
func ParseProviders2ConfigurationProfileAssignmentReportID(input string) (*Providers2ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2ConfigurationProfileAssignmentReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviders2ConfigurationProfileAssignmentReportIDInsensitively parses 'input' case-insensitively into a Providers2ConfigurationProfileAssignmentReportId
// note: this method should only be used for API response data and not user input
func ParseProviders2ConfigurationProfileAssignmentReportIDInsensitively(input string) (*Providers2ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2ConfigurationProfileAssignmentReportId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *Providers2ConfigurationProfileAssignmentReportId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VirtualMachineName, ok = input.Parsed["virtualMachineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", input)
	}

	if id.ConfigurationProfileAssignmentName, ok = input.Parsed["configurationProfileAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", input)
	}

	if id.ReportName, ok = input.Parsed["reportName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "reportName", input)
	}

	return nil
}

// ValidateProviders2ConfigurationProfileAssignmentReportID checks that 'input' can be parsed as a Providers 2 Configuration Profile Assignment Report ID
func ValidateProviders2ConfigurationProfileAssignmentReportID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2ConfigurationProfileAssignmentReportID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Configuration Profile Assignment Report ID
func (id Providers2ConfigurationProfileAssignmentReportId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.ConfigurationProfileAssignmentName, id.ReportName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Configuration Profile Assignment Report ID
func (id Providers2ConfigurationProfileAssignmentReportId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReports", "reports", "reports"),
		resourceids.UserSpecifiedSegment("reportName", "reportValue"),
	}
}

// String returns a human-readable description of this Providers 2 Configuration Profile Assignment Report ID
func (id Providers2ConfigurationProfileAssignmentReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
		fmt.Sprintf("Report Name: %q", id.ReportName),
	}
	return fmt.Sprintf("Providers 2 Configuration Profile Assignment Report (%s)", strings.Join(components, "\n"))
}
