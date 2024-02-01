package restorabledroppedsqlpools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RestorableDroppedSqlPoolId{}

// RestorableDroppedSqlPoolId is a struct representing the Resource ID for a Restorable Dropped Sql Pool
type RestorableDroppedSqlPoolId struct {
	SubscriptionId             string
	ResourceGroupName          string
	WorkspaceName              string
	RestorableDroppedSqlPoolId string
}

// NewRestorableDroppedSqlPoolID returns a new RestorableDroppedSqlPoolId struct
func NewRestorableDroppedSqlPoolID(subscriptionId string, resourceGroupName string, workspaceName string, restorableDroppedSqlPoolId string) RestorableDroppedSqlPoolId {
	return RestorableDroppedSqlPoolId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		WorkspaceName:              workspaceName,
		RestorableDroppedSqlPoolId: restorableDroppedSqlPoolId,
	}
}

// ParseRestorableDroppedSqlPoolID parses 'input' into a RestorableDroppedSqlPoolId
func ParseRestorableDroppedSqlPoolID(input string) (*RestorableDroppedSqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RestorableDroppedSqlPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RestorableDroppedSqlPoolId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRestorableDroppedSqlPoolIDInsensitively parses 'input' case-insensitively into a RestorableDroppedSqlPoolId
// note: this method should only be used for API response data and not user input
func ParseRestorableDroppedSqlPoolIDInsensitively(input string) (*RestorableDroppedSqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RestorableDroppedSqlPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RestorableDroppedSqlPoolId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RestorableDroppedSqlPoolId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RestorableDroppedSqlPoolId, ok = input.Parsed["restorableDroppedSqlPoolId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "restorableDroppedSqlPoolId", input)
	}

	return nil
}

// ValidateRestorableDroppedSqlPoolID checks that 'input' can be parsed as a Restorable Dropped Sql Pool ID
func ValidateRestorableDroppedSqlPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRestorableDroppedSqlPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Restorable Dropped Sql Pool ID
func (id RestorableDroppedSqlPoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/restorableDroppedSqlPools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.RestorableDroppedSqlPoolId)
}

// Segments returns a slice of Resource ID Segments which comprise this Restorable Dropped Sql Pool ID
func (id RestorableDroppedSqlPoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticRestorableDroppedSqlPools", "restorableDroppedSqlPools", "restorableDroppedSqlPools"),
		resourceids.UserSpecifiedSegment("restorableDroppedSqlPoolId", "restorableDroppedSqlPoolIdValue"),
	}
}

// String returns a human-readable description of this Restorable Dropped Sql Pool ID
func (id RestorableDroppedSqlPoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Restorable Dropped Sql Pool: %q", id.RestorableDroppedSqlPoolId),
	}
	return fmt.Sprintf("Restorable Dropped Sql Pool (%s)", strings.Join(components, "\n"))
}
