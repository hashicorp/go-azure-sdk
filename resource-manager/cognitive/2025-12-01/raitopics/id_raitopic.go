package raitopics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RaitopicId{})
}

var _ resourceids.ResourceId = &RaitopicId{}

// RaitopicId is a struct representing the Resource ID for a Raitopic
type RaitopicId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	RaitopicName      string
}

// NewRaitopicID returns a new RaitopicId struct
func NewRaitopicID(subscriptionId string, resourceGroupName string, accountName string, raitopicName string) RaitopicId {
	return RaitopicId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		RaitopicName:      raitopicName,
	}
}

// ParseRaitopicID parses 'input' into a RaitopicId
func ParseRaitopicID(input string) (*RaitopicId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaitopicId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaitopicId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRaitopicIDInsensitively parses 'input' case-insensitively into a RaitopicId
// note: this method should only be used for API response data and not user input
func ParseRaitopicIDInsensitively(input string) (*RaitopicId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaitopicId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaitopicId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RaitopicId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AccountName, ok = input.Parsed["accountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountName", input)
	}

	if id.RaitopicName, ok = input.Parsed["raitopicName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "raitopicName", input)
	}

	return nil
}

// ValidateRaitopicID checks that 'input' can be parsed as a Raitopic ID
func ValidateRaitopicID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRaitopicID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Raitopic ID
func (id RaitopicId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/raitopics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.RaitopicName)
}

// Segments returns a slice of Resource ID Segments which comprise this Raitopic ID
func (id RaitopicId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticRaitopics", "raitopics", "raitopics"),
		resourceids.UserSpecifiedSegment("raitopicName", "raitopicName"),
	}
}

// String returns a human-readable description of this Raitopic ID
func (id RaitopicId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Raitopic Name: %q", id.RaitopicName),
	}
	return fmt.Sprintf("Raitopic (%s)", strings.Join(components, "\n"))
}
