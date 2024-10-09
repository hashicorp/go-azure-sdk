package myworkbooksapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MyWorkbookId{})
}

var _ resourceids.ResourceId = &MyWorkbookId{}

// MyWorkbookId is a struct representing the Resource ID for a My Workbook
type MyWorkbookId struct {
	SubscriptionId    string
	ResourceGroupName string
	MyWorkbookName    string
}

// NewMyWorkbookID returns a new MyWorkbookId struct
func NewMyWorkbookID(subscriptionId string, resourceGroupName string, myWorkbookName string) MyWorkbookId {
	return MyWorkbookId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		MyWorkbookName:    myWorkbookName,
	}
}

// ParseMyWorkbookID parses 'input' into a MyWorkbookId
func ParseMyWorkbookID(input string) (*MyWorkbookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MyWorkbookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MyWorkbookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMyWorkbookIDInsensitively parses 'input' case-insensitively into a MyWorkbookId
// note: this method should only be used for API response data and not user input
func ParseMyWorkbookIDInsensitively(input string) (*MyWorkbookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MyWorkbookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MyWorkbookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MyWorkbookId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.MyWorkbookName, ok = input.Parsed["myWorkbookName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "myWorkbookName", input)
	}

	return nil
}

// ValidateMyWorkbookID checks that 'input' can be parsed as a My Workbook ID
func ValidateMyWorkbookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMyWorkbookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted My Workbook ID
func (id MyWorkbookId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/myWorkbooks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MyWorkbookName)
}

// Segments returns a slice of Resource ID Segments which comprise this My Workbook ID
func (id MyWorkbookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticMyWorkbooks", "myWorkbooks", "myWorkbooks"),
		resourceids.UserSpecifiedSegment("myWorkbookName", "myWorkbookName"),
	}
}

// String returns a human-readable description of this My Workbook ID
func (id MyWorkbookId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("My Workbook Name: %q", id.MyWorkbookName),
	}
	return fmt.Sprintf("My Workbook (%s)", strings.Join(components, "\n"))
}
