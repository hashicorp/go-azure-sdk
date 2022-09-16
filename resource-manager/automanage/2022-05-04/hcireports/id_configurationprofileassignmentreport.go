package hcireports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ConfigurationProfileAssignmentReportId{}

// ConfigurationProfileAssignmentReportId is a struct representing the Resource ID for a Configuration Profile Assignment Report
type ConfigurationProfileAssignmentReportId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	ClusterName                        string
	ConfigurationProfileAssignmentName string
	ReportName                         string
}

// NewConfigurationProfileAssignmentReportID returns a new ConfigurationProfileAssignmentReportId struct
func NewConfigurationProfileAssignmentReportID(subscriptionId string, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, reportName string) ConfigurationProfileAssignmentReportId {
	return ConfigurationProfileAssignmentReportId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		ClusterName:                        clusterName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
		ReportName:                         reportName,
	}
}

// ParseConfigurationProfileAssignmentReportID parses 'input' into a ConfigurationProfileAssignmentReportId
func ParseConfigurationProfileAssignmentReportID(input string) (*ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationProfileAssignmentName' was not found in the resource id %q", input)
	}

	if id.ReportName, ok = parsed.Parsed["reportName"]; !ok {
		return nil, fmt.Errorf("the segment 'reportName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseConfigurationProfileAssignmentReportIDInsensitively parses 'input' case-insensitively into a ConfigurationProfileAssignmentReportId
// note: this method should only be used for API response data and not user input
func ParseConfigurationProfileAssignmentReportIDInsensitively(input string) (*ConfigurationProfileAssignmentReportId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileAssignmentReportId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileAssignmentReportId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationProfileAssignmentName' was not found in the resource id %q", input)
	}

	if id.ReportName, ok = parsed.Parsed["reportName"]; !ok {
		return nil, fmt.Errorf("the segment 'reportName' was not found in the resource id %q", input)
	}

	return &id, nil
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
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AzureStackHCI/clusters/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s/reports/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ConfigurationProfileAssignmentName, id.ReportName)
}

// Segments returns a slice of Resource ID Segments which comprise this Configuration Profile Assignment Report ID
func (id ConfigurationProfileAssignmentReportId) Segments() []resourceids.Segment {
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

// String returns a human-readable description of this Configuration Profile Assignment Report ID
func (id ConfigurationProfileAssignmentReportId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
		fmt.Sprintf("Report Name: %q", id.ReportName),
	}
	return fmt.Sprintf("Configuration Profile Assignment Report (%s)", strings.Join(components, "\n"))
}
