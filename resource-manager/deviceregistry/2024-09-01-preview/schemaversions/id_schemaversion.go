package schemaversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SchemaVersionId{})
}

var _ resourceids.ResourceId = &SchemaVersionId{}

// SchemaVersionId is a struct representing the Resource ID for a Schema Version
type SchemaVersionId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SchemaRegistryName string
	SchemaName         string
	SchemaVersionName  string
}

// NewSchemaVersionID returns a new SchemaVersionId struct
func NewSchemaVersionID(subscriptionId string, resourceGroupName string, schemaRegistryName string, schemaName string, schemaVersionName string) SchemaVersionId {
	return SchemaVersionId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SchemaRegistryName: schemaRegistryName,
		SchemaName:         schemaName,
		SchemaVersionName:  schemaVersionName,
	}
}

// ParseSchemaVersionID parses 'input' into a SchemaVersionId
func ParseSchemaVersionID(input string) (*SchemaVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSchemaVersionIDInsensitively parses 'input' case-insensitively into a SchemaVersionId
// note: this method should only be used for API response data and not user input
func ParseSchemaVersionIDInsensitively(input string) (*SchemaVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SchemaVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SchemaName, ok = input.Parsed["schemaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schemaName", input)
	}

	if id.SchemaVersionName, ok = input.Parsed["schemaVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schemaVersionName", input)
	}

	return nil
}

// ValidateSchemaVersionID checks that 'input' can be parsed as a Schema Version ID
func ValidateSchemaVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSchemaVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Schema Version ID
func (id SchemaVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DeviceRegistry/schemaRegistries/%s/schemas/%s/schemaVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SchemaRegistryName, id.SchemaName, id.SchemaVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Schema Version ID
func (id SchemaVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticSchemaRegistries", "schemaRegistries", "schemaRegistries"),
		resourceids.UserSpecifiedSegment("schemaRegistryName", "schemaRegistryName"),
		resourceids.StaticSegment("staticSchemas", "schemas", "schemas"),
		resourceids.UserSpecifiedSegment("schemaName", "schemaName"),
		resourceids.StaticSegment("staticSchemaVersions", "schemaVersions", "schemaVersions"),
		resourceids.UserSpecifiedSegment("schemaVersionName", "schemaVersionName"),
	}
}

// String returns a human-readable description of this Schema Version ID
func (id SchemaVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Schema Registry Name: %q", id.SchemaRegistryName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Schema Version Name: %q", id.SchemaVersionName),
	}
	return fmt.Sprintf("Schema Version (%s)", strings.Join(components, "\n"))
}
