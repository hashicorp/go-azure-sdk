package senderusernames

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SenderUsernameId{}

// SenderUsernameId is a struct representing the Resource ID for a Sender Username
type SenderUsernameId struct {
	SubscriptionId     string
	ResourceGroupName  string
	EmailServiceName   string
	DomainName         string
	SenderUsernameName string
}

// NewSenderUsernameID returns a new SenderUsernameId struct
func NewSenderUsernameID(subscriptionId string, resourceGroupName string, emailServiceName string, domainName string, senderUsernameName string) SenderUsernameId {
	return SenderUsernameId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		EmailServiceName:   emailServiceName,
		DomainName:         domainName,
		SenderUsernameName: senderUsernameName,
	}
}

// ParseSenderUsernameID parses 'input' into a SenderUsernameId
func ParseSenderUsernameID(input string) (*SenderUsernameId, error) {
	parser := resourceids.NewParserFromResourceIdType(SenderUsernameId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SenderUsernameId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.EmailServiceName, ok = parsed.Parsed["emailServiceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "emailServiceName", *parsed)
	}

	if id.DomainName, ok = parsed.Parsed["domainName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainName", *parsed)
	}

	if id.SenderUsernameName, ok = parsed.Parsed["senderUsernameName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "senderUsernameName", *parsed)
	}

	return &id, nil
}

// ParseSenderUsernameIDInsensitively parses 'input' case-insensitively into a SenderUsernameId
// note: this method should only be used for API response data and not user input
func ParseSenderUsernameIDInsensitively(input string) (*SenderUsernameId, error) {
	parser := resourceids.NewParserFromResourceIdType(SenderUsernameId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SenderUsernameId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.EmailServiceName, ok = parsed.Parsed["emailServiceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "emailServiceName", *parsed)
	}

	if id.DomainName, ok = parsed.Parsed["domainName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "domainName", *parsed)
	}

	if id.SenderUsernameName, ok = parsed.Parsed["senderUsernameName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "senderUsernameName", *parsed)
	}

	return &id, nil
}

// ValidateSenderUsernameID checks that 'input' can be parsed as a Sender Username ID
func ValidateSenderUsernameID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSenderUsernameID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sender Username ID
func (id SenderUsernameId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Communication/emailServices/%s/domains/%s/senderUsernames/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.EmailServiceName, id.DomainName, id.SenderUsernameName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sender Username ID
func (id SenderUsernameId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCommunication", "Microsoft.Communication", "Microsoft.Communication"),
		resourceids.StaticSegment("staticEmailServices", "emailServices", "emailServices"),
		resourceids.UserSpecifiedSegment("emailServiceName", "emailServiceValue"),
		resourceids.StaticSegment("staticDomains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainName", "domainValue"),
		resourceids.StaticSegment("staticSenderUsernames", "senderUsernames", "senderUsernames"),
		resourceids.UserSpecifiedSegment("senderUsernameName", "senderUsernameValue"),
	}
}

// String returns a human-readable description of this Sender Username ID
func (id SenderUsernameId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Email Service Name: %q", id.EmailServiceName),
		fmt.Sprintf("Domain Name: %q", id.DomainName),
		fmt.Sprintf("Sender Username Name: %q", id.SenderUsernameName),
	}
	return fmt.Sprintf("Sender Username (%s)", strings.Join(components, "\n"))
}
