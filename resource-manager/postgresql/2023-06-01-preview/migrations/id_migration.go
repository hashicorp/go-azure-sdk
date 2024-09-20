package migrations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MigrationId{})
}

var _ resourceids.ResourceId = &MigrationId{}

// MigrationId is a struct representing the Resource ID for a Migration
type MigrationId struct {
	SubscriptionId     string
	ResourceGroupName  string
	FlexibleServerName string
	MigrationName      string
}

// NewMigrationID returns a new MigrationId struct
func NewMigrationID(subscriptionId string, resourceGroupName string, flexibleServerName string, migrationName string) MigrationId {
	return MigrationId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		FlexibleServerName: flexibleServerName,
		MigrationName:      migrationName,
	}
}

// ParseMigrationID parses 'input' into a MigrationId
func ParseMigrationID(input string) (*MigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MigrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMigrationIDInsensitively parses 'input' case-insensitively into a MigrationId
// note: this method should only be used for API response data and not user input
func ParseMigrationIDInsensitively(input string) (*MigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MigrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MigrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.FlexibleServerName, ok = input.Parsed["flexibleServerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "flexibleServerName", input)
	}

	if id.MigrationName, ok = input.Parsed["migrationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "migrationName", input)
	}

	return nil
}

// ValidateMigrationID checks that 'input' can be parsed as a Migration ID
func ValidateMigrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMigrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Migration ID
func (id MigrationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/flexibleServers/%s/migrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FlexibleServerName, id.MigrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Migration ID
func (id MigrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticFlexibleServers", "flexibleServers", "flexibleServers"),
		resourceids.UserSpecifiedSegment("flexibleServerName", "targetDbServerName"),
		resourceids.StaticSegment("staticMigrations", "migrations", "migrations"),
		resourceids.UserSpecifiedSegment("migrationName", "migrationName"),
	}
}

// String returns a human-readable description of this Migration ID
func (id MigrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Flexible Server Name: %q", id.FlexibleServerName),
		fmt.Sprintf("Migration Name: %q", id.MigrationName),
	}
	return fmt.Sprintf("Migration (%s)", strings.Join(components, "\n"))
}
