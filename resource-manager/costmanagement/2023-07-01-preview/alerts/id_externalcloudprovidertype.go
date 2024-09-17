package alerts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ExternalCloudProviderTypeId{})
}

var _ resourceids.ResourceId = &ExternalCloudProviderTypeId{}

// ExternalCloudProviderTypeId is a struct representing the Resource ID for a External Cloud Provider Type
type ExternalCloudProviderTypeId struct {
	ExternalCloudProviderType ExternalCloudProviderType
	ExternalCloudProviderId   string
}

// NewExternalCloudProviderTypeID returns a new ExternalCloudProviderTypeId struct
func NewExternalCloudProviderTypeID(externalCloudProviderType ExternalCloudProviderType, externalCloudProviderId string) ExternalCloudProviderTypeId {
	return ExternalCloudProviderTypeId{
		ExternalCloudProviderType: externalCloudProviderType,
		ExternalCloudProviderId:   externalCloudProviderId,
	}
}

// ParseExternalCloudProviderTypeID parses 'input' into a ExternalCloudProviderTypeId
func ParseExternalCloudProviderTypeID(input string) (*ExternalCloudProviderTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExternalCloudProviderTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExternalCloudProviderTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExternalCloudProviderTypeIDInsensitively parses 'input' case-insensitively into a ExternalCloudProviderTypeId
// note: this method should only be used for API response data and not user input
func ParseExternalCloudProviderTypeIDInsensitively(input string) (*ExternalCloudProviderTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExternalCloudProviderTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExternalCloudProviderTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExternalCloudProviderTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if v, ok := input.Parsed["externalCloudProviderType"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "externalCloudProviderType", input)
		}

		externalCloudProviderType, err := parseExternalCloudProviderType(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.ExternalCloudProviderType = *externalCloudProviderType
	}

	if id.ExternalCloudProviderId, ok = input.Parsed["externalCloudProviderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "externalCloudProviderId", input)
	}

	return nil
}

// ValidateExternalCloudProviderTypeID checks that 'input' can be parsed as a External Cloud Provider Type ID
func ValidateExternalCloudProviderTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExternalCloudProviderTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted External Cloud Provider Type ID
func (id ExternalCloudProviderTypeId) ID() string {
	fmtString := "/providers/Microsoft.CostManagement/%s/%s"
	return fmt.Sprintf(fmtString, string(id.ExternalCloudProviderType), id.ExternalCloudProviderId)
}

// Segments returns a slice of Resource ID Segments which comprise this External Cloud Provider Type ID
func (id ExternalCloudProviderTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCostManagement", "Microsoft.CostManagement", "Microsoft.CostManagement"),
		resourceids.ConstantSegment("externalCloudProviderType", PossibleValuesForExternalCloudProviderType(), "externalBillingAccounts"),
		resourceids.UserSpecifiedSegment("externalCloudProviderId", "externalCloudProviderIdValue"),
	}
}

// String returns a human-readable description of this External Cloud Provider Type ID
func (id ExternalCloudProviderTypeId) String() string {
	components := []string{
		fmt.Sprintf("External Cloud Provider Type: %q", string(id.ExternalCloudProviderType)),
		fmt.Sprintf("External Cloud Provider: %q", id.ExternalCloudProviderId),
	}
	return fmt.Sprintf("External Cloud Provider Type (%s)", strings.Join(components, "\n"))
}
