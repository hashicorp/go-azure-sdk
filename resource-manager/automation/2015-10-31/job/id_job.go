// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = JobId{}

// JobId is a struct representing the Resource ID for a Job
type JobId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	JobId                 string
}

// NewJobID returns a new JobId struct
func NewJobID(subscriptionId string, resourceGroupName string, automationAccountName string, jobId string) JobId {
	return JobId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		JobId:                 jobId,
	}
}

// ParseJobID parses 'input' into a JobId
func ParseJobID(input string) (*JobId, error) {
	parser := resourceids.NewParserFromResourceIdType(JobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.JobId, ok = parsed.Parsed["jobId"]; !ok {
		return nil, fmt.Errorf("the segment 'jobId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseJobIDInsensitively parses 'input' case-insensitively into a JobId
// note: this method should only be used for API response data and not user input
func ParseJobIDInsensitively(input string) (*JobId, error) {
	parser := resourceids.NewParserFromResourceIdType(JobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JobId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, fmt.Errorf("the segment 'automationAccountName' was not found in the resource id %q", input)
	}

	if id.JobId, ok = parsed.Parsed["jobId"]; !ok {
		return nil, fmt.Errorf("the segment 'jobId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateJobID checks that 'input' can be parsed as a Job ID
func ValidateJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Job ID
func (id JobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/jobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.JobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Job ID
func (id JobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("jobId", "jobIdValue"),
	}
}

// String returns a human-readable description of this Job ID
func (id JobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Job: %q", id.JobId),
	}
	return fmt.Sprintf("Job (%s)", strings.Join(components, "\n"))
}
