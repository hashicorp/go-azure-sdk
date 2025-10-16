package encryptionscopes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&EncryptionScopeId{})
}

var _ resourceids.ResourceId = &EncryptionScopeId{}

// EncryptionScopeId is a struct representing the Resource ID for a Encryption Scope
type EncryptionScopeId struct {
	SubscriptionId      string
	ResourceGroupName   string
	AccountName         string
	EncryptionScopeName string
}

// NewEncryptionScopeID returns a new EncryptionScopeId struct
func NewEncryptionScopeID(subscriptionId string, resourceGroupName string, accountName string, encryptionScopeName string) EncryptionScopeId {
	return EncryptionScopeId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		AccountName:         accountName,
		EncryptionScopeName: encryptionScopeName,
	}
}

// ParseEncryptionScopeID parses 'input' into a EncryptionScopeId
func ParseEncryptionScopeID(input string) (*EncryptionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EncryptionScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EncryptionScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEncryptionScopeIDInsensitively parses 'input' case-insensitively into a EncryptionScopeId
// note: this method should only be used for API response data and not user input
func ParseEncryptionScopeIDInsensitively(input string) (*EncryptionScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EncryptionScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EncryptionScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EncryptionScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AccountName, ok = input.Parsed["accountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountName", input)
	}

	if id.EncryptionScopeName, ok = input.Parsed["encryptionScopeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "encryptionScopeName", input)
	}

	return nil
}

// ValidateEncryptionScopeID checks that 'input' can be parsed as a Encryption Scope ID
func ValidateEncryptionScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEncryptionScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Encryption Scope ID
func (id EncryptionScopeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/encryptionScopes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.EncryptionScopeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Encryption Scope ID
func (id EncryptionScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticEncryptionScopes", "encryptionScopes", "encryptionScopes"),
		resourceids.UserSpecifiedSegment("encryptionScopeName", "encryptionScopeName"),
	}
}

// String returns a human-readable description of this Encryption Scope ID
func (id EncryptionScopeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Encryption Scope Name: %q", id.EncryptionScopeName),
	}
	return fmt.Sprintf("Encryption Scope (%s)", strings.Join(components, "\n"))
}
