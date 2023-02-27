package invitation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = InvitationId{}

// InvitationId is a struct representing the Resource ID for a Invitation
type InvitationId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	ShareName         string
	InvitationName    string
}

// NewInvitationID returns a new InvitationId struct
func NewInvitationID(subscriptionId string, resourceGroupName string, accountName string, shareName string, invitationName string) InvitationId {
	return InvitationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		ShareName:         shareName,
		InvitationName:    invitationName,
	}
}

// ParseInvitationID parses 'input' into a InvitationId
func ParseInvitationID(input string) (*InvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(InvitationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := InvitationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ShareName, ok = parsed.Parsed["shareName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareName' was not found in the resource id %q", input)
	}

	if id.InvitationName, ok = parsed.Parsed["invitationName"]; !ok {
		return nil, fmt.Errorf("the segment 'invitationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseInvitationIDInsensitively parses 'input' case-insensitively into a InvitationId
// note: this method should only be used for API response data and not user input
func ParseInvitationIDInsensitively(input string) (*InvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(InvitationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := InvitationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.ShareName, ok = parsed.Parsed["shareName"]; !ok {
		return nil, fmt.Errorf("the segment 'shareName' was not found in the resource id %q", input)
	}

	if id.InvitationName, ok = parsed.Parsed["invitationName"]; !ok {
		return nil, fmt.Errorf("the segment 'invitationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateInvitationID checks that 'input' can be parsed as a Invitation ID
func ValidateInvitationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvitationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invitation ID
func (id InvitationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataShare/accounts/%s/shares/%s/invitations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ShareName, id.InvitationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invitation ID
func (id InvitationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataShare", "Microsoft.DataShare", "Microsoft.DataShare"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticShares", "shares", "shares"),
		resourceids.UserSpecifiedSegment("shareName", "shareValue"),
		resourceids.StaticSegment("staticInvitations", "invitations", "invitations"),
		resourceids.UserSpecifiedSegment("invitationName", "invitationValue"),
	}
}

// String returns a human-readable description of this Invitation ID
func (id InvitationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Share Name: %q", id.ShareName),
		fmt.Sprintf("Invitation Name: %q", id.InvitationName),
	}
	return fmt.Sprintf("Invitation (%s)", strings.Join(components, "\n"))
}
