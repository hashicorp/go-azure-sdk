package testraiexternalsafetyprovider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TestRaiExternalSafetyProviderId{})
}

var _ resourceids.ResourceId = &TestRaiExternalSafetyProviderId{}

// TestRaiExternalSafetyProviderId is a struct representing the Resource ID for a Test Rai External Safety Provider
type TestRaiExternalSafetyProviderId struct {
	SubscriptionId                    string
	ResourceGroupName                 string
	AccountName                       string
	TestRaiExternalSafetyProviderName string
}

// NewTestRaiExternalSafetyProviderID returns a new TestRaiExternalSafetyProviderId struct
func NewTestRaiExternalSafetyProviderID(subscriptionId string, resourceGroupName string, accountName string, testRaiExternalSafetyProviderName string) TestRaiExternalSafetyProviderId {
	return TestRaiExternalSafetyProviderId{
		SubscriptionId:                    subscriptionId,
		ResourceGroupName:                 resourceGroupName,
		AccountName:                       accountName,
		TestRaiExternalSafetyProviderName: testRaiExternalSafetyProviderName,
	}
}

// ParseTestRaiExternalSafetyProviderID parses 'input' into a TestRaiExternalSafetyProviderId
func ParseTestRaiExternalSafetyProviderID(input string) (*TestRaiExternalSafetyProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TestRaiExternalSafetyProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TestRaiExternalSafetyProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTestRaiExternalSafetyProviderIDInsensitively parses 'input' case-insensitively into a TestRaiExternalSafetyProviderId
// note: this method should only be used for API response data and not user input
func ParseTestRaiExternalSafetyProviderIDInsensitively(input string) (*TestRaiExternalSafetyProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TestRaiExternalSafetyProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TestRaiExternalSafetyProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TestRaiExternalSafetyProviderId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TestRaiExternalSafetyProviderName, ok = input.Parsed["testRaiExternalSafetyProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "testRaiExternalSafetyProviderName", input)
	}

	return nil
}

// ValidateTestRaiExternalSafetyProviderID checks that 'input' can be parsed as a Test Rai External Safety Provider ID
func ValidateTestRaiExternalSafetyProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTestRaiExternalSafetyProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Test Rai External Safety Provider ID
func (id TestRaiExternalSafetyProviderId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/testRaiExternalSafetyProvider/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.TestRaiExternalSafetyProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Test Rai External Safety Provider ID
func (id TestRaiExternalSafetyProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticTestRaiExternalSafetyProvider", "testRaiExternalSafetyProvider", "testRaiExternalSafetyProvider"),
		resourceids.UserSpecifiedSegment("testRaiExternalSafetyProviderName", "testRaiExternalSafetyProviderName"),
	}
}

// String returns a human-readable description of this Test Rai External Safety Provider ID
func (id TestRaiExternalSafetyProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Test Rai External Safety Provider Name: %q", id.TestRaiExternalSafetyProviderName),
	}
	return fmt.Sprintf("Test Rai External Safety Provider (%s)", strings.Join(components, "\n"))
}
