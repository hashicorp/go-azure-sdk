package scriptexecutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScriptExecutionId{}

// ScriptExecutionId is a struct representing the Resource ID for a Script Execution
type ScriptExecutionId struct {
	SubscriptionId      string
	ResourceGroupName   string
	PrivateCloudName    string
	ScriptExecutionName string
}

// NewScriptExecutionID returns a new ScriptExecutionId struct
func NewScriptExecutionID(subscriptionId string, resourceGroupName string, privateCloudName string, scriptExecutionName string) ScriptExecutionId {
	return ScriptExecutionId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		PrivateCloudName:    privateCloudName,
		ScriptExecutionName: scriptExecutionName,
	}
}

// ParseScriptExecutionID parses 'input' into a ScriptExecutionId
func ParseScriptExecutionID(input string) (*ScriptExecutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScriptExecutionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScriptExecutionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScriptExecutionIDInsensitively parses 'input' case-insensitively into a ScriptExecutionId
// note: this method should only be used for API response data and not user input
func ParseScriptExecutionIDInsensitively(input string) (*ScriptExecutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScriptExecutionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScriptExecutionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScriptExecutionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateCloudName, ok = input.Parsed["privateCloudName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateCloudName", input)
	}

	if id.ScriptExecutionName, ok = input.Parsed["scriptExecutionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scriptExecutionName", input)
	}

	return nil
}

// ValidateScriptExecutionID checks that 'input' can be parsed as a Script Execution ID
func ValidateScriptExecutionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScriptExecutionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Script Execution ID
func (id ScriptExecutionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/scriptExecutions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ScriptExecutionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Script Execution ID
func (id ScriptExecutionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticScriptExecutions", "scriptExecutions", "scriptExecutions"),
		resourceids.UserSpecifiedSegment("scriptExecutionName", "scriptExecutionValue"),
	}
}

// String returns a human-readable description of this Script Execution ID
func (id ScriptExecutionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Script Execution Name: %q", id.ScriptExecutionName),
	}
	return fmt.Sprintf("Script Execution (%s)", strings.Join(components, "\n"))
}
