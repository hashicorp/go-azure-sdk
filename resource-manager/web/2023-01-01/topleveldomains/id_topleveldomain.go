package topleveldomains

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TopLevelDomainId{})
}

var _ resourceids.ResourceId = &TopLevelDomainId{}

// TopLevelDomainId is a struct representing the Resource ID for a Top Level Domain
type TopLevelDomainId struct {
	SubscriptionId     string
	TopLevelDomainName string
}

// NewTopLevelDomainID returns a new TopLevelDomainId struct
func NewTopLevelDomainID(subscriptionId string, topLevelDomainName string) TopLevelDomainId {
	return TopLevelDomainId{
		SubscriptionId:     subscriptionId,
		TopLevelDomainName: topLevelDomainName,
	}
}

// ParseTopLevelDomainID parses 'input' into a TopLevelDomainId
func ParseTopLevelDomainID(input string) (*TopLevelDomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TopLevelDomainId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TopLevelDomainId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTopLevelDomainIDInsensitively parses 'input' case-insensitively into a TopLevelDomainId
// note: this method should only be used for API response data and not user input
func ParseTopLevelDomainIDInsensitively(input string) (*TopLevelDomainId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TopLevelDomainId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TopLevelDomainId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TopLevelDomainId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.TopLevelDomainName, ok = input.Parsed["topLevelDomainName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "topLevelDomainName", input)
	}

	return nil
}

// ValidateTopLevelDomainID checks that 'input' can be parsed as a Top Level Domain ID
func ValidateTopLevelDomainID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTopLevelDomainID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Top Level Domain ID
func (id TopLevelDomainId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.DomainRegistration/topLevelDomains/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.TopLevelDomainName)
}

// Segments returns a slice of Resource ID Segments which comprise this Top Level Domain ID
func (id TopLevelDomainId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDomainRegistration", "Microsoft.DomainRegistration", "Microsoft.DomainRegistration"),
		resourceids.StaticSegment("staticTopLevelDomains", "topLevelDomains", "topLevelDomains"),
		resourceids.UserSpecifiedSegment("topLevelDomainName", "topLevelDomainValue"),
	}
}

// String returns a human-readable description of this Top Level Domain ID
func (id TopLevelDomainId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Top Level Domain Name: %q", id.TopLevelDomainName),
	}
	return fmt.Sprintf("Top Level Domain (%s)", strings.Join(components, "\n"))
}
