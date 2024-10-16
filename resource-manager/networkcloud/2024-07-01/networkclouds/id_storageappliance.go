package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&StorageApplianceId{})
}

var _ resourceids.ResourceId = &StorageApplianceId{}

// StorageApplianceId is a struct representing the Resource ID for a Storage Appliance
type StorageApplianceId struct {
	SubscriptionId       string
	ResourceGroupName    string
	StorageApplianceName string
}

// NewStorageApplianceID returns a new StorageApplianceId struct
func NewStorageApplianceID(subscriptionId string, resourceGroupName string, storageApplianceName string) StorageApplianceId {
	return StorageApplianceId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		StorageApplianceName: storageApplianceName,
	}
}

// ParseStorageApplianceID parses 'input' into a StorageApplianceId
func ParseStorageApplianceID(input string) (*StorageApplianceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StorageApplianceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StorageApplianceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStorageApplianceIDInsensitively parses 'input' case-insensitively into a StorageApplianceId
// note: this method should only be used for API response data and not user input
func ParseStorageApplianceIDInsensitively(input string) (*StorageApplianceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StorageApplianceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StorageApplianceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StorageApplianceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.StorageApplianceName, ok = input.Parsed["storageApplianceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "storageApplianceName", input)
	}

	return nil
}

// ValidateStorageApplianceID checks that 'input' can be parsed as a Storage Appliance ID
func ValidateStorageApplianceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStorageApplianceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Storage Appliance ID
func (id StorageApplianceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/storageAppliances/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StorageApplianceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Storage Appliance ID
func (id StorageApplianceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticStorageAppliances", "storageAppliances", "storageAppliances"),
		resourceids.UserSpecifiedSegment("storageApplianceName", "storageApplianceName"),
	}
}

// String returns a human-readable description of this Storage Appliance ID
func (id StorageApplianceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Storage Appliance Name: %q", id.StorageApplianceName),
	}
	return fmt.Sprintf("Storage Appliance (%s)", strings.Join(components, "\n"))
}
