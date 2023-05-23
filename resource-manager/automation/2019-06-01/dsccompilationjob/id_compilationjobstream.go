package dsccompilationjob

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CompilationJobStreamId{}

// CompilationJobStreamId is a struct representing the Resource ID for a Compilation Job Stream
type CompilationJobStreamId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	CompilationJobName    string
	JobStreamId           string
}

// NewCompilationJobStreamID returns a new CompilationJobStreamId struct
func NewCompilationJobStreamID(subscriptionId string, resourceGroupName string, automationAccountName string, compilationJobName string, jobStreamId string) CompilationJobStreamId {
	return CompilationJobStreamId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		CompilationJobName:    compilationJobName,
		JobStreamId:           jobStreamId,
	}
}

// ParseCompilationJobStreamID parses 'input' into a CompilationJobStreamId
func ParseCompilationJobStreamID(input string) (*CompilationJobStreamId, error) {
	parser := resourceids.NewParserFromResourceIdType(CompilationJobStreamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CompilationJobStreamId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "automationAccountName", *parsed)
	}

	if id.CompilationJobName, ok = parsed.Parsed["compilationJobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "compilationJobName", *parsed)
	}

	if id.JobStreamId, ok = parsed.Parsed["jobStreamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobStreamId", *parsed)
	}

	return &id, nil
}

// ParseCompilationJobStreamIDInsensitively parses 'input' case-insensitively into a CompilationJobStreamId
// note: this method should only be used for API response data and not user input
func ParseCompilationJobStreamIDInsensitively(input string) (*CompilationJobStreamId, error) {
	parser := resourceids.NewParserFromResourceIdType(CompilationJobStreamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CompilationJobStreamId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "automationAccountName", *parsed)
	}

	if id.CompilationJobName, ok = parsed.Parsed["compilationJobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "compilationJobName", *parsed)
	}

	if id.JobStreamId, ok = parsed.Parsed["jobStreamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobStreamId", *parsed)
	}

	return &id, nil
}

// ValidateCompilationJobStreamID checks that 'input' can be parsed as a Compilation Job Stream ID
func ValidateCompilationJobStreamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCompilationJobStreamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Compilation Job Stream ID
func (id CompilationJobStreamId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/compilationJobs/%s/streams/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.CompilationJobName, id.JobStreamId)
}

// Segments returns a slice of Resource ID Segments which comprise this Compilation Job Stream ID
func (id CompilationJobStreamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticCompilationJobs", "compilationJobs", "compilationJobs"),
		resourceids.UserSpecifiedSegment("compilationJobName", "compilationJobValue"),
		resourceids.StaticSegment("staticStreams", "streams", "streams"),
		resourceids.UserSpecifiedSegment("jobStreamId", "jobStreamIdValue"),
	}
}

// String returns a human-readable description of this Compilation Job Stream ID
func (id CompilationJobStreamId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Compilation Job Name: %q", id.CompilationJobName),
		fmt.Sprintf("Job Stream: %q", id.JobStreamId),
	}
	return fmt.Sprintf("Compilation Job Stream (%s)", strings.Join(components, "\n"))
}
