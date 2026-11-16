package majorversionupgradeprecheck

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MajorVersionUpgradePrecheckId{})
}

var _ resourceids.ResourceId = &MajorVersionUpgradePrecheckId{}

// MajorVersionUpgradePrecheckId is a struct representing the Resource ID for a Major Version Upgrade Precheck
type MajorVersionUpgradePrecheckId struct {
	SubscriptionId       string
	ResourceGroupName    string
	FlexibleServerName   string
	PrecheckValidationId string
}

// NewMajorVersionUpgradePrecheckID returns a new MajorVersionUpgradePrecheckId struct
func NewMajorVersionUpgradePrecheckID(subscriptionId string, resourceGroupName string, flexibleServerName string, precheckValidationId string) MajorVersionUpgradePrecheckId {
	return MajorVersionUpgradePrecheckId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		FlexibleServerName:   flexibleServerName,
		PrecheckValidationId: precheckValidationId,
	}
}

// ParseMajorVersionUpgradePrecheckID parses 'input' into a MajorVersionUpgradePrecheckId
func ParseMajorVersionUpgradePrecheckID(input string) (*MajorVersionUpgradePrecheckId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MajorVersionUpgradePrecheckId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MajorVersionUpgradePrecheckId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMajorVersionUpgradePrecheckIDInsensitively parses 'input' case-insensitively into a MajorVersionUpgradePrecheckId
// note: this method should only be used for API response data and not user input
func ParseMajorVersionUpgradePrecheckIDInsensitively(input string) (*MajorVersionUpgradePrecheckId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MajorVersionUpgradePrecheckId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MajorVersionUpgradePrecheckId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MajorVersionUpgradePrecheckId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PrecheckValidationId, ok = input.Parsed["precheckValidationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "precheckValidationId", input)
	}

	return nil
}

// ValidateMajorVersionUpgradePrecheckID checks that 'input' can be parsed as a Major Version Upgrade Precheck ID
func ValidateMajorVersionUpgradePrecheckID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMajorVersionUpgradePrecheckID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Major Version Upgrade Precheck ID
func (id MajorVersionUpgradePrecheckId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/flexibleServers/%s/majorVersionUpgradePrecheck/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FlexibleServerName, id.PrecheckValidationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Major Version Upgrade Precheck ID
func (id MajorVersionUpgradePrecheckId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticFlexibleServers", "flexibleServers", "flexibleServers"),
		resourceids.UserSpecifiedSegment("flexibleServerName", "flexibleServerName"),
		resourceids.StaticSegment("staticMajorVersionUpgradePrecheck", "majorVersionUpgradePrecheck", "majorVersionUpgradePrecheck"),
		resourceids.UserSpecifiedSegment("precheckValidationId", "precheckValidationId"),
	}
}

// String returns a human-readable description of this Major Version Upgrade Precheck ID
func (id MajorVersionUpgradePrecheckId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Flexible Server Name: %q", id.FlexibleServerName),
		fmt.Sprintf("Precheck Validation: %q", id.PrecheckValidationId),
	}
	return fmt.Sprintf("Major Version Upgrade Precheck (%s)", strings.Join(components, "\n"))
}
