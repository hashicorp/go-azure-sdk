package workloadclassifiers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkloadClassifierId{})
}

var _ resourceids.ResourceId = &WorkloadClassifierId{}

// WorkloadClassifierId is a struct representing the Resource ID for a Workload Classifier
type WorkloadClassifierId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ServerName             string
	DatabaseName           string
	WorkloadGroupName      string
	WorkloadClassifierName string
}

// NewWorkloadClassifierID returns a new WorkloadClassifierId struct
func NewWorkloadClassifierID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string) WorkloadClassifierId {
	return WorkloadClassifierId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ServerName:             serverName,
		DatabaseName:           databaseName,
		WorkloadGroupName:      workloadGroupName,
		WorkloadClassifierName: workloadClassifierName,
	}
}

// ParseWorkloadClassifierID parses 'input' into a WorkloadClassifierId
func ParseWorkloadClassifierID(input string) (*WorkloadClassifierId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkloadClassifierId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkloadClassifierId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkloadClassifierIDInsensitively parses 'input' case-insensitively into a WorkloadClassifierId
// note: this method should only be used for API response data and not user input
func ParseWorkloadClassifierIDInsensitively(input string) (*WorkloadClassifierId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkloadClassifierId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkloadClassifierId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkloadClassifierId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServerName, ok = input.Parsed["serverName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverName", input)
	}

	if id.DatabaseName, ok = input.Parsed["databaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseName", input)
	}

	if id.WorkloadGroupName, ok = input.Parsed["workloadGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workloadGroupName", input)
	}

	if id.WorkloadClassifierName, ok = input.Parsed["workloadClassifierName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workloadClassifierName", input)
	}

	return nil
}

// ValidateWorkloadClassifierID checks that 'input' can be parsed as a Workload Classifier ID
func ValidateWorkloadClassifierID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkloadClassifierID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workload Classifier ID
func (id WorkloadClassifierId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/workloadGroups/%s/workloadClassifiers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.WorkloadGroupName, id.WorkloadClassifierName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workload Classifier ID
func (id WorkloadClassifierId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticWorkloadClassifiers", "workloadClassifiers", "workloadClassifiers"),
		resourceids.UserSpecifiedSegment("workloadClassifierName", "workloadClassifierValue"),
	}
}

// String returns a human-readable description of this Workload Classifier ID
func (id WorkloadClassifierId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Workload Group Name: %q", id.WorkloadGroupName),
		fmt.Sprintf("Workload Classifier Name: %q", id.WorkloadClassifierName),
	}
	return fmt.Sprintf("Workload Classifier (%s)", strings.Join(components, "\n"))
}
