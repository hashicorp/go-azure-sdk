package operationstatus

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkflowOperationId{}

// WorkflowOperationId is a struct representing the Resource ID for a Workflow Operation
type WorkflowOperationId struct {
	SubscriptionId    string
	ResourceGroupName string
	LocationName      string
	WorkflowId        string
	OperationId       string
}

// NewWorkflowOperationID returns a new WorkflowOperationId struct
func NewWorkflowOperationID(subscriptionId string, resourceGroupName string, locationName string, workflowId string, operationId string) WorkflowOperationId {
	return WorkflowOperationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LocationName:      locationName,
		WorkflowId:        workflowId,
		OperationId:       operationId,
	}
}

// ParseWorkflowOperationID parses 'input' into a WorkflowOperationId
func ParseWorkflowOperationID(input string) (*WorkflowOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkflowOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkflowOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.WorkflowId, ok = parsed.Parsed["workflowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workflowId", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseWorkflowOperationIDInsensitively parses 'input' case-insensitively into a WorkflowOperationId
// note: this method should only be used for API response data and not user input
func ParseWorkflowOperationIDInsensitively(input string) (*WorkflowOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkflowOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkflowOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.WorkflowId, ok = parsed.Parsed["workflowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workflowId", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateWorkflowOperationID checks that 'input' can be parsed as a Workflow Operation ID
func ValidateWorkflowOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkflowOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workflow Operation ID
func (id WorkflowOperationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StorageSync/locations/%s/workflows/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.WorkflowId, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Workflow Operation ID
func (id WorkflowOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorageSync", "Microsoft.StorageSync", "Microsoft.StorageSync"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticWorkflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Workflow Operation ID
func (id WorkflowOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Workflow Operation (%s)", strings.Join(components, "\n"))
}
