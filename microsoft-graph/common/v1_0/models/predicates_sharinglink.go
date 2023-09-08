package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingLinkOperationPredicate struct {
	ODataType        *string
	PreventsDownload *bool
	Scope            *string
	Type             *string
	WebHtml          *string
	WebUrl           *string
}

func (p SharingLinkOperationPredicate) Matches(input SharingLink) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PreventsDownload != nil && (input.PreventsDownload == nil || *p.PreventsDownload != *input.PreventsDownload) {
		return false
	}

	if p.Scope != nil && (input.Scope == nil || *p.Scope != *input.Scope) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.WebHtml != nil && (input.WebHtml == nil || *p.WebHtml != *input.WebHtml) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
