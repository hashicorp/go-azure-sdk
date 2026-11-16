package parametergroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProviderParameterGroupId{})
}

var _ resourceids.ResourceId = &ProviderParameterGroupId{}

// ProviderParameterGroupId is a struct representing the Resource ID for a Provider Parameter Group
type ProviderParameterGroupId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ParameterGroupName string
}

// NewProviderParameterGroupID returns a new ProviderParameterGroupId struct
func NewProviderParameterGroupID(subscriptionId string, resourceGroupName string, parameterGroupName string) ProviderParameterGroupId {
	return ProviderParameterGroupId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ParameterGroupName: parameterGroupName,
	}
}

// ParseProviderParameterGroupID parses 'input' into a ProviderParameterGroupId
func ParseProviderParameterGroupID(input string) (*ProviderParameterGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderParameterGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderParameterGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderParameterGroupIDInsensitively parses 'input' case-insensitively into a ProviderParameterGroupId
// note: this method should only be used for API response data and not user input
func ParseProviderParameterGroupIDInsensitively(input string) (*ProviderParameterGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderParameterGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderParameterGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderParameterGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ParameterGroupName, ok = input.Parsed["parameterGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "parameterGroupName", input)
	}

	return nil
}

// ValidateProviderParameterGroupID checks that 'input' can be parsed as a Provider Parameter Group ID
func ValidateProviderParameterGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderParameterGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Parameter Group ID
func (id ProviderParameterGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/parameterGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ParameterGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Parameter Group ID
func (id ProviderParameterGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHorizonDb", "Microsoft.HorizonDb", "Microsoft.HorizonDb"),
		resourceids.StaticSegment("staticParameterGroups", "parameterGroups", "parameterGroups"),
		resourceids.UserSpecifiedSegment("parameterGroupName", "parameterGroupName"),
	}
}

// String returns a human-readable description of this Provider Parameter Group ID
func (id ProviderParameterGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Parameter Group Name: %q", id.ParameterGroupName),
	}
	return fmt.Sprintf("Provider Parameter Group (%s)", strings.Join(components, "\n"))
}
