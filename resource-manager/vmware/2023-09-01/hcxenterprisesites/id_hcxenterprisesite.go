package hcxenterprisesites

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&HcxEnterpriseSiteId{})
}

var _ resourceids.ResourceId = &HcxEnterpriseSiteId{}

// HcxEnterpriseSiteId is a struct representing the Resource ID for a Hcx Enterprise Site
type HcxEnterpriseSiteId struct {
	SubscriptionId        string
	ResourceGroupName     string
	PrivateCloudName      string
	HcxEnterpriseSiteName string
}

// NewHcxEnterpriseSiteID returns a new HcxEnterpriseSiteId struct
func NewHcxEnterpriseSiteID(subscriptionId string, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string) HcxEnterpriseSiteId {
	return HcxEnterpriseSiteId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		PrivateCloudName:      privateCloudName,
		HcxEnterpriseSiteName: hcxEnterpriseSiteName,
	}
}

// ParseHcxEnterpriseSiteID parses 'input' into a HcxEnterpriseSiteId
func ParseHcxEnterpriseSiteID(input string) (*HcxEnterpriseSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HcxEnterpriseSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HcxEnterpriseSiteId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHcxEnterpriseSiteIDInsensitively parses 'input' case-insensitively into a HcxEnterpriseSiteId
// note: this method should only be used for API response data and not user input
func ParseHcxEnterpriseSiteIDInsensitively(input string) (*HcxEnterpriseSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HcxEnterpriseSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HcxEnterpriseSiteId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HcxEnterpriseSiteId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.HcxEnterpriseSiteName, ok = input.Parsed["hcxEnterpriseSiteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hcxEnterpriseSiteName", input)
	}

	return nil
}

// ValidateHcxEnterpriseSiteID checks that 'input' can be parsed as a Hcx Enterprise Site ID
func ValidateHcxEnterpriseSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHcxEnterpriseSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hcx Enterprise Site ID
func (id HcxEnterpriseSiteId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/hcxEnterpriseSites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.HcxEnterpriseSiteName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hcx Enterprise Site ID
func (id HcxEnterpriseSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticHcxEnterpriseSites", "hcxEnterpriseSites", "hcxEnterpriseSites"),
		resourceids.UserSpecifiedSegment("hcxEnterpriseSiteName", "hcxEnterpriseSiteValue"),
	}
}

// String returns a human-readable description of this Hcx Enterprise Site ID
func (id HcxEnterpriseSiteId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Hcx Enterprise Site Name: %q", id.HcxEnterpriseSiteName),
	}
	return fmt.Sprintf("Hcx Enterprise Site (%s)", strings.Join(components, "\n"))
}
