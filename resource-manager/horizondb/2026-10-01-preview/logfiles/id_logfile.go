package logfiles

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LogFileId{})
}

var _ resourceids.ResourceId = &LogFileId{}

// LogFileId is a struct representing the Resource ID for a Log File
type LogFileId struct {
	SubscriptionId    string
	ResourceGroupName string
	ClusterName       string
	LogFileName       string
}

// NewLogFileID returns a new LogFileId struct
func NewLogFileID(subscriptionId string, resourceGroupName string, clusterName string, logFileName string) LogFileId {
	return LogFileId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ClusterName:       clusterName,
		LogFileName:       logFileName,
	}
}

// ParseLogFileID parses 'input' into a LogFileId
func ParseLogFileID(input string) (*LogFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LogFileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LogFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLogFileIDInsensitively parses 'input' case-insensitively into a LogFileId
// note: this method should only be used for API response data and not user input
func ParseLogFileIDInsensitively(input string) (*LogFileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LogFileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LogFileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LogFileId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LogFileName, ok = input.Parsed["logFileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "logFileName", input)
	}

	return nil
}

// ValidateLogFileID checks that 'input' can be parsed as a Log File ID
func ValidateLogFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLogFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Log File ID
func (id LogFileId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/clusters/%s/logFiles/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.LogFileName)
}

// Segments returns a slice of Resource ID Segments which comprise this Log File ID
func (id LogFileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHorizonDb", "Microsoft.HorizonDb", "Microsoft.HorizonDb"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticLogFiles", "logFiles", "logFiles"),
		resourceids.UserSpecifiedSegment("logFileName", "logFileName"),
	}
}

// String returns a human-readable description of this Log File ID
func (id LogFileId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Log File Name: %q", id.LogFileName),
	}
	return fmt.Sprintf("Log File (%s)", strings.Join(components, "\n"))
}
