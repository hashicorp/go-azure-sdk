package schemaregistries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SchemaRegistryId{})
}

var _ resourceids.ResourceId = &SchemaRegistryId{}

// SchemaRegistryId is a struct representing the Resource ID for a Schema Registry
type SchemaRegistryId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SchemaRegistryName string
}

// NewSchemaRegistryID returns a new SchemaRegistryId struct
func NewSchemaRegistryID(subscriptionId string, resourceGroupName string, schemaRegistryName string) SchemaRegistryId {
	return SchemaRegistryId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SchemaRegistryName: schemaRegistryName,
	}
}

// ParseSchemaRegistryID parses 'input' into a SchemaRegistryId
func ParseSchemaRegistryID(input string) (*SchemaRegistryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaRegistryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaRegistryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSchemaRegistryIDInsensitively parses 'input' case-insensitively into a SchemaRegistryId
// note: this method should only be used for API response data and not user input
func ParseSchemaRegistryIDInsensitively(input string) (*SchemaRegistryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaRegistryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaRegistryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SchemaRegistryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SchemaRegistryName, ok = input.Parsed["schemaRegistryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schemaRegistryName", input)
	}

	return nil
}

// ValidateSchemaRegistryID checks that 'input' can be parsed as a Schema Registry ID
func ValidateSchemaRegistryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSchemaRegistryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Schema Registry ID
func (id SchemaRegistryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DeviceRegistry/schemaRegistries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SchemaRegistryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Schema Registry ID
func (id SchemaRegistryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticSchemaRegistries", "schemaRegistries", "schemaRegistries"),
		resourceids.UserSpecifiedSegment("schemaRegistryName", "schemaRegistryName"),
	}
}

// String returns a human-readable description of this Schema Registry ID
func (id SchemaRegistryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Schema Registry Name: %q", id.SchemaRegistryName),
	}
	return fmt.Sprintf("Schema Registry (%s)", strings.Join(components, "\n"))
}
