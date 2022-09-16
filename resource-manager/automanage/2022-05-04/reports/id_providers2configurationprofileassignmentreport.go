package reports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = Providers2ConfigurationProfileAssignmentReportId{}

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
	parser := resourceids.NewParserFromResourceIdType(Providers2ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ConfigurationProfileAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineName' was not found in the resource id %q", input)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationProfileAssignmentName' was not found in the resource id %q", input)
	}

	if id.ReportName, ok = parsed.Parsed["reportName"]; !ok {
		return nil, fmt.Errorf("the segment 'reportName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviders2ConfigurationProfileAssignmentReportIDInsensitively parses 'input' case-insensitively into a Providers2ConfigurationProfileAssignmentReportId
// note: this method should only be used for API response data and not user input
func ParseProviders2ConfigurationProfileAssignmentReportIDInsensitively(input string) (*Providers2ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ConfigurationProfileAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineName' was not found in the resource id %q", input)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationProfileAssignmentName' was not found in the resource id %q", input)
	}

	if id.ReportName, ok = parsed.Parsed["reportName"]; !ok {
		return nil, fmt.Errorf("the segment 'reportName' was not found in the resource id %q", input)
	}

	return &id, nil
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
