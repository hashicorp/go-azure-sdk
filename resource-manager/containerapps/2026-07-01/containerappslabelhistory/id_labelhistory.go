package containerappslabelhistory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LabelHistoryId{})
}

var _ resourceids.ResourceId = &LabelHistoryId{}

// LabelHistoryId is a struct representing the Resource ID for a Label History
type LabelHistoryId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
	LabelHistoryName  string
}

// NewLabelHistoryID returns a new LabelHistoryId struct
func NewLabelHistoryID(subscriptionId string, resourceGroupName string, containerAppName string, labelHistoryName string) LabelHistoryId {
	return LabelHistoryId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
		LabelHistoryName:  labelHistoryName,
	}
}

// ParseLabelHistoryID parses 'input' into a LabelHistoryId
func ParseLabelHistoryID(input string) (*LabelHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LabelHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LabelHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLabelHistoryIDInsensitively parses 'input' case-insensitively into a LabelHistoryId
// note: this method should only be used for API response data and not user input
func ParseLabelHistoryIDInsensitively(input string) (*LabelHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LabelHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LabelHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LabelHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ContainerAppName, ok = input.Parsed["containerAppName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "containerAppName", input)
	}

	if id.LabelHistoryName, ok = input.Parsed["labelHistoryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labelHistoryName", input)
	}

	return nil
}

// ValidateLabelHistoryID checks that 'input' can be parsed as a Label History ID
func ValidateLabelHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLabelHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Label History ID
func (id LabelHistoryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/containerApps/%s/labelHistories/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName, id.LabelHistoryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Label History ID
func (id LabelHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppName"),
		resourceids.StaticSegment("staticLabelHistories", "labelHistories", "labelHistories"),
		resourceids.UserSpecifiedSegment("labelHistoryName", "labelHistoryName"),
	}
}

// String returns a human-readable description of this Label History ID
func (id LabelHistoryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
		fmt.Sprintf("Label History Name: %q", id.LabelHistoryName),
	}
	return fmt.Sprintf("Label History (%s)", strings.Join(components, "\n"))
}
