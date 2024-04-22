package global

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DeletedSiteId{})
}

var _ resourceids.ResourceId = &DeletedSiteId{}

// DeletedSiteId is a struct representing the Resource ID for a Deleted Site
type DeletedSiteId struct {
	SubscriptionId string
	DeletedSiteId  string
}

// NewDeletedSiteID returns a new DeletedSiteId struct
func NewDeletedSiteID(subscriptionId string, deletedSiteId string) DeletedSiteId {
	return DeletedSiteId{
		SubscriptionId: subscriptionId,
		DeletedSiteId:  deletedSiteId,
	}
}

// ParseDeletedSiteID parses 'input' into a DeletedSiteId
func ParseDeletedSiteID(input string) (*DeletedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeletedSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeletedSiteId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeletedSiteIDInsensitively parses 'input' case-insensitively into a DeletedSiteId
// note: this method should only be used for API response data and not user input
func ParseDeletedSiteIDInsensitively(input string) (*DeletedSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeletedSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeletedSiteId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeletedSiteId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.DeletedSiteId, ok = input.Parsed["deletedSiteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deletedSiteId", input)
	}

	return nil
}

// ValidateDeletedSiteID checks that 'input' can be parsed as a Deleted Site ID
func ValidateDeletedSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeletedSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deleted Site ID
func (id DeletedSiteId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Web/deletedSites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.DeletedSiteId)
}

// Segments returns a slice of Resource ID Segments which comprise this Deleted Site ID
func (id DeletedSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticDeletedSites", "deletedSites", "deletedSites"),
		resourceids.UserSpecifiedSegment("deletedSiteId", "deletedSiteIdValue"),
	}
}

// String returns a human-readable description of this Deleted Site ID
func (id DeletedSiteId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Deleted Site: %q", id.DeletedSiteId),
	}
	return fmt.Sprintf("Deleted Site (%s)", strings.Join(components, "\n"))
}
