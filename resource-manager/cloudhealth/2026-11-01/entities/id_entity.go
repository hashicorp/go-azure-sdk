package entities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&EntityId{})
}

var _ resourceids.ResourceId = &EntityId{}

// EntityId is a struct representing the Resource ID for a Entity
type EntityId struct {
	SubscriptionId    string
	ResourceGroupName string
	HealthmodelName   string
	EntityName        string
}

// NewEntityID returns a new EntityId struct
func NewEntityID(subscriptionId string, resourceGroupName string, healthmodelName string, entityName string) EntityId {
	return EntityId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		HealthmodelName:   healthmodelName,
		EntityName:        entityName,
	}
}

// ParseEntityID parses 'input' into a EntityId
func ParseEntityID(input string) (*EntityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EntityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EntityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEntityIDInsensitively parses 'input' case-insensitively into a EntityId
// note: this method should only be used for API response data and not user input
func ParseEntityIDInsensitively(input string) (*EntityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EntityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EntityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EntityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HealthmodelName, ok = input.Parsed["healthmodelName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "healthmodelName", input)
	}

	if id.EntityName, ok = input.Parsed["entityName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "entityName", input)
	}

	return nil
}

// ValidateEntityID checks that 'input' can be parsed as a Entity ID
func ValidateEntityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEntityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Entity ID
func (id EntityId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CloudHealth/healthmodels/%s/entities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HealthmodelName, id.EntityName)
}

// Segments returns a slice of Resource ID Segments which comprise this Entity ID
func (id EntityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCloudHealth", "Microsoft.CloudHealth", "Microsoft.CloudHealth"),
		resourceids.StaticSegment("staticHealthmodels", "healthmodels", "healthmodels"),
		resourceids.UserSpecifiedSegment("healthmodelName", "healthmodelName"),
		resourceids.StaticSegment("staticEntities", "entities", "entities"),
		resourceids.UserSpecifiedSegment("entityName", "entityName"),
	}
}

// String returns a human-readable description of this Entity ID
func (id EntityId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Healthmodel Name: %q", id.HealthmodelName),
		fmt.Sprintf("Entity Name: %q", id.EntityName),
	}
	return fmt.Sprintf("Entity (%s)", strings.Join(components, "\n"))
}
