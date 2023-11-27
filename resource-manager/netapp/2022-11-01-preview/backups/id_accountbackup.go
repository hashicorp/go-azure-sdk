package backups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AccountBackupId{}

// AccountBackupId is a struct representing the Resource ID for a Account Backup
type AccountBackupId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	AccountBackupName string
}

// NewAccountBackupID returns a new AccountBackupId struct
func NewAccountBackupID(subscriptionId string, resourceGroupName string, netAppAccountName string, accountBackupName string) AccountBackupId {
	return AccountBackupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		AccountBackupName: accountBackupName,
	}
}

// ParseAccountBackupID parses 'input' into a AccountBackupId
func ParseAccountBackupID(input string) (*AccountBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountBackupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AccountBackupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAccountBackupIDInsensitively parses 'input' case-insensitively into a AccountBackupId
// note: this method should only be used for API response data and not user input
func ParseAccountBackupIDInsensitively(input string) (*AccountBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountBackupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AccountBackupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AccountBackupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetAppAccountName, ok = input.Parsed["netAppAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", input)
	}

	if id.AccountBackupName, ok = input.Parsed["accountBackupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountBackupName", input)
	}

	return nil
}

// ValidateAccountBackupID checks that 'input' can be parsed as a Account Backup ID
func ValidateAccountBackupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccountBackupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Account Backup ID
func (id AccountBackupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/accountBackups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.AccountBackupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Account Backup ID
func (id AccountBackupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountValue"),
		resourceids.StaticSegment("staticAccountBackups", "accountBackups", "accountBackups"),
		resourceids.UserSpecifiedSegment("accountBackupName", "accountBackupValue"),
	}
}

// String returns a human-readable description of this Account Backup ID
func (id AccountBackupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Account Backup Name: %q", id.AccountBackupName),
	}
	return fmt.Sprintf("Account Backup (%s)", strings.Join(components, "\n"))
}
