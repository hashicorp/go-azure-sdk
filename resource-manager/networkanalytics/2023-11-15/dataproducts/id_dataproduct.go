package dataproducts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DataProductId{})
}

var _ resourceids.ResourceId = &DataProductId{}

// DataProductId is a struct representing the Resource ID for a Data Product
type DataProductId struct {
	SubscriptionId    string
	ResourceGroupName string
	DataProductName   string
}

// NewDataProductID returns a new DataProductId struct
func NewDataProductID(subscriptionId string, resourceGroupName string, dataProductName string) DataProductId {
	return DataProductId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		DataProductName:   dataProductName,
	}
}

// ParseDataProductID parses 'input' into a DataProductId
func ParseDataProductID(input string) (*DataProductId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataProductId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataProductId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDataProductIDInsensitively parses 'input' case-insensitively into a DataProductId
// note: this method should only be used for API response data and not user input
func ParseDataProductIDInsensitively(input string) (*DataProductId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataProductId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataProductId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DataProductId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DataProductName, ok = input.Parsed["dataProductName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataProductName", input)
	}

	return nil
}

// ValidateDataProductID checks that 'input' can be parsed as a Data Product ID
func ValidateDataProductID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataProductID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data Product ID
func (id DataProductId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkAnalytics/dataProducts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataProductName)
}

// Segments returns a slice of Resource ID Segments which comprise this Data Product ID
func (id DataProductId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkAnalytics", "Microsoft.NetworkAnalytics", "Microsoft.NetworkAnalytics"),
		resourceids.StaticSegment("staticDataProducts", "dataProducts", "dataProducts"),
		resourceids.UserSpecifiedSegment("dataProductName", "dataProductValue"),
	}
}

// String returns a human-readable description of this Data Product ID
func (id DataProductId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Product Name: %q", id.DataProductName),
	}
	return fmt.Sprintf("Data Product (%s)", strings.Join(components, "\n"))
}
