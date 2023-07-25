package appserviceplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkerId{}

// WorkerId is a struct representing the Resource ID for a Worker
type WorkerId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerFarmName    string
	WorkerName        string
}

// NewWorkerID returns a new WorkerId struct
func NewWorkerID(subscriptionId string, resourceGroupName string, serverFarmName string, workerName string) WorkerId {
	return WorkerId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerFarmName:    serverFarmName,
		WorkerName:        workerName,
	}
}

// ParseWorkerID parses 'input' into a WorkerId
func ParseWorkerID(input string) (*WorkerId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	if id.WorkerName, ok = parsed.Parsed["workerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workerName", *parsed)
	}

	return &id, nil
}

// ParseWorkerIDInsensitively parses 'input' case-insensitively into a WorkerId
// note: this method should only be used for API response data and not user input
func ParseWorkerIDInsensitively(input string) (*WorkerId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	if id.WorkerName, ok = parsed.Parsed["workerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workerName", *parsed)
	}

	return &id, nil
}

// ValidateWorkerID checks that 'input' can be parsed as a Worker ID
func ValidateWorkerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Worker ID
func (id WorkerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/serverFarms/%s/workers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerFarmName, id.WorkerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Worker ID
func (id WorkerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticServerFarms", "serverFarms", "serverFarms"),
		resourceids.UserSpecifiedSegment("serverFarmName", "serverFarmValue"),
		resourceids.StaticSegment("staticWorkers", "workers", "workers"),
		resourceids.UserSpecifiedSegment("workerName", "workerValue"),
	}
}

// String returns a human-readable description of this Worker ID
func (id WorkerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Farm Name: %q", id.ServerFarmName),
		fmt.Sprintf("Worker Name: %q", id.WorkerName),
	}
	return fmt.Sprintf("Worker (%s)", strings.Join(components, "\n"))
}
