package workloadclassifiers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkloadGroupId{}

// WorkloadGroupId is a struct representing the Resource ID for a Workload Group
type WorkloadGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	DatabaseName      string
	WorkloadGroupName string
}

// NewWorkloadGroupID returns a new WorkloadGroupId struct
func NewWorkloadGroupID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, workloadGroupName string) WorkloadGroupId {
	return WorkloadGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		DatabaseName:      databaseName,
		WorkloadGroupName: workloadGroupName,
	}
}

// ParseWorkloadGroupID parses 'input' into a WorkloadGroupId
func ParseWorkloadGroupID(input string) (*WorkloadGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkloadGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkloadGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.WorkloadGroupName, ok = parsed.Parsed["workloadGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workloadGroupName", *parsed)
	}

	return &id, nil
}

// ParseWorkloadGroupIDInsensitively parses 'input' case-insensitively into a WorkloadGroupId
// note: this method should only be used for API response data and not user input
func ParseWorkloadGroupIDInsensitively(input string) (*WorkloadGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkloadGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkloadGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.WorkloadGroupName, ok = parsed.Parsed["workloadGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workloadGroupName", *parsed)
	}

	return &id, nil
}

// ValidateWorkloadGroupID checks that 'input' can be parsed as a Workload Group ID
func ValidateWorkloadGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkloadGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workload Group ID
func (id WorkloadGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/workloadGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.WorkloadGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workload Group ID
func (id WorkloadGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticWorkloadGroups", "workloadGroups", "workloadGroups"),
		resourceids.UserSpecifiedSegment("workloadGroupName", "workloadGroupValue"),
	}
}

// String returns a human-readable description of this Workload Group ID
func (id WorkloadGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Workload Group Name: %q", id.WorkloadGroupName),
	}
	return fmt.Sprintf("Workload Group (%s)", strings.Join(components, "\n"))
}
