package subvolumes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SubVolumeId{}

// SubVolumeId is a struct representing the Resource ID for a Sub Volume
type SubVolumeId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	CapacityPoolName  string
	VolumeName        string
	SubVolumeName     string
}

// NewSubVolumeID returns a new SubVolumeId struct
func NewSubVolumeID(subscriptionId string, resourceGroupName string, netAppAccountName string, capacityPoolName string, volumeName string, subVolumeName string) SubVolumeId {
	return SubVolumeId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		CapacityPoolName:  capacityPoolName,
		VolumeName:        volumeName,
		SubVolumeName:     subVolumeName,
	}
}

// ParseSubVolumeID parses 'input' into a SubVolumeId
func ParseSubVolumeID(input string) (*SubVolumeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubVolumeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubVolumeId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSubVolumeIDInsensitively parses 'input' case-insensitively into a SubVolumeId
// note: this method should only be used for API response data and not user input
func ParseSubVolumeIDInsensitively(input string) (*SubVolumeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SubVolumeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SubVolumeId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SubVolumeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetAppAccountName, ok = input.Parsed["netAppAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", input)
	}

	if id.CapacityPoolName, ok = input.Parsed["capacityPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", input)
	}

	if id.VolumeName, ok = input.Parsed["volumeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "volumeName", input)
	}

	if id.SubVolumeName, ok = input.Parsed["subVolumeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subVolumeName", input)
	}

	return nil
}

// ValidateSubVolumeID checks that 'input' can be parsed as a Sub Volume ID
func ValidateSubVolumeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSubVolumeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sub Volume ID
func (id SubVolumeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/capacityPools/%s/volumes/%s/subVolumes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.CapacityPoolName, id.VolumeName, id.SubVolumeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sub Volume ID
func (id SubVolumeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountValue"),
		resourceids.StaticSegment("staticCapacityPools", "capacityPools", "capacityPools"),
		resourceids.UserSpecifiedSegment("capacityPoolName", "capacityPoolValue"),
		resourceids.StaticSegment("staticVolumes", "volumes", "volumes"),
		resourceids.UserSpecifiedSegment("volumeName", "volumeValue"),
		resourceids.StaticSegment("staticSubVolumes", "subVolumes", "subVolumes"),
		resourceids.UserSpecifiedSegment("subVolumeName", "subVolumeValue"),
	}
}

// String returns a human-readable description of this Sub Volume ID
func (id SubVolumeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Capacity Pool Name: %q", id.CapacityPoolName),
		fmt.Sprintf("Volume Name: %q", id.VolumeName),
		fmt.Sprintf("Sub Volume Name: %q", id.SubVolumeName),
	}
	return fmt.Sprintf("Sub Volume (%s)", strings.Join(components, "\n"))
}
