package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplicationOperationPredicate struct {
	HomePageUrl *string
	LogoutUrl   *string
	ODataType   *string
}

func (p WebApplicationOperationPredicate) Matches(input WebApplication) bool {

	if p.HomePageUrl != nil && (input.HomePageUrl == nil || *p.HomePageUrl != *input.HomePageUrl) {
		return false
	}

	if p.LogoutUrl != nil && (input.LogoutUrl == nil || *p.LogoutUrl != *input.LogoutUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
