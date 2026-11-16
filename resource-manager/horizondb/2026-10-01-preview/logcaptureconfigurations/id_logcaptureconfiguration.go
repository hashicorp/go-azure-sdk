package logcaptureconfigurations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LogCaptureConfigurationId{})
}

var _ resourceids.ResourceId = &LogCaptureConfigurationId{}

// LogCaptureConfigurationId is a struct representing the Resource ID for a Log Capture Configuration
type LogCaptureConfigurationId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ClusterName                 string
	LogCaptureConfigurationName string
}

// NewLogCaptureConfigurationID returns a new LogCaptureConfigurationId struct
func NewLogCaptureConfigurationID(subscriptionId string, resourceGroupName string, clusterName string, logCaptureConfigurationName string) LogCaptureConfigurationId {
	return LogCaptureConfigurationId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ClusterName:                 clusterName,
		LogCaptureConfigurationName: logCaptureConfigurationName,
	}
}

// ParseLogCaptureConfigurationID parses 'input' into a LogCaptureConfigurationId
func ParseLogCaptureConfigurationID(input string) (*LogCaptureConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LogCaptureConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LogCaptureConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLogCaptureConfigurationIDInsensitively parses 'input' case-insensitively into a LogCaptureConfigurationId
// note: this method should only be used for API response data and not user input
func ParseLogCaptureConfigurationIDInsensitively(input string) (*LogCaptureConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LogCaptureConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LogCaptureConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LogCaptureConfigurationId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LogCaptureConfigurationName, ok = input.Parsed["logCaptureConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "logCaptureConfigurationName", input)
	}

	return nil
}

// ValidateLogCaptureConfigurationID checks that 'input' can be parsed as a Log Capture Configuration ID
func ValidateLogCaptureConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLogCaptureConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Log Capture Configuration ID
func (id LogCaptureConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/clusters/%s/logCaptureConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.LogCaptureConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Log Capture Configuration ID
func (id LogCaptureConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHorizonDb", "Microsoft.HorizonDb", "Microsoft.HorizonDb"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticLogCaptureConfigurations", "logCaptureConfigurations", "logCaptureConfigurations"),
		resourceids.UserSpecifiedSegment("logCaptureConfigurationName", "logCaptureConfigurationName"),
	}
}

// String returns a human-readable description of this Log Capture Configuration ID
func (id LogCaptureConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Log Capture Configuration Name: %q", id.LogCaptureConfigurationName),
	}
	return fmt.Sprintf("Log Capture Configuration (%s)", strings.Join(components, "\n"))
}
