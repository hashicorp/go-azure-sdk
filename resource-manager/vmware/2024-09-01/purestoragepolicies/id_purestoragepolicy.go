package purestoragepolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PureStoragePolicyId{})
}

var _ resourceids.ResourceId = &PureStoragePolicyId{}

// PureStoragePolicyId is a struct representing the Resource ID for a Pure Storage Policy
type PureStoragePolicyId struct {
	SubscriptionId        string
	ResourceGroupName     string
	PrivateCloudName      string
	PureStoragePolicyName string
}

// NewPureStoragePolicyID returns a new PureStoragePolicyId struct
func NewPureStoragePolicyID(subscriptionId string, resourceGroupName string, privateCloudName string, pureStoragePolicyName string) PureStoragePolicyId {
	return PureStoragePolicyId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		PrivateCloudName:      privateCloudName,
		PureStoragePolicyName: pureStoragePolicyName,
	}
}

// ParsePureStoragePolicyID parses 'input' into a PureStoragePolicyId
func ParsePureStoragePolicyID(input string) (*PureStoragePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PureStoragePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PureStoragePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePureStoragePolicyIDInsensitively parses 'input' case-insensitively into a PureStoragePolicyId
// note: this method should only be used for API response data and not user input
func ParsePureStoragePolicyIDInsensitively(input string) (*PureStoragePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PureStoragePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PureStoragePolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PureStoragePolicyId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PureStoragePolicyName, ok = input.Parsed["pureStoragePolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pureStoragePolicyName", input)
	}

	return nil
}

// ValidatePureStoragePolicyID checks that 'input' can be parsed as a Pure Storage Policy ID
func ValidatePureStoragePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePureStoragePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pure Storage Policy ID
func (id PureStoragePolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/pureStoragePolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.PureStoragePolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Pure Storage Policy ID
func (id PureStoragePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudName"),
		resourceids.StaticSegment("staticPureStoragePolicies", "pureStoragePolicies", "pureStoragePolicies"),
		resourceids.UserSpecifiedSegment("pureStoragePolicyName", "pureStoragePolicyName"),
	}
}

// String returns a human-readable description of this Pure Storage Policy ID
func (id PureStoragePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Pure Storage Policy Name: %q", id.PureStoragePolicyName),
	}
	return fmt.Sprintf("Pure Storage Policy (%s)", strings.Join(components, "\n"))
}
