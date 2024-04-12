package securitycontacts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SecurityContactId{})
}

var _ resourceids.ResourceId = &SecurityContactId{}

// SecurityContactId is a struct representing the Resource ID for a Security Contact
type SecurityContactId struct {
	SubscriptionId      string
	SecurityContactName string
}

// NewSecurityContactID returns a new SecurityContactId struct
func NewSecurityContactID(subscriptionId string, securityContactName string) SecurityContactId {
	return SecurityContactId{
		SubscriptionId:      subscriptionId,
		SecurityContactName: securityContactName,
	}
}

// ParseSecurityContactID parses 'input' into a SecurityContactId
func ParseSecurityContactID(input string) (*SecurityContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecurityContactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecurityContactId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSecurityContactIDInsensitively parses 'input' case-insensitively into a SecurityContactId
// note: this method should only be used for API response data and not user input
func ParseSecurityContactIDInsensitively(input string) (*SecurityContactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SecurityContactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SecurityContactId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SecurityContactId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.SecurityContactName, ok = input.Parsed["securityContactName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "securityContactName", input)
	}

	return nil
}

// ValidateSecurityContactID checks that 'input' can be parsed as a Security Contact ID
func ValidateSecurityContactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSecurityContactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Security Contact ID
func (id SecurityContactId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/securityContacts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.SecurityContactName)
}

// Segments returns a slice of Resource ID Segments which comprise this Security Contact ID
func (id SecurityContactId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecurityContacts", "securityContacts", "securityContacts"),
		resourceids.UserSpecifiedSegment("securityContactName", "securityContactValue"),
	}
}

// String returns a human-readable description of this Security Contact ID
func (id SecurityContactId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Security Contact Name: %q", id.SecurityContactName),
	}
	return fmt.Sprintf("Security Contact (%s)", strings.Join(components, "\n"))
}
