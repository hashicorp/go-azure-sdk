package privatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PrivateLinkForAzureAdId{})
}

var _ resourceids.ResourceId = &PrivateLinkForAzureAdId{}

// PrivateLinkForAzureAdId is a struct representing the Resource ID for a Private Link For Azure Ad
type PrivateLinkForAzureAdId struct {
	SubscriptionId            string
	ResourceGroupName         string
	PrivateLinkForAzureAdName string
}

// NewPrivateLinkForAzureAdID returns a new PrivateLinkForAzureAdId struct
func NewPrivateLinkForAzureAdID(subscriptionId string, resourceGroupName string, privateLinkForAzureAdName string) PrivateLinkForAzureAdId {
	return PrivateLinkForAzureAdId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		PrivateLinkForAzureAdName: privateLinkForAzureAdName,
	}
}

// ParsePrivateLinkForAzureAdID parses 'input' into a PrivateLinkForAzureAdId
func ParsePrivateLinkForAzureAdID(input string) (*PrivateLinkForAzureAdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PrivateLinkForAzureAdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PrivateLinkForAzureAdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePrivateLinkForAzureAdIDInsensitively parses 'input' case-insensitively into a PrivateLinkForAzureAdId
// note: this method should only be used for API response data and not user input
func ParsePrivateLinkForAzureAdIDInsensitively(input string) (*PrivateLinkForAzureAdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PrivateLinkForAzureAdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PrivateLinkForAzureAdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PrivateLinkForAzureAdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateLinkForAzureAdName, ok = input.Parsed["privateLinkForAzureAdName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateLinkForAzureAdName", input)
	}

	return nil
}

// ValidatePrivateLinkForAzureAdID checks that 'input' can be parsed as a Private Link For Azure Ad ID
func ValidatePrivateLinkForAzureAdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrivateLinkForAzureAdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Private Link For Azure Ad ID
func (id PrivateLinkForAzureAdId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AADIAM/privateLinkForAzureAd/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateLinkForAzureAdName)
}

// Segments returns a slice of Resource ID Segments which comprise this Private Link For Azure Ad ID
func (id PrivateLinkForAzureAdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAADIAM", "Microsoft.AADIAM", "Microsoft.AADIAM"),
		resourceids.StaticSegment("staticPrivateLinkForAzureAd", "privateLinkForAzureAd", "privateLinkForAzureAd"),
		resourceids.UserSpecifiedSegment("privateLinkForAzureAdName", "privateLinkForAzureAdName"),
	}
}

// String returns a human-readable description of this Private Link For Azure Ad ID
func (id PrivateLinkForAzureAdId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Link For Azure Ad Name: %q", id.PrivateLinkForAzureAdName),
	}
	return fmt.Sprintf("Private Link For Azure Ad (%s)", strings.Join(components, "\n"))
}
