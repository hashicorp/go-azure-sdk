package raitoollabels

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RaiToolLabelId{})
}

var _ resourceids.ResourceId = &RaiToolLabelId{}

// RaiToolLabelId is a struct representing the Resource ID for a Rai Tool Label
type RaiToolLabelId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	RaiToolLabelName  string
}

// NewRaiToolLabelID returns a new RaiToolLabelId struct
func NewRaiToolLabelID(subscriptionId string, resourceGroupName string, accountName string, raiToolLabelName string) RaiToolLabelId {
	return RaiToolLabelId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		RaiToolLabelName:  raiToolLabelName,
	}
}

// ParseRaiToolLabelID parses 'input' into a RaiToolLabelId
func ParseRaiToolLabelID(input string) (*RaiToolLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiToolLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiToolLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRaiToolLabelIDInsensitively parses 'input' case-insensitively into a RaiToolLabelId
// note: this method should only be used for API response data and not user input
func ParseRaiToolLabelIDInsensitively(input string) (*RaiToolLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiToolLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiToolLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RaiToolLabelId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RaiToolLabelName, ok = input.Parsed["raiToolLabelName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "raiToolLabelName", input)
	}

	return nil
}

// ValidateRaiToolLabelID checks that 'input' can be parsed as a Rai Tool Label ID
func ValidateRaiToolLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRaiToolLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rai Tool Label ID
func (id RaiToolLabelId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/raiToolLabels/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.RaiToolLabelName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rai Tool Label ID
func (id RaiToolLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticRaiToolLabels", "raiToolLabels", "raiToolLabels"),
		resourceids.UserSpecifiedSegment("raiToolLabelName", "raiToolLabelName"),
	}
}

// String returns a human-readable description of this Rai Tool Label ID
func (id RaiToolLabelId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Rai Tool Label Name: %q", id.RaiToolLabelName),
	}
	return fmt.Sprintf("Rai Tool Label (%s)", strings.Join(components, "\n"))
}
