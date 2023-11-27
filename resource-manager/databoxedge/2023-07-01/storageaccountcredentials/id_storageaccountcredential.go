package storageaccountcredentials

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = StorageAccountCredentialId{}

// StorageAccountCredentialId is a struct representing the Resource ID for a Storage Account Credential
type StorageAccountCredentialId struct {
	SubscriptionId               string
	ResourceGroupName            string
	DataBoxEdgeDeviceName        string
	StorageAccountCredentialName string
}

// NewStorageAccountCredentialID returns a new StorageAccountCredentialId struct
func NewStorageAccountCredentialID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, storageAccountCredentialName string) StorageAccountCredentialId {
	return StorageAccountCredentialId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		DataBoxEdgeDeviceName:        dataBoxEdgeDeviceName,
		StorageAccountCredentialName: storageAccountCredentialName,
	}
}

// ParseStorageAccountCredentialID parses 'input' into a StorageAccountCredentialId
func ParseStorageAccountCredentialID(input string) (*StorageAccountCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(StorageAccountCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StorageAccountCredentialId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStorageAccountCredentialIDInsensitively parses 'input' case-insensitively into a StorageAccountCredentialId
// note: this method should only be used for API response data and not user input
func ParseStorageAccountCredentialIDInsensitively(input string) (*StorageAccountCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(StorageAccountCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StorageAccountCredentialId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StorageAccountCredentialId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DataBoxEdgeDeviceName, ok = input.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataBoxEdgeDeviceName", input)
	}

	if id.StorageAccountCredentialName, ok = input.Parsed["storageAccountCredentialName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "storageAccountCredentialName", input)
	}

	return nil
}

// ValidateStorageAccountCredentialID checks that 'input' can be parsed as a Storage Account Credential ID
func ValidateStorageAccountCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStorageAccountCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Storage Account Credential ID
func (id StorageAccountCredentialId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/storageAccountCredentials/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.StorageAccountCredentialName)
}

// Segments returns a slice of Resource ID Segments which comprise this Storage Account Credential ID
func (id StorageAccountCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticStorageAccountCredentials", "storageAccountCredentials", "storageAccountCredentials"),
		resourceids.UserSpecifiedSegment("storageAccountCredentialName", "storageAccountCredentialValue"),
	}
}

// String returns a human-readable description of this Storage Account Credential ID
func (id StorageAccountCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Storage Account Credential Name: %q", id.StorageAccountCredentialName),
	}
	return fmt.Sprintf("Storage Account Credential (%s)", strings.Join(components, "\n"))
}
