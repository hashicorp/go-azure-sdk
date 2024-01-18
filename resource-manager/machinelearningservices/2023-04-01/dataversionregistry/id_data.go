package dataversionregistry

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DataId{}

// DataId is a struct representing the Resource ID for a Data
type DataId struct {
	SubscriptionId    string
	ResourceGroupName string
	RegistryName      string
	DataName          string
}

// NewDataID returns a new DataId struct
func NewDataID(subscriptionId string, resourceGroupName string, registryName string, dataName string) DataId {
	return DataId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RegistryName:      registryName,
		DataName:          dataName,
	}
}

// ParseDataID parses 'input' into a DataId
func ParseDataID(input string) (*DataId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDataIDInsensitively parses 'input' case-insensitively into a DataId
// note: this method should only be used for API response data and not user input
func ParseDataIDInsensitively(input string) (*DataId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DataId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.RegistryName, ok = input.Parsed["registryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "registryName", input)
	}

	if id.DataName, ok = input.Parsed["dataName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataName", input)
	}

	return nil
}

// ValidateDataID checks that 'input' can be parsed as a Data ID
func ValidateDataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data ID
func (id DataId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/registries/%s/data/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RegistryName, id.DataName)
}

// Segments returns a slice of Resource ID Segments which comprise this Data ID
func (id DataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticRegistries", "registries", "registries"),
		resourceids.UserSpecifiedSegment("registryName", "registryValue"),
		resourceids.StaticSegment("staticData", "data", "data"),
		resourceids.UserSpecifiedSegment("dataName", "dataValue"),
	}
}

// String returns a human-readable description of this Data ID
func (id DataId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Registry Name: %q", id.RegistryName),
		fmt.Sprintf("Data Name: %q", id.DataName),
	}
	return fmt.Sprintf("Data (%s)", strings.Join(components, "\n"))
}
