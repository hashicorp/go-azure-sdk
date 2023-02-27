package privateendpointconnection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MasterSiteId{}

// MasterSiteId is a struct representing the Resource ID for a Master Site
type MasterSiteId struct {
	SubscriptionId    string
	ResourceGroupName string
	MasterSiteName    string
}

// NewMasterSiteID returns a new MasterSiteId struct
func NewMasterSiteID(subscriptionId string, resourceGroupName string, masterSiteName string) MasterSiteId {
	return MasterSiteId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		MasterSiteName:    masterSiteName,
	}
}

// ParseMasterSiteID parses 'input' into a MasterSiteId
func ParseMasterSiteID(input string) (*MasterSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(MasterSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MasterSiteId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.MasterSiteName, ok = parsed.Parsed["masterSiteName"]; !ok {
		return nil, fmt.Errorf("the segment 'masterSiteName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseMasterSiteIDInsensitively parses 'input' case-insensitively into a MasterSiteId
// note: this method should only be used for API response data and not user input
func ParseMasterSiteIDInsensitively(input string) (*MasterSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(MasterSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MasterSiteId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.MasterSiteName, ok = parsed.Parsed["masterSiteName"]; !ok {
		return nil, fmt.Errorf("the segment 'masterSiteName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateMasterSiteID checks that 'input' can be parsed as a Master Site ID
func ValidateMasterSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMasterSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Master Site ID
func (id MasterSiteId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OffAzure/masterSites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MasterSiteName)
}

// Segments returns a slice of Resource ID Segments which comprise this Master Site ID
func (id MasterSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOffAzure", "Microsoft.OffAzure", "Microsoft.OffAzure"),
		resourceids.StaticSegment("staticMasterSites", "masterSites", "masterSites"),
		resourceids.UserSpecifiedSegment("masterSiteName", "masterSiteValue"),
	}
}

// String returns a human-readable description of this Master Site ID
func (id MasterSiteId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Master Site Name: %q", id.MasterSiteName),
	}
	return fmt.Sprintf("Master Site (%s)", strings.Join(components, "\n"))
}
