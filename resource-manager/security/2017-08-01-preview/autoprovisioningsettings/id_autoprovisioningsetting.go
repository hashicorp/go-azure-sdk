package autoprovisioningsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AutoProvisioningSettingId{}

// AutoProvisioningSettingId is a struct representing the Resource ID for a Auto Provisioning Setting
type AutoProvisioningSettingId struct {
	SubscriptionId              string
	AutoProvisioningSettingName string
}

// NewAutoProvisioningSettingID returns a new AutoProvisioningSettingId struct
func NewAutoProvisioningSettingID(subscriptionId string, autoProvisioningSettingName string) AutoProvisioningSettingId {
	return AutoProvisioningSettingId{
		SubscriptionId:              subscriptionId,
		AutoProvisioningSettingName: autoProvisioningSettingName,
	}
}

// ParseAutoProvisioningSettingID parses 'input' into a AutoProvisioningSettingId
func ParseAutoProvisioningSettingID(input string) (*AutoProvisioningSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(AutoProvisioningSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AutoProvisioningSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.AutoProvisioningSettingName, ok = parsed.Parsed["autoProvisioningSettingName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "autoProvisioningSettingName", *parsed)
	}

	return &id, nil
}

// ParseAutoProvisioningSettingIDInsensitively parses 'input' case-insensitively into a AutoProvisioningSettingId
// note: this method should only be used for API response data and not user input
func ParseAutoProvisioningSettingIDInsensitively(input string) (*AutoProvisioningSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(AutoProvisioningSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AutoProvisioningSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.AutoProvisioningSettingName, ok = parsed.Parsed["autoProvisioningSettingName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "autoProvisioningSettingName", *parsed)
	}

	return &id, nil
}

// ValidateAutoProvisioningSettingID checks that 'input' can be parsed as a Auto Provisioning Setting ID
func ValidateAutoProvisioningSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAutoProvisioningSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Auto Provisioning Setting ID
func (id AutoProvisioningSettingId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/autoProvisioningSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.AutoProvisioningSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Auto Provisioning Setting ID
func (id AutoProvisioningSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticAutoProvisioningSettings", "autoProvisioningSettings", "autoProvisioningSettings"),
		resourceids.UserSpecifiedSegment("autoProvisioningSettingName", "autoProvisioningSettingValue"),
	}
}

// String returns a human-readable description of this Auto Provisioning Setting ID
func (id AutoProvisioningSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Auto Provisioning Setting Name: %q", id.AutoProvisioningSettingName),
	}
	return fmt.Sprintf("Auto Provisioning Setting (%s)", strings.Join(components, "\n"))
}
