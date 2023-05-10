package datasets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DatasetId{}

// DatasetId is a struct representing the Resource ID for a Dataset
type DatasetId struct {
	SubscriptionId    string
	ResourceGroupName string
	FactoryName       string
	DatasetName       string
}

// NewDatasetID returns a new DatasetId struct
func NewDatasetID(subscriptionId string, resourceGroupName string, factoryName string, datasetName string) DatasetId {
	return DatasetId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FactoryName:       factoryName,
		DatasetName:       datasetName,
	}
}

// ParseDatasetID parses 'input' into a DatasetId
func ParseDatasetID(input string) (*DatasetId, error) {
	parser := resourceids.NewParserFromResourceIdType(DatasetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DatasetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "factoryName", *parsed)
	}

	if id.DatasetName, ok = parsed.Parsed["datasetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "datasetName", *parsed)
	}

	return &id, nil
}

// ParseDatasetIDInsensitively parses 'input' case-insensitively into a DatasetId
// note: this method should only be used for API response data and not user input
func ParseDatasetIDInsensitively(input string) (*DatasetId, error) {
	parser := resourceids.NewParserFromResourceIdType(DatasetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DatasetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "factoryName", *parsed)
	}

	if id.DatasetName, ok = parsed.Parsed["datasetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "datasetName", *parsed)
	}

	return &id, nil
}

// ValidateDatasetID checks that 'input' can be parsed as a Dataset ID
func ValidateDatasetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDatasetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dataset ID
func (id DatasetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/datasets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.DatasetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dataset ID
func (id DatasetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticDatasets", "datasets", "datasets"),
		resourceids.UserSpecifiedSegment("datasetName", "datasetValue"),
	}
}

// String returns a human-readable description of this Dataset ID
func (id DatasetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Dataset Name: %q", id.DatasetName),
	}
	return fmt.Sprintf("Dataset (%s)", strings.Join(components, "\n"))
}
