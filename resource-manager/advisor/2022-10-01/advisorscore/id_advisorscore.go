package advisorscore

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AdvisorScoreId{})
}

var _ resourceids.ResourceId = &AdvisorScoreId{}

// AdvisorScoreId is a struct representing the Resource ID for a Advisor Score
type AdvisorScoreId struct {
	SubscriptionId   string
	AdvisorScoreName string
}

// NewAdvisorScoreID returns a new AdvisorScoreId struct
func NewAdvisorScoreID(subscriptionId string, advisorScoreName string) AdvisorScoreId {
	return AdvisorScoreId{
		SubscriptionId:   subscriptionId,
		AdvisorScoreName: advisorScoreName,
	}
}

// ParseAdvisorScoreID parses 'input' into a AdvisorScoreId
func ParseAdvisorScoreID(input string) (*AdvisorScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdvisorScoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdvisorScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAdvisorScoreIDInsensitively parses 'input' case-insensitively into a AdvisorScoreId
// note: this method should only be used for API response data and not user input
func ParseAdvisorScoreIDInsensitively(input string) (*AdvisorScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AdvisorScoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AdvisorScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AdvisorScoreId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.AdvisorScoreName, ok = input.Parsed["advisorScoreName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "advisorScoreName", input)
	}

	return nil
}

// ValidateAdvisorScoreID checks that 'input' can be parsed as a Advisor Score ID
func ValidateAdvisorScoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdvisorScoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Advisor Score ID
func (id AdvisorScoreId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Advisor/advisorScore/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.AdvisorScoreName)
}

// Segments returns a slice of Resource ID Segments which comprise this Advisor Score ID
func (id AdvisorScoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAdvisor", "Microsoft.Advisor", "Microsoft.Advisor"),
		resourceids.StaticSegment("staticAdvisorScore", "advisorScore", "advisorScore"),
		resourceids.UserSpecifiedSegment("advisorScoreName", "advisorScoreValue"),
	}
}

// String returns a human-readable description of this Advisor Score ID
func (id AdvisorScoreId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Advisor Score Name: %q", id.AdvisorScoreName),
	}
	return fmt.Sprintf("Advisor Score (%s)", strings.Join(components, "\n"))
}
