package defenderforaisettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DefenderForAISettingId{})
}

var _ resourceids.ResourceId = &DefenderForAISettingId{}

// DefenderForAISettingId is a struct representing the Resource ID for a Defender For A I Setting
type DefenderForAISettingId struct {
	SubscriptionId           string
	ResourceGroupName        string
	AccountName              string
	DefenderForAISettingName string
}

// NewDefenderForAISettingID returns a new DefenderForAISettingId struct
func NewDefenderForAISettingID(subscriptionId string, resourceGroupName string, accountName string, defenderForAISettingName string) DefenderForAISettingId {
	return DefenderForAISettingId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		AccountName:              accountName,
		DefenderForAISettingName: defenderForAISettingName,
	}
}

// ParseDefenderForAISettingID parses 'input' into a DefenderForAISettingId
func ParseDefenderForAISettingID(input string) (*DefenderForAISettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DefenderForAISettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DefenderForAISettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDefenderForAISettingIDInsensitively parses 'input' case-insensitively into a DefenderForAISettingId
// note: this method should only be used for API response data and not user input
func ParseDefenderForAISettingIDInsensitively(input string) (*DefenderForAISettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DefenderForAISettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DefenderForAISettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DefenderForAISettingId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DefenderForAISettingName, ok = input.Parsed["defenderForAISettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "defenderForAISettingName", input)
	}

	return nil
}

// ValidateDefenderForAISettingID checks that 'input' can be parsed as a Defender For A I Setting ID
func ValidateDefenderForAISettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDefenderForAISettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Defender For A I Setting ID
func (id DefenderForAISettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/defenderForAISettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.DefenderForAISettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Defender For A I Setting ID
func (id DefenderForAISettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticDefenderForAISettings", "defenderForAISettings", "defenderForAISettings"),
		resourceids.UserSpecifiedSegment("defenderForAISettingName", "defenderForAISettingName"),
	}
}

// String returns a human-readable description of this Defender For A I Setting ID
func (id DefenderForAISettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Defender For A I Setting Name: %q", id.DefenderForAISettingName),
	}
	return fmt.Sprintf("Defender For A I Setting (%s)", strings.Join(components, "\n"))
}
