package jobtargetexecutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TargetId{})
}

var _ resourceids.ResourceId = &TargetId{}

// TargetId is a struct representing the Resource ID for a Target
type TargetId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	JobAgentName      string
	JobName           string
	JobExecutionId    string
	StepName          string
	TargetId          string
}

// NewTargetID returns a new TargetId struct
func NewTargetID(subscriptionId string, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionId string, stepName string, targetId string) TargetId {
	return TargetId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		JobAgentName:      jobAgentName,
		JobName:           jobName,
		JobExecutionId:    jobExecutionId,
		StepName:          stepName,
		TargetId:          targetId,
	}
}

// ParseTargetID parses 'input' into a TargetId
func ParseTargetID(input string) (*TargetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TargetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TargetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTargetIDInsensitively parses 'input' case-insensitively into a TargetId
// note: this method should only be used for API response data and not user input
func ParseTargetIDInsensitively(input string) (*TargetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TargetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TargetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TargetId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.JobAgentName, ok = input.Parsed["jobAgentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", input)
	}

	if id.JobName, ok = input.Parsed["jobName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jobName", input)
	}

	if id.JobExecutionId, ok = input.Parsed["jobExecutionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jobExecutionId", input)
	}

	if id.StepName, ok = input.Parsed["stepName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "stepName", input)
	}

	if id.TargetId, ok = input.Parsed["targetId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "targetId", input)
	}

	return nil
}

// ValidateTargetID checks that 'input' can be parsed as a Target ID
func ValidateTargetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTargetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Target ID
func (id TargetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/jobAgents/%s/jobs/%s/executions/%s/steps/%s/targets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.JobAgentName, id.JobName, id.JobExecutionId, id.StepName, id.TargetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Target ID
func (id TargetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverName"),
		resourceids.StaticSegment("staticJobAgents", "jobAgents", "jobAgents"),
		resourceids.UserSpecifiedSegment("jobAgentName", "jobAgentName"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("jobName", "jobName"),
		resourceids.StaticSegment("staticExecutions", "executions", "executions"),
		resourceids.UserSpecifiedSegment("jobExecutionId", "jobExecutionId"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("stepName", "stepName"),
		resourceids.StaticSegment("staticTargets", "targets", "targets"),
		resourceids.UserSpecifiedSegment("targetId", "targetId"),
	}
}

// String returns a human-readable description of this Target ID
func (id TargetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Job Agent Name: %q", id.JobAgentName),
		fmt.Sprintf("Job Name: %q", id.JobName),
		fmt.Sprintf("Job Execution: %q", id.JobExecutionId),
		fmt.Sprintf("Step Name: %q", id.StepName),
		fmt.Sprintf("Target: %q", id.TargetId),
	}
	return fmt.Sprintf("Target (%s)", strings.Join(components, "\n"))
}
