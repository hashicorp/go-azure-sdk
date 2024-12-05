package officeconsents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&OfficeConsentId{})
}

var _ resourceids.ResourceId = &OfficeConsentId{}

// OfficeConsentId is a struct representing the Resource ID for a Office Consent
type OfficeConsentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	ConsentId         string
}

// NewOfficeConsentID returns a new OfficeConsentId struct
func NewOfficeConsentID(subscriptionId string, resourceGroupName string, workspaceName string, consentId string) OfficeConsentId {
	return OfficeConsentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		ConsentId:         consentId,
	}
}

// ParseOfficeConsentID parses 'input' into a OfficeConsentId
func ParseOfficeConsentID(input string) (*OfficeConsentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OfficeConsentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OfficeConsentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOfficeConsentIDInsensitively parses 'input' case-insensitively into a OfficeConsentId
// note: this method should only be used for API response data and not user input
func ParseOfficeConsentIDInsensitively(input string) (*OfficeConsentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OfficeConsentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OfficeConsentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OfficeConsentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.ConsentId, ok = input.Parsed["consentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "consentId", input)
	}

	return nil
}

// ValidateOfficeConsentID checks that 'input' can be parsed as a Office Consent ID
func ValidateOfficeConsentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOfficeConsentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Office Consent ID
func (id OfficeConsentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/officeConsents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ConsentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Office Consent ID
func (id OfficeConsentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticOfficeConsents", "officeConsents", "officeConsents"),
		resourceids.UserSpecifiedSegment("consentId", "consentId"),
	}
}

// String returns a human-readable description of this Office Consent ID
func (id OfficeConsentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Consent: %q", id.ConsentId),
	}
	return fmt.Sprintf("Office Consent (%s)", strings.Join(components, "\n"))
}
