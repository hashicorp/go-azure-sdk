// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incrementalrestorepoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DiskRestorePointId{}

// DiskRestorePointId is a struct representing the Resource ID for a Disk Restore Point
type DiskRestorePointId struct {
	SubscriptionId             string
	ResourceGroupName          string
	RestorePointCollectionName string
	RestorePointName           string
	DiskRestorePointName       string
}

// NewDiskRestorePointID returns a new DiskRestorePointId struct
func NewDiskRestorePointID(subscriptionId string, resourceGroupName string, restorePointCollectionName string, restorePointName string, diskRestorePointName string) DiskRestorePointId {
	return DiskRestorePointId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		RestorePointCollectionName: restorePointCollectionName,
		RestorePointName:           restorePointName,
		DiskRestorePointName:       diskRestorePointName,
	}
}

// ParseDiskRestorePointID parses 'input' into a DiskRestorePointId
func ParseDiskRestorePointID(input string) (*DiskRestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(DiskRestorePointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DiskRestorePointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RestorePointCollectionName, ok = parsed.Parsed["restorePointCollectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointCollectionName' was not found in the resource id %q", input)
	}

	if id.RestorePointName, ok = parsed.Parsed["restorePointName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointName' was not found in the resource id %q", input)
	}

	if id.DiskRestorePointName, ok = parsed.Parsed["diskRestorePointName"]; !ok {
		return nil, fmt.Errorf("the segment 'diskRestorePointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDiskRestorePointIDInsensitively parses 'input' case-insensitively into a DiskRestorePointId
// note: this method should only be used for API response data and not user input
func ParseDiskRestorePointIDInsensitively(input string) (*DiskRestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(DiskRestorePointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DiskRestorePointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RestorePointCollectionName, ok = parsed.Parsed["restorePointCollectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointCollectionName' was not found in the resource id %q", input)
	}

	if id.RestorePointName, ok = parsed.Parsed["restorePointName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointName' was not found in the resource id %q", input)
	}

	if id.DiskRestorePointName, ok = parsed.Parsed["diskRestorePointName"]; !ok {
		return nil, fmt.Errorf("the segment 'diskRestorePointName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDiskRestorePointID checks that 'input' can be parsed as a Disk Restore Point ID
func ValidateDiskRestorePointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiskRestorePointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Disk Restore Point ID
func (id DiskRestorePointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/restorePointCollections/%s/restorePoints/%s/diskRestorePoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RestorePointCollectionName, id.RestorePointName, id.DiskRestorePointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Disk Restore Point ID
func (id DiskRestorePointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticRestorePointCollections", "restorePointCollections", "restorePointCollections"),
		resourceids.UserSpecifiedSegment("restorePointCollectionName", "restorePointCollectionValue"),
		resourceids.StaticSegment("staticRestorePoints", "restorePoints", "restorePoints"),
		resourceids.UserSpecifiedSegment("restorePointName", "restorePointValue"),
		resourceids.StaticSegment("staticDiskRestorePoints", "diskRestorePoints", "diskRestorePoints"),
		resourceids.UserSpecifiedSegment("diskRestorePointName", "diskRestorePointValue"),
	}
}

// String returns a human-readable description of this Disk Restore Point ID
func (id DiskRestorePointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Restore Point Collection Name: %q", id.RestorePointCollectionName),
		fmt.Sprintf("Restore Point Name: %q", id.RestorePointName),
		fmt.Sprintf("Disk Restore Point Name: %q", id.DiskRestorePointName),
	}
	return fmt.Sprintf("Disk Restore Point (%s)", strings.Join(components, "\n"))
}
