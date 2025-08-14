package get

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ServiceTaskId{})
}

var _ resourceids.ResourceId = &ServiceTaskId{}

// ServiceTaskId is a struct representing the Resource ID for a Service Task
type ServiceTaskId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	ServiceTaskName   string
}

// NewServiceTaskID returns a new ServiceTaskId struct
func NewServiceTaskID(subscriptionId string, resourceGroupName string, serviceName string, serviceTaskName string) ServiceTaskId {
	return ServiceTaskId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		ServiceTaskName:   serviceTaskName,
	}
}

// ParseServiceTaskID parses 'input' into a ServiceTaskId
func ParseServiceTaskID(input string) (*ServiceTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServiceTaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServiceTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServiceTaskIDInsensitively parses 'input' case-insensitively into a ServiceTaskId
// note: this method should only be used for API response data and not user input
func ParseServiceTaskIDInsensitively(input string) (*ServiceTaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServiceTaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServiceTaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServiceTaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServiceName, ok = input.Parsed["serviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceName", input)
	}

	if id.ServiceTaskName, ok = input.Parsed["serviceTaskName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceTaskName", input)
	}

	return nil
}

// ValidateServiceTaskID checks that 'input' can be parsed as a Service Task ID
func ValidateServiceTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServiceTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Task ID
func (id ServiceTaskId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataMigration/services/%s/serviceTasks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.ServiceTaskName)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Task ID
func (id ServiceTaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.UserSpecifiedSegment("resourceGroupName", "resourceGroupName"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticServices", "services", "services"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceName"),
		resourceids.StaticSegment("staticServiceTasks", "serviceTasks", "serviceTasks"),
		resourceids.UserSpecifiedSegment("serviceTaskName", "serviceTaskName"),
	}
}

// String returns a human-readable description of this Service Task ID
func (id ServiceTaskId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Service Task Name: %q", id.ServiceTaskName),
	}
	return fmt.Sprintf("Service Task (%s)", strings.Join(components, "\n"))
}
