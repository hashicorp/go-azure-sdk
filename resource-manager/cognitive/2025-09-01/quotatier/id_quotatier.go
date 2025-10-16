package quotatier

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&QuotaTierId{})
}

var _ resourceids.ResourceId = &QuotaTierId{}

// QuotaTierId is a struct representing the Resource ID for a Quota Tier
type QuotaTierId struct {
	SubscriptionId string
	QuotaTierName  string
}

// NewQuotaTierID returns a new QuotaTierId struct
func NewQuotaTierID(subscriptionId string, quotaTierName string) QuotaTierId {
	return QuotaTierId{
		SubscriptionId: subscriptionId,
		QuotaTierName:  quotaTierName,
	}
}

// ParseQuotaTierID parses 'input' into a QuotaTierId
func ParseQuotaTierID(input string) (*QuotaTierId, error) {
	parser := resourceids.NewParserFromResourceIdType(&QuotaTierId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := QuotaTierId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseQuotaTierIDInsensitively parses 'input' case-insensitively into a QuotaTierId
// note: this method should only be used for API response data and not user input
func ParseQuotaTierIDInsensitively(input string) (*QuotaTierId, error) {
	parser := resourceids.NewParserFromResourceIdType(&QuotaTierId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := QuotaTierId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *QuotaTierId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.QuotaTierName, ok = input.Parsed["quotaTierName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "quotaTierName", input)
	}

	return nil
}

// ValidateQuotaTierID checks that 'input' can be parsed as a Quota Tier ID
func ValidateQuotaTierID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseQuotaTierID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Quota Tier ID
func (id QuotaTierId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.CognitiveServices/quotaTiers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.QuotaTierName)
}

// Segments returns a slice of Resource ID Segments which comprise this Quota Tier ID
func (id QuotaTierId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticQuotaTiers", "quotaTiers", "quotaTiers"),
		resourceids.UserSpecifiedSegment("quotaTierName", "quotaTierName"),
	}
}

// String returns a human-readable description of this Quota Tier ID
func (id QuotaTierId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Quota Tier Name: %q", id.QuotaTierName),
	}
	return fmt.Sprintf("Quota Tier (%s)", strings.Join(components, "\n"))
}
