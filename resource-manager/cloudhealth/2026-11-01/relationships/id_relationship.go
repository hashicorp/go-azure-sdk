package relationships

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RelationshipId{})
}

var _ resourceids.ResourceId = &RelationshipId{}

// RelationshipId is a struct representing the Resource ID for a Relationship
type RelationshipId struct {
	SubscriptionId    string
	ResourceGroupName string
	HealthmodelName   string
	RelationshipName  string
}

// NewRelationshipID returns a new RelationshipId struct
func NewRelationshipID(subscriptionId string, resourceGroupName string, healthmodelName string, relationshipName string) RelationshipId {
	return RelationshipId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		HealthmodelName:   healthmodelName,
		RelationshipName:  relationshipName,
	}
}

// ParseRelationshipID parses 'input' into a RelationshipId
func ParseRelationshipID(input string) (*RelationshipId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RelationshipId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RelationshipId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRelationshipIDInsensitively parses 'input' case-insensitively into a RelationshipId
// note: this method should only be used for API response data and not user input
func ParseRelationshipIDInsensitively(input string) (*RelationshipId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RelationshipId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RelationshipId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RelationshipId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RelationshipName, ok = input.Parsed["relationshipName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "relationshipName", input)
	}

	return nil
}

// ValidateRelationshipID checks that 'input' can be parsed as a Relationship ID
func ValidateRelationshipID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRelationshipID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Relationship ID
func (id RelationshipId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CloudHealth/healthmodels/%s/relationships/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HealthmodelName, id.RelationshipName)
}

// Segments returns a slice of Resource ID Segments which comprise this Relationship ID
func (id RelationshipId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCloudHealth", "Microsoft.CloudHealth", "Microsoft.CloudHealth"),
		resourceids.StaticSegment("staticHealthmodels", "healthmodels", "healthmodels"),
		resourceids.UserSpecifiedSegment("healthmodelName", "healthmodelName"),
		resourceids.StaticSegment("staticRelationships", "relationships", "relationships"),
		resourceids.UserSpecifiedSegment("relationshipName", "relationshipName"),
	}
}

// String returns a human-readable description of this Relationship ID
func (id RelationshipId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Healthmodel Name: %q", id.HealthmodelName),
		fmt.Sprintf("Relationship Name: %q", id.RelationshipName),
	}
	return fmt.Sprintf("Relationship (%s)", strings.Join(components, "\n"))
}
