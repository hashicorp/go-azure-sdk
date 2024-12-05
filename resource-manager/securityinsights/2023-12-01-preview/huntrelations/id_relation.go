package huntrelations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RelationId{})
}

var _ resourceids.ResourceId = &RelationId{}

// RelationId is a struct representing the Resource ID for a Relation
type RelationId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	HuntId            string
	HuntRelationId    string
}

// NewRelationID returns a new RelationId struct
func NewRelationID(subscriptionId string, resourceGroupName string, workspaceName string, huntId string, huntRelationId string) RelationId {
	return RelationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		HuntId:            huntId,
		HuntRelationId:    huntRelationId,
	}
}

// ParseRelationID parses 'input' into a RelationId
func ParseRelationID(input string) (*RelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RelationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRelationIDInsensitively parses 'input' case-insensitively into a RelationId
// note: this method should only be used for API response data and not user input
func ParseRelationIDInsensitively(input string) (*RelationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RelationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RelationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RelationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.HuntId, ok = input.Parsed["huntId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "huntId", input)
	}

	if id.HuntRelationId, ok = input.Parsed["huntRelationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "huntRelationId", input)
	}

	return nil
}

// ValidateRelationID checks that 'input' can be parsed as a Relation ID
func ValidateRelationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRelationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Relation ID
func (id RelationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/hunts/%s/relations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.HuntId, id.HuntRelationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Relation ID
func (id RelationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticHunts", "hunts", "hunts"),
		resourceids.UserSpecifiedSegment("huntId", "huntId"),
		resourceids.StaticSegment("staticRelations", "relations", "relations"),
		resourceids.UserSpecifiedSegment("huntRelationId", "huntRelationId"),
	}
}

// String returns a human-readable description of this Relation ID
func (id RelationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Hunt: %q", id.HuntId),
		fmt.Sprintf("Hunt Relation: %q", id.HuntRelationId),
	}
	return fmt.Sprintf("Relation (%s)", strings.Join(components, "\n"))
}
