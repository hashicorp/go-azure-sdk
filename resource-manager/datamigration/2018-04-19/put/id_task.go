// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package put

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = TaskId{}

// TaskId is a struct representing the Resource ID for a Task
type TaskId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ProjectName       string
	TaskName          string
}

// NewTaskID returns a new TaskId struct
func NewTaskID(subscriptionId string, resourceGroupName string, serviceName string, projectName string, taskName string) TaskId {
	return TaskId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ProjectName:       projectName,
		TaskName:          taskName,
	}
}

// ParseTaskID parses 'input' into a TaskId
func ParseTaskID(input string) (*TaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(TaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TaskId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ProjectName, ok = parsed.Parsed["projectName"]; !ok {
		return nil, fmt.Errorf("the segment 'projectName' was not found in the resource id %q", input)
	}

	if id.TaskName, ok = parsed.Parsed["taskName"]; !ok {
		return nil, fmt.Errorf("the segment 'taskName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseTaskIDInsensitively parses 'input' case-insensitively into a TaskId
// note: this method should only be used for API response data and not user input
func ParseTaskIDInsensitively(input string) (*TaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(TaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TaskId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, fmt.Errorf("the segment 'serviceName' was not found in the resource id %q", input)
	}

	if id.ProjectName, ok = parsed.Parsed["projectName"]; !ok {
		return nil, fmt.Errorf("the segment 'projectName' was not found in the resource id %q", input)
	}

	if id.TaskName, ok = parsed.Parsed["taskName"]; !ok {
		return nil, fmt.Errorf("the segment 'taskName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateTaskID checks that 'input' can be parsed as a Task ID
func ValidateTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Task ID
func (id TaskId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataMigration/services/%s/projects/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ProjectName, id.TaskName)
}

// Segments returns a slice of Resource ID Segments which comprise this Task ID
func (id TaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.UserSpecifiedSegment("resourceGroupName", "resourceGroupValue"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticServices", "services", "services"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectName", "projectValue"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskName", "taskValue"),
	}
}

// String returns a human-readable description of this Task ID
func (id TaskId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
		fmt.Sprintf("Task Name: %q", id.TaskName),
	}
	return fmt.Sprintf("Task (%s)", strings.Join(components, "\n"))
}
