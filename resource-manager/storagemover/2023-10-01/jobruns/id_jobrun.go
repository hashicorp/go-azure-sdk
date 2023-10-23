package jobruns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = JobRunId{}

// JobRunId is a struct representing the Resource ID for a Job Run
type JobRunId struct {
	SubscriptionId    string
	ResourceGroupName string
	StorageMoverName  string
	ProjectName       string
	JobDefinitionName string
	JobRunName        string
}

// NewJobRunID returns a new JobRunId struct
func NewJobRunID(subscriptionId string, resourceGroupName string, storageMoverName string, projectName string, jobDefinitionName string, jobRunName string) JobRunId {
	return JobRunId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		StorageMoverName:  storageMoverName,
		ProjectName:       projectName,
		JobDefinitionName: jobDefinitionName,
		JobRunName:        jobRunName,
	}
}

// ParseJobRunID parses 'input' into a JobRunId
func ParseJobRunID(input string) (*JobRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(JobRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JobRunId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.StorageMoverName, ok = parsed.Parsed["storageMoverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageMoverName", *parsed)
	}

	if id.ProjectName, ok = parsed.Parsed["projectName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "projectName", *parsed)
	}

	if id.JobDefinitionName, ok = parsed.Parsed["jobDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobDefinitionName", *parsed)
	}

	if id.JobRunName, ok = parsed.Parsed["jobRunName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobRunName", *parsed)
	}

	return &id, nil
}

// ParseJobRunIDInsensitively parses 'input' case-insensitively into a JobRunId
// note: this method should only be used for API response data and not user input
func ParseJobRunIDInsensitively(input string) (*JobRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(JobRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JobRunId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.StorageMoverName, ok = parsed.Parsed["storageMoverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageMoverName", *parsed)
	}

	if id.ProjectName, ok = parsed.Parsed["projectName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "projectName", *parsed)
	}

	if id.JobDefinitionName, ok = parsed.Parsed["jobDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobDefinitionName", *parsed)
	}

	if id.JobRunName, ok = parsed.Parsed["jobRunName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobRunName", *parsed)
	}

	return &id, nil
}

// ValidateJobRunID checks that 'input' can be parsed as a Job Run ID
func ValidateJobRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseJobRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Job Run ID
func (id JobRunId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StorageMover/storageMovers/%s/projects/%s/jobDefinitions/%s/jobRuns/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StorageMoverName, id.ProjectName, id.JobDefinitionName, id.JobRunName)
}

// Segments returns a slice of Resource ID Segments which comprise this Job Run ID
func (id JobRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorageMover", "Microsoft.StorageMover", "Microsoft.StorageMover"),
		resourceids.StaticSegment("staticStorageMovers", "storageMovers", "storageMovers"),
		resourceids.UserSpecifiedSegment("storageMoverName", "storageMoverValue"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectName", "projectValue"),
		resourceids.StaticSegment("staticJobDefinitions", "jobDefinitions", "jobDefinitions"),
		resourceids.UserSpecifiedSegment("jobDefinitionName", "jobDefinitionValue"),
		resourceids.StaticSegment("staticJobRuns", "jobRuns", "jobRuns"),
		resourceids.UserSpecifiedSegment("jobRunName", "jobRunValue"),
	}
}

// String returns a human-readable description of this Job Run ID
func (id JobRunId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Storage Mover Name: %q", id.StorageMoverName),
		fmt.Sprintf("Project Name: %q", id.ProjectName),
		fmt.Sprintf("Job Definition Name: %q", id.JobDefinitionName),
		fmt.Sprintf("Job Run Name: %q", id.JobRunName),
	}
	return fmt.Sprintf("Job Run (%s)", strings.Join(components, "\n"))
}
