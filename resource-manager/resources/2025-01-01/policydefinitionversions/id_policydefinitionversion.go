package policydefinitionversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PolicyDefinitionVersionId{})
}

var _ resourceids.ResourceId = &PolicyDefinitionVersionId{}

// PolicyDefinitionVersionId is a struct representing the Resource ID for a Policy Definition Version
type PolicyDefinitionVersionId struct {
	SubscriptionId       string
	PolicyDefinitionName string
	VersionName          string
}

// NewPolicyDefinitionVersionID returns a new PolicyDefinitionVersionId struct
func NewPolicyDefinitionVersionID(subscriptionId string, policyDefinitionName string, versionName string) PolicyDefinitionVersionId {
	return PolicyDefinitionVersionId{
		SubscriptionId:       subscriptionId,
		PolicyDefinitionName: policyDefinitionName,
		VersionName:          versionName,
	}
}

// ParsePolicyDefinitionVersionID parses 'input' into a PolicyDefinitionVersionId
func ParsePolicyDefinitionVersionID(input string) (*PolicyDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyDefinitionVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyDefinitionVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyDefinitionVersionIDInsensitively parses 'input' case-insensitively into a PolicyDefinitionVersionId
// note: this method should only be used for API response data and not user input
func ParsePolicyDefinitionVersionIDInsensitively(input string) (*PolicyDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyDefinitionVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyDefinitionVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyDefinitionVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicyDefinitionName, ok = input.Parsed["policyDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidatePolicyDefinitionVersionID checks that 'input' can be parsed as a Policy Definition Version ID
func ValidatePolicyDefinitionVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyDefinitionVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Definition Version ID
func (id PolicyDefinitionVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policyDefinitions/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicyDefinitionName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Definition Version ID
func (id PolicyDefinitionVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyDefinitions", "policyDefinitions", "policyDefinitions"),
		resourceids.UserSpecifiedSegment("policyDefinitionName", "policyDefinitionName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Policy Definition Version ID
func (id PolicyDefinitionVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy Definition Name: %q", id.PolicyDefinitionName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Policy Definition Version (%s)", strings.Join(components, "\n"))
}
