package maintenanceevents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MaintenanceEventId{})
}

var _ resourceids.ResourceId = &MaintenanceEventId{}

// MaintenanceEventId is a struct representing the Resource ID for a Maintenance Event
type MaintenanceEventId struct {
	SubscriptionId     string
	ResourceGroupName  string
	FlexibleServerName string
	MaintenanceEventId string
}

// NewMaintenanceEventID returns a new MaintenanceEventId struct
func NewMaintenanceEventID(subscriptionId string, resourceGroupName string, flexibleServerName string, maintenanceEventId string) MaintenanceEventId {
	return MaintenanceEventId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		FlexibleServerName: flexibleServerName,
		MaintenanceEventId: maintenanceEventId,
	}
}

// ParseMaintenanceEventID parses 'input' into a MaintenanceEventId
func ParseMaintenanceEventID(input string) (*MaintenanceEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MaintenanceEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MaintenanceEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMaintenanceEventIDInsensitively parses 'input' case-insensitively into a MaintenanceEventId
// note: this method should only be used for API response data and not user input
func ParseMaintenanceEventIDInsensitively(input string) (*MaintenanceEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MaintenanceEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MaintenanceEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MaintenanceEventId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.MaintenanceEventId, ok = input.Parsed["maintenanceEventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "maintenanceEventId", input)
	}

	return nil
}

// ValidateMaintenanceEventID checks that 'input' can be parsed as a Maintenance Event ID
func ValidateMaintenanceEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMaintenanceEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Maintenance Event ID
func (id MaintenanceEventId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/flexibleServers/%s/maintenanceEvents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FlexibleServerName, id.MaintenanceEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Maintenance Event ID
func (id MaintenanceEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticFlexibleServers", "flexibleServers", "flexibleServers"),
		resourceids.UserSpecifiedSegment("flexibleServerName", "flexibleServerName"),
		resourceids.StaticSegment("staticMaintenanceEvents", "maintenanceEvents", "maintenanceEvents"),
		resourceids.UserSpecifiedSegment("maintenanceEventId", "maintenanceEventId"),
	}
}

// String returns a human-readable description of this Maintenance Event ID
func (id MaintenanceEventId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Flexible Server Name: %q", id.FlexibleServerName),
		fmt.Sprintf("Maintenance Event: %q", id.MaintenanceEventId),
	}
	return fmt.Sprintf("Maintenance Event (%s)", strings.Join(components, "\n"))
}
