package dsccompilationjob

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CompilationJobStreamId{})
}

var _ resourceids.ResourceId = &CompilationJobStreamId{}

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
	parser := resourceids.NewParserFromResourceIdType(&CompilationJobStreamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CompilationJobStreamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCompilationJobStreamIDInsensitively parses 'input' case-insensitively into a CompilationJobStreamId
// note: this method should only be used for API response data and not user input
func ParseCompilationJobStreamIDInsensitively(input string) (*CompilationJobStreamId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CompilationJobStreamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CompilationJobStreamId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CompilationJobStreamId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AutomationAccountName, ok = input.Parsed["automationAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "automationAccountName", input)
	}

	if id.CompilationJobName, ok = input.Parsed["compilationJobName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "compilationJobName", input)
	}

	if id.JobStreamId, ok = input.Parsed["jobStreamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jobStreamId", input)
	}

	return nil
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
