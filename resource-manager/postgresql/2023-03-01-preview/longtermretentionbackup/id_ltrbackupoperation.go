package longtermretentionbackup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LtrBackupOperationId{}

// LtrBackupOperationId is a struct representing the Resource ID for a Ltr Backup Operation
type LtrBackupOperationId struct {
	SubscriptionId         string
	ResourceGroupName      string
	FlexibleServerName     string
	LtrBackupOperationName string
}

// NewLtrBackupOperationID returns a new LtrBackupOperationId struct
func NewLtrBackupOperationID(subscriptionId string, resourceGroupName string, flexibleServerName string, ltrBackupOperationName string) LtrBackupOperationId {
	return LtrBackupOperationId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		FlexibleServerName:     flexibleServerName,
		LtrBackupOperationName: ltrBackupOperationName,
	}
}

// ParseLtrBackupOperationID parses 'input' into a LtrBackupOperationId
func ParseLtrBackupOperationID(input string) (*LtrBackupOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(LtrBackupOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LtrBackupOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FlexibleServerName, ok = parsed.Parsed["flexibleServerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "flexibleServerName", *parsed)
	}

	if id.LtrBackupOperationName, ok = parsed.Parsed["ltrBackupOperationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "ltrBackupOperationName", *parsed)
	}

	return &id, nil
}

// ParseLtrBackupOperationIDInsensitively parses 'input' case-insensitively into a LtrBackupOperationId
// note: this method should only be used for API response data and not user input
func ParseLtrBackupOperationIDInsensitively(input string) (*LtrBackupOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(LtrBackupOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LtrBackupOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FlexibleServerName, ok = parsed.Parsed["flexibleServerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "flexibleServerName", *parsed)
	}

	if id.LtrBackupOperationName, ok = parsed.Parsed["ltrBackupOperationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "ltrBackupOperationName", *parsed)
	}

	return &id, nil
}

// ValidateLtrBackupOperationID checks that 'input' can be parsed as a Ltr Backup Operation ID
func ValidateLtrBackupOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLtrBackupOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Ltr Backup Operation ID
func (id LtrBackupOperationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/flexibleServers/%s/ltrBackupOperations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FlexibleServerName, id.LtrBackupOperationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Ltr Backup Operation ID
func (id LtrBackupOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticFlexibleServers", "flexibleServers", "flexibleServers"),
		resourceids.UserSpecifiedSegment("flexibleServerName", "flexibleServerValue"),
		resourceids.StaticSegment("staticLtrBackupOperations", "ltrBackupOperations", "ltrBackupOperations"),
		resourceids.UserSpecifiedSegment("ltrBackupOperationName", "ltrBackupOperationValue"),
	}
}

// String returns a human-readable description of this Ltr Backup Operation ID
func (id LtrBackupOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Flexible Server Name: %q", id.FlexibleServerName),
		fmt.Sprintf("Ltr Backup Operation Name: %q", id.LtrBackupOperationName),
	}
	return fmt.Sprintf("Ltr Backup Operation (%s)", strings.Join(components, "\n"))
}
