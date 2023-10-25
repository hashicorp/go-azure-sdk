package deletedservers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DeletedServerId{}

// DeletedServerId is a struct representing the Resource ID for a Deleted Server
type DeletedServerId struct {
	SubscriptionId    string
	LocationName      string
	DeletedServerName string
}

// NewDeletedServerID returns a new DeletedServerId struct
func NewDeletedServerID(subscriptionId string, locationName string, deletedServerName string) DeletedServerId {
	return DeletedServerId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		DeletedServerName: deletedServerName,
	}
}

// ParseDeletedServerID parses 'input' into a DeletedServerId
func ParseDeletedServerID(input string) (*DeletedServerId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeletedServerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeletedServerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.DeletedServerName, ok = parsed.Parsed["deletedServerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deletedServerName", *parsed)
	}

	return &id, nil
}

// ParseDeletedServerIDInsensitively parses 'input' case-insensitively into a DeletedServerId
// note: this method should only be used for API response data and not user input
func ParseDeletedServerIDInsensitively(input string) (*DeletedServerId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeletedServerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeletedServerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.DeletedServerName, ok = parsed.Parsed["deletedServerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deletedServerName", *parsed)
	}

	return &id, nil
}

// ValidateDeletedServerID checks that 'input' can be parsed as a Deleted Server ID
func ValidateDeletedServerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeletedServerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deleted Server ID
func (id DeletedServerId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/deletedServers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.DeletedServerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Deleted Server ID
func (id DeletedServerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticDeletedServers", "deletedServers", "deletedServers"),
		resourceids.UserSpecifiedSegment("deletedServerName", "deletedServerValue"),
	}
}

// String returns a human-readable description of this Deleted Server ID
func (id DeletedServerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Deleted Server Name: %q", id.DeletedServerName),
	}
	return fmt.Sprintf("Deleted Server (%s)", strings.Join(components, "\n"))
}
