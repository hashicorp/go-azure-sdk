package policysetdefinitionversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProviderPolicySetDefinitionVersionId{})
}

var _ resourceids.ResourceId = &ProviderPolicySetDefinitionVersionId{}

// ProviderPolicySetDefinitionVersionId is a struct representing the Resource ID for a Provider Policy Set Definition Version
type ProviderPolicySetDefinitionVersionId struct {
	SubscriptionId          string
	PolicySetDefinitionName string
	VersionName             string
}

// NewProviderPolicySetDefinitionVersionID returns a new ProviderPolicySetDefinitionVersionId struct
func NewProviderPolicySetDefinitionVersionID(subscriptionId string, policySetDefinitionName string, versionName string) ProviderPolicySetDefinitionVersionId {
	return ProviderPolicySetDefinitionVersionId{
		SubscriptionId:          subscriptionId,
		PolicySetDefinitionName: policySetDefinitionName,
		VersionName:             versionName,
	}
}

// ParseProviderPolicySetDefinitionVersionID parses 'input' into a ProviderPolicySetDefinitionVersionId
func ParseProviderPolicySetDefinitionVersionID(input string) (*ProviderPolicySetDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderPolicySetDefinitionVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderPolicySetDefinitionVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderPolicySetDefinitionVersionIDInsensitively parses 'input' case-insensitively into a ProviderPolicySetDefinitionVersionId
// note: this method should only be used for API response data and not user input
func ParseProviderPolicySetDefinitionVersionIDInsensitively(input string) (*ProviderPolicySetDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderPolicySetDefinitionVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderPolicySetDefinitionVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderPolicySetDefinitionVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicySetDefinitionName, ok = input.Parsed["policySetDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateProviderPolicySetDefinitionVersionID checks that 'input' can be parsed as a Provider Policy Set Definition Version ID
func ValidateProviderPolicySetDefinitionVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderPolicySetDefinitionVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Policy Set Definition Version ID
func (id ProviderPolicySetDefinitionVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policySetDefinitions/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicySetDefinitionName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Policy Set Definition Version ID
func (id ProviderPolicySetDefinitionVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicySetDefinitions", "policySetDefinitions", "policySetDefinitions"),
		resourceids.UserSpecifiedSegment("policySetDefinitionName", "policySetDefinitionName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Provider Policy Set Definition Version ID
func (id ProviderPolicySetDefinitionVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy Set Definition Name: %q", id.PolicySetDefinitionName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Provider Policy Set Definition Version (%s)", strings.Join(components, "\n"))
}
