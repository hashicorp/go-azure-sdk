package jobtargetexecutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ExecutionStepId{}

// ExecutionStepId is a struct representing the Resource ID for a Execution Step
type ExecutionStepId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	JobAgentName      string
	JobName           string
	JobExecutionId    string
	StepName          string
}

// NewExecutionStepID returns a new ExecutionStepId struct
func NewExecutionStepID(subscriptionId string, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionId string, stepName string) ExecutionStepId {
	return ExecutionStepId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		JobAgentName:      jobAgentName,
		JobName:           jobName,
		JobExecutionId:    jobExecutionId,
		StepName:          stepName,
	}
}

// ParseExecutionStepID parses 'input' into a ExecutionStepId
func ParseExecutionStepID(input string) (*ExecutionStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExecutionStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExecutionStepId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.JobAgentName, ok = parsed.Parsed["jobAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", *parsed)
	}

	if id.JobName, ok = parsed.Parsed["jobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobName", *parsed)
	}

	if id.JobExecutionId, ok = parsed.Parsed["jobExecutionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobExecutionId", *parsed)
	}

	if id.StepName, ok = parsed.Parsed["stepName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "stepName", *parsed)
	}

	return &id, nil
}

// ParseExecutionStepIDInsensitively parses 'input' case-insensitively into a ExecutionStepId
// note: this method should only be used for API response data and not user input
func ParseExecutionStepIDInsensitively(input string) (*ExecutionStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExecutionStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExecutionStepId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.JobAgentName, ok = parsed.Parsed["jobAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", *parsed)
	}

	if id.JobName, ok = parsed.Parsed["jobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobName", *parsed)
	}

	if id.JobExecutionId, ok = parsed.Parsed["jobExecutionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobExecutionId", *parsed)
	}

	if id.StepName, ok = parsed.Parsed["stepName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "stepName", *parsed)
	}

	return &id, nil
}

// ValidateExecutionStepID checks that 'input' can be parsed as a Execution Step ID
func ValidateExecutionStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExecutionStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Execution Step ID
func (id ExecutionStepId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/jobAgents/%s/jobs/%s/executions/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.JobAgentName, id.JobName, id.JobExecutionId, id.StepName)
}

// Segments returns a slice of Resource ID Segments which comprise this Execution Step ID
func (id ExecutionStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticJobAgents", "jobAgents", "jobAgents"),
		resourceids.UserSpecifiedSegment("jobAgentName", "jobAgentValue"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("jobName", "jobValue"),
		resourceids.StaticSegment("staticExecutions", "executions", "executions"),
		resourceids.UserSpecifiedSegment("jobExecutionId", "jobExecutionIdValue"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("stepName", "stepValue"),
	}
}

// String returns a human-readable description of this Execution Step ID
func (id ExecutionStepId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Job Agent Name: %q", id.JobAgentName),
		fmt.Sprintf("Job Name: %q", id.JobName),
		fmt.Sprintf("Job Execution: %q", id.JobExecutionId),
		fmt.Sprintf("Step Name: %q", id.StepName),
	}
	return fmt.Sprintf("Execution Step (%s)", strings.Join(components, "\n"))
}
