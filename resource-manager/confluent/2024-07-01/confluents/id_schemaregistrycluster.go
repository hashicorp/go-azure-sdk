package confluents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SchemaRegistryClusterId{})
}

var _ resourceids.ResourceId = &SchemaRegistryClusterId{}

// SchemaRegistryClusterId is a struct representing the Resource ID for a Schema Registry Cluster
type SchemaRegistryClusterId struct {
	SubscriptionId    string
	ResourceGroupName string
	OrganizationName  string
	EnvironmentId     string
	ClusterId         string
}

// NewSchemaRegistryClusterID returns a new SchemaRegistryClusterId struct
func NewSchemaRegistryClusterID(subscriptionId string, resourceGroupName string, organizationName string, environmentId string, clusterId string) SchemaRegistryClusterId {
	return SchemaRegistryClusterId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		OrganizationName:  organizationName,
		EnvironmentId:     environmentId,
		ClusterId:         clusterId,
	}
}

// ParseSchemaRegistryClusterID parses 'input' into a SchemaRegistryClusterId
func ParseSchemaRegistryClusterID(input string) (*SchemaRegistryClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaRegistryClusterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaRegistryClusterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSchemaRegistryClusterIDInsensitively parses 'input' case-insensitively into a SchemaRegistryClusterId
// note: this method should only be used for API response data and not user input
func ParseSchemaRegistryClusterIDInsensitively(input string) (*SchemaRegistryClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SchemaRegistryClusterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SchemaRegistryClusterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SchemaRegistryClusterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.OrganizationName, ok = input.Parsed["organizationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "organizationName", input)
	}

	if id.EnvironmentId, ok = input.Parsed["environmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "environmentId", input)
	}

	if id.ClusterId, ok = input.Parsed["clusterId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterId", input)
	}

	return nil
}

// ValidateSchemaRegistryClusterID checks that 'input' can be parsed as a Schema Registry Cluster ID
func ValidateSchemaRegistryClusterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSchemaRegistryClusterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Schema Registry Cluster ID
func (id SchemaRegistryClusterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Confluent/organizations/%s/environments/%s/schemaRegistryClusters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.OrganizationName, id.EnvironmentId, id.ClusterId)
}

// Segments returns a slice of Resource ID Segments which comprise this Schema Registry Cluster ID
func (id SchemaRegistryClusterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConfluent", "Microsoft.Confluent", "Microsoft.Confluent"),
		resourceids.StaticSegment("staticOrganizations", "organizations", "organizations"),
		resourceids.UserSpecifiedSegment("organizationName", "organizationName"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.UserSpecifiedSegment("environmentId", "environmentId"),
		resourceids.StaticSegment("staticSchemaRegistryClusters", "schemaRegistryClusters", "schemaRegistryClusters"),
		resourceids.UserSpecifiedSegment("clusterId", "clusterId"),
	}
}

// String returns a human-readable description of this Schema Registry Cluster ID
func (id SchemaRegistryClusterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Organization Name: %q", id.OrganizationName),
		fmt.Sprintf("Environment: %q", id.EnvironmentId),
		fmt.Sprintf("Cluster: %q", id.ClusterId),
	}
	return fmt.Sprintf("Schema Registry Cluster (%s)", strings.Join(components, "\n"))
}
