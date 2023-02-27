// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syncsets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SyncSetId{}

// SyncSetId is a struct representing the Resource ID for a Sync Set
type SyncSetId struct {
	SubscriptionId       string
	ResourceGroupName    string
	OpenShiftClusterName string
	SyncSetName          string
}

// NewSyncSetID returns a new SyncSetId struct
func NewSyncSetID(subscriptionId string, resourceGroupName string, openShiftClusterName string, syncSetName string) SyncSetId {
	return SyncSetId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		OpenShiftClusterName: openShiftClusterName,
		SyncSetName:          syncSetName,
	}
}

// ParseSyncSetID parses 'input' into a SyncSetId
func ParseSyncSetID(input string) (*SyncSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(SyncSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SyncSetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.OpenShiftClusterName, ok = parsed.Parsed["openShiftClusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'openShiftClusterName' was not found in the resource id %q", input)
	}

	if id.SyncSetName, ok = parsed.Parsed["syncSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'syncSetName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSyncSetIDInsensitively parses 'input' case-insensitively into a SyncSetId
// note: this method should only be used for API response data and not user input
func ParseSyncSetIDInsensitively(input string) (*SyncSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(SyncSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SyncSetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.OpenShiftClusterName, ok = parsed.Parsed["openShiftClusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'openShiftClusterName' was not found in the resource id %q", input)
	}

	if id.SyncSetName, ok = parsed.Parsed["syncSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'syncSetName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateSyncSetID checks that 'input' can be parsed as a Sync Set ID
func ValidateSyncSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSyncSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sync Set ID
func (id SyncSetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RedHatOpenShift/openShiftClusters/%s/syncSet/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.OpenShiftClusterName, id.SyncSetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sync Set ID
func (id SyncSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticOpenShiftClusters", "openShiftClusters", "openShiftClusters"),
		resourceids.UserSpecifiedSegment("openShiftClusterName", "openShiftClusterValue"),
		resourceids.StaticSegment("staticSyncSet", "syncSet", "syncSet"),
		resourceids.UserSpecifiedSegment("syncSetName", "syncSetValue"),
	}
}

// String returns a human-readable description of this Sync Set ID
func (id SyncSetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Open Shift Cluster Name: %q", id.OpenShiftClusterName),
		fmt.Sprintf("Sync Set Name: %q", id.SyncSetName),
	}
	return fmt.Sprintf("Sync Set (%s)", strings.Join(components, "\n"))
}
