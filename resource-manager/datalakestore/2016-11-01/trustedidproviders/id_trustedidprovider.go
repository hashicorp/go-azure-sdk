package trustedidproviders

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TrustedIdProviderId{})
}

var _ resourceids.ResourceId = &TrustedIdProviderId{}

// TrustedIdProviderId is a struct representing the Resource ID for a Trusted Id Provider
type TrustedIdProviderId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AccountName           string
	TrustedIdProviderName string
}

// NewTrustedIdProviderID returns a new TrustedIdProviderId struct
func NewTrustedIdProviderID(subscriptionId string, resourceGroupName string, accountName string, trustedIdProviderName string) TrustedIdProviderId {
	return TrustedIdProviderId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AccountName:           accountName,
		TrustedIdProviderName: trustedIdProviderName,
	}
}

// ParseTrustedIdProviderID parses 'input' into a TrustedIdProviderId
func ParseTrustedIdProviderID(input string) (*TrustedIdProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TrustedIdProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TrustedIdProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTrustedIdProviderIDInsensitively parses 'input' case-insensitively into a TrustedIdProviderId
// note: this method should only be used for API response data and not user input
func ParseTrustedIdProviderIDInsensitively(input string) (*TrustedIdProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TrustedIdProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TrustedIdProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TrustedIdProviderId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TrustedIdProviderName, ok = input.Parsed["trustedIdProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "trustedIdProviderName", input)
	}

	return nil
}

// ValidateTrustedIdProviderID checks that 'input' can be parsed as a Trusted Id Provider ID
func ValidateTrustedIdProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTrustedIdProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trusted Id Provider ID
func (id TrustedIdProviderId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataLakeStore/accounts/%s/trustedIdProviders/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.TrustedIdProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Trusted Id Provider ID
func (id TrustedIdProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataLakeStore", "Microsoft.DataLakeStore", "Microsoft.DataLakeStore"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticTrustedIdProviders", "trustedIdProviders", "trustedIdProviders"),
		resourceids.UserSpecifiedSegment("trustedIdProviderName", "trustedIdProviderName"),
	}
}

// String returns a human-readable description of this Trusted Id Provider ID
func (id TrustedIdProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Trusted Id Provider Name: %q", id.TrustedIdProviderName),
	}
	return fmt.Sprintf("Trusted Id Provider (%s)", strings.Join(components, "\n"))
}
