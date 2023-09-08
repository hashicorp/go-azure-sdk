package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectionSettingsOperationPredicate struct {
	EnableSingleSignOn *bool
	ODataType          *string
}

func (p CloudPCConnectionSettingsOperationPredicate) Matches(input CloudPCConnectionSettings) bool {

	if p.EnableSingleSignOn != nil && (input.EnableSingleSignOn == nil || *p.EnableSingleSignOn != *input.EnableSingleSignOn) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
