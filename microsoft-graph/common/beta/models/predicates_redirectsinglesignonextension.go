package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RedirectSingleSignOnExtensionOperationPredicate struct {
	ExtensionIdentifier *string
	ODataType           *string
	TeamIdentifier      *string
}

func (p RedirectSingleSignOnExtensionOperationPredicate) Matches(input RedirectSingleSignOnExtension) bool {

	if p.ExtensionIdentifier != nil && (input.ExtensionIdentifier == nil || *p.ExtensionIdentifier != *input.ExtensionIdentifier) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamIdentifier != nil && (input.TeamIdentifier == nil || *p.TeamIdentifier != *input.TeamIdentifier) {
		return false
	}

	return true
}
