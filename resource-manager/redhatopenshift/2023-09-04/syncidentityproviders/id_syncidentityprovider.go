package syncidentityproviders

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SyncIdentityProviderId{})
}

var _ resourceids.ResourceId = &SyncIdentityProviderId{}

// SyncIdentityProviderId is a struct representing the Resource ID for a Sync Identity Provider
type SyncIdentityProviderId struct {
	SubscriptionId           string
	ResourceGroupName        string
	OpenShiftClusterName     string
	SyncIdentityProviderName string
}

// NewSyncIdentityProviderID returns a new SyncIdentityProviderId struct
func NewSyncIdentityProviderID(subscriptionId string, resourceGroupName string, openShiftClusterName string, syncIdentityProviderName string) SyncIdentityProviderId {
	return SyncIdentityProviderId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		OpenShiftClusterName:     openShiftClusterName,
		SyncIdentityProviderName: syncIdentityProviderName,
	}
}

// ParseSyncIdentityProviderID parses 'input' into a SyncIdentityProviderId
func ParseSyncIdentityProviderID(input string) (*SyncIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SyncIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SyncIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSyncIdentityProviderIDInsensitively parses 'input' case-insensitively into a SyncIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseSyncIdentityProviderIDInsensitively(input string) (*SyncIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SyncIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SyncIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SyncIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.OpenShiftClusterName, ok = input.Parsed["openShiftClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftClusterName", input)
	}

	if id.SyncIdentityProviderName, ok = input.Parsed["syncIdentityProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "syncIdentityProviderName", input)
	}

	return nil
}

// ValidateSyncIdentityProviderID checks that 'input' can be parsed as a Sync Identity Provider ID
func ValidateSyncIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSyncIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sync Identity Provider ID
func (id SyncIdentityProviderId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RedHatOpenShift/openShiftClusters/%s/syncIdentityProvider/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.OpenShiftClusterName, id.SyncIdentityProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sync Identity Provider ID
func (id SyncIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticOpenShiftClusters", "openShiftClusters", "openShiftClusters"),
		resourceids.UserSpecifiedSegment("openShiftClusterName", "openShiftClusterName"),
		resourceids.StaticSegment("staticSyncIdentityProvider", "syncIdentityProvider", "syncIdentityProvider"),
		resourceids.UserSpecifiedSegment("syncIdentityProviderName", "syncIdentityProviderName"),
	}
}

// String returns a human-readable description of this Sync Identity Provider ID
func (id SyncIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Open Shift Cluster Name: %q", id.OpenShiftClusterName),
		fmt.Sprintf("Sync Identity Provider Name: %q", id.SyncIdentityProviderName),
	}
	return fmt.Sprintf("Sync Identity Provider (%s)", strings.Join(components, "\n"))
}
