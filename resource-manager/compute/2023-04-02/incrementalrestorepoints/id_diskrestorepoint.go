package incrementalrestorepoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DiskRestorePointId{})
}

var _ resourceids.ResourceId = &DiskRestorePointId{}

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
	parser := resourceids.NewParserFromResourceIdType(&DiskRestorePointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiskRestorePointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiskRestorePointIDInsensitively parses 'input' case-insensitively into a DiskRestorePointId
// note: this method should only be used for API response data and not user input
func ParseDiskRestorePointIDInsensitively(input string) (*DiskRestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiskRestorePointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiskRestorePointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiskRestorePointId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.RestorePointCollectionName, ok = input.Parsed["restorePointCollectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "restorePointCollectionName", input)
	}

	if id.RestorePointName, ok = input.Parsed["restorePointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "restorePointName", input)
	}

	if id.DiskRestorePointName, ok = input.Parsed["diskRestorePointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "diskRestorePointName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("restorePointCollectionName", "restorePointCollectionName"),
		resourceids.StaticSegment("staticRestorePoints", "restorePoints", "restorePoints"),
		resourceids.UserSpecifiedSegment("restorePointName", "vmRestorePointName"),
		resourceids.StaticSegment("staticDiskRestorePoints", "diskRestorePoints", "diskRestorePoints"),
		resourceids.UserSpecifiedSegment("diskRestorePointName", "diskRestorePointName"),
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
