package experiments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ExecutionDetailId{})
}

var _ resourceids.ResourceId = &ExecutionDetailId{}

// ExecutionDetailId is a struct representing the Resource ID for a Execution Detail
type ExecutionDetailId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ExperimentName     string
	ExecutionDetailsId string
}

// NewExecutionDetailID returns a new ExecutionDetailId struct
func NewExecutionDetailID(subscriptionId string, resourceGroupName string, experimentName string, executionDetailsId string) ExecutionDetailId {
	return ExecutionDetailId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ExperimentName:     experimentName,
		ExecutionDetailsId: executionDetailsId,
	}
}

// ParseExecutionDetailID parses 'input' into a ExecutionDetailId
func ParseExecutionDetailID(input string) (*ExecutionDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExecutionDetailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExecutionDetailId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExecutionDetailIDInsensitively parses 'input' case-insensitively into a ExecutionDetailId
// note: this method should only be used for API response data and not user input
func ParseExecutionDetailIDInsensitively(input string) (*ExecutionDetailId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExecutionDetailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExecutionDetailId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExecutionDetailId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ExperimentName, ok = input.Parsed["experimentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "experimentName", input)
	}

	if id.ExecutionDetailsId, ok = input.Parsed["executionDetailsId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "executionDetailsId", input)
	}

	return nil
}

// ValidateExecutionDetailID checks that 'input' can be parsed as a Execution Detail ID
func ValidateExecutionDetailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExecutionDetailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Execution Detail ID
func (id ExecutionDetailId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Chaos/experiments/%s/executionDetails/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ExperimentName, id.ExecutionDetailsId)
}

// Segments returns a slice of Resource ID Segments which comprise this Execution Detail ID
func (id ExecutionDetailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticExperiments", "experiments", "experiments"),
		resourceids.UserSpecifiedSegment("experimentName", "experimentValue"),
		resourceids.StaticSegment("staticExecutionDetails", "executionDetails", "executionDetails"),
		resourceids.UserSpecifiedSegment("executionDetailsId", "executionDetailsIdValue"),
	}
}

// String returns a human-readable description of this Execution Detail ID
func (id ExecutionDetailId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Experiment Name: %q", id.ExperimentName),
		fmt.Sprintf("Execution Details: %q", id.ExecutionDetailsId),
	}
	return fmt.Sprintf("Execution Detail (%s)", strings.Join(components, "\n"))
}
