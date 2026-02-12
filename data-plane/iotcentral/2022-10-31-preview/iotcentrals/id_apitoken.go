package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApiTokenId{}

// ApiTokenId is a struct representing the Resource ID for a Api Token
type ApiTokenId struct {
	BaseURI string
	TokenId string
}

// NewApiTokenID returns a new ApiTokenId struct
func NewApiTokenID(baseURI string, tokenId string) ApiTokenId {
	return ApiTokenId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		TokenId: tokenId,
	}
}

// ParseApiTokenID parses 'input' into a ApiTokenId
func ParseApiTokenID(input string) (*ApiTokenId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApiTokenId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApiTokenId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApiTokenIDInsensitively parses 'input' case-insensitively into a ApiTokenId
// note: this method should only be used for API response data and not user input
func ParseApiTokenIDInsensitively(input string) (*ApiTokenId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApiTokenId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApiTokenId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApiTokenId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.TokenId, ok = input.Parsed["tokenId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tokenId", input)
	}

	return nil
}

// ValidateApiTokenID checks that 'input' can be parsed as a Api Token ID
func ValidateApiTokenID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApiTokenID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Api Token ID
func (id ApiTokenId) ID() string {
	fmtString := "%s/apiTokens/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.TokenId)
}

// Path returns the formatted Api Token ID without the BaseURI
func (id ApiTokenId) Path() string {
	fmtString := "/apiTokens/%s"
	return fmt.Sprintf(fmtString, id.TokenId)
}

// PathElements returns the values of Api Token ID Segments without the BaseURI
func (id ApiTokenId) PathElements() []any {
	return []any{id.TokenId}
}

// Segments returns a slice of Resource ID Segments which comprise this Api Token ID
func (id ApiTokenId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticApiTokens", "apiTokens", "apiTokens"),
		resourceids.UserSpecifiedSegment("tokenId", "tokenId"),
	}
}

// String returns a human-readable description of this Api Token ID
func (id ApiTokenId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Token: %q", id.TokenId),
	}
	return fmt.Sprintf("Api Token (%s)", strings.Join(components, "\n"))
}
