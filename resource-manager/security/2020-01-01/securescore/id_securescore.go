package securescore

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SecureScoreId{})
}

var _ resourceids.ResourceId = &SecureScoreId{}

// SecureScoreId is a struct representing the Resource ID for a Secure Score
type SecureScoreId struct {
	SubscriptionId  string
	SecureScoreName string
}

// NewSecureScoreID returns a new SecureScoreId struct
func NewSecureScoreID(subscriptionId string, secureScoreName string) SecureScoreId {
	return SecureScoreId{
		SubscriptionId:  subscriptionId,
		SecureScoreName: secureScoreName,
	}
}

// ParseSecureScoreID parses 'input' into a SecureScoreId
func ParseSecureScoreID(input string) (*SecureScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecureScoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecureScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSecureScoreIDInsensitively parses 'input' case-insensitively into a SecureScoreId
// note: this method should only be used for API response data and not user input
func ParseSecureScoreIDInsensitively(input string) (*SecureScoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecureScoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecureScoreId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SecureScoreId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.SecureScoreName, ok = input.Parsed["secureScoreName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "secureScoreName", input)
	}

	return nil
}

// ValidateSecureScoreID checks that 'input' can be parsed as a Secure Score ID
func ValidateSecureScoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSecureScoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Secure Score ID
func (id SecureScoreId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/secureScores/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.SecureScoreName)
}

// Segments returns a slice of Resource ID Segments which comprise this Secure Score ID
func (id SecureScoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecureScores", "secureScores", "secureScores"),
		resourceids.UserSpecifiedSegment("secureScoreName", "secureScoreName"),
	}
}

// String returns a human-readable description of this Secure Score ID
func (id SecureScoreId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Secure Score Name: %q", id.SecureScoreName),
	}
	return fmt.Sprintf("Secure Score (%s)", strings.Join(components, "\n"))
}
