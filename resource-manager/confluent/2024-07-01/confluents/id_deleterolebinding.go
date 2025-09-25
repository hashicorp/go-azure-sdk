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
	recaser.RegisterResourceId(&DeleteRoleBindingId{})
}

var _ resourceids.ResourceId = &DeleteRoleBindingId{}

// DeleteRoleBindingId is a struct representing the Resource ID for a Delete Role Binding
type DeleteRoleBindingId struct {
	SubscriptionId    string
	ResourceGroupName string
	OrganizationName  string
	RoleBindingId     string
}

// NewDeleteRoleBindingID returns a new DeleteRoleBindingId struct
func NewDeleteRoleBindingID(subscriptionId string, resourceGroupName string, organizationName string, roleBindingId string) DeleteRoleBindingId {
	return DeleteRoleBindingId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		OrganizationName:  organizationName,
		RoleBindingId:     roleBindingId,
	}
}

// ParseDeleteRoleBindingID parses 'input' into a DeleteRoleBindingId
func ParseDeleteRoleBindingID(input string) (*DeleteRoleBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeleteRoleBindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeleteRoleBindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeleteRoleBindingIDInsensitively parses 'input' case-insensitively into a DeleteRoleBindingId
// note: this method should only be used for API response data and not user input
func ParseDeleteRoleBindingIDInsensitively(input string) (*DeleteRoleBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeleteRoleBindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeleteRoleBindingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeleteRoleBindingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RoleBindingId, ok = input.Parsed["roleBindingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleBindingId", input)
	}

	return nil
}

// ValidateDeleteRoleBindingID checks that 'input' can be parsed as a Delete Role Binding ID
func ValidateDeleteRoleBindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeleteRoleBindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Delete Role Binding ID
func (id DeleteRoleBindingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Confluent/organizations/%s/access/default/deleteRoleBinding/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.OrganizationName, id.RoleBindingId)
}

// Segments returns a slice of Resource ID Segments which comprise this Delete Role Binding ID
func (id DeleteRoleBindingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConfluent", "Microsoft.Confluent", "Microsoft.Confluent"),
		resourceids.StaticSegment("staticOrganizations", "organizations", "organizations"),
		resourceids.UserSpecifiedSegment("organizationName", "organizationName"),
		resourceids.StaticSegment("staticAccess", "access", "access"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticDeleteRoleBinding", "deleteRoleBinding", "deleteRoleBinding"),
		resourceids.UserSpecifiedSegment("roleBindingId", "roleBindingId"),
	}
}

// String returns a human-readable description of this Delete Role Binding ID
func (id DeleteRoleBindingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Organization Name: %q", id.OrganizationName),
		fmt.Sprintf("Role Binding: %q", id.RoleBindingId),
	}
	return fmt.Sprintf("Delete Role Binding (%s)", strings.Join(components, "\n"))
}
