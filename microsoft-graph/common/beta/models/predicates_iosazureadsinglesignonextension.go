package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosAzureAdSingleSignOnExtensionOperationPredicate struct {
	EnableSharedDeviceMode *bool
	ODataType              *string
}

func (p IosAzureAdSingleSignOnExtensionOperationPredicate) Matches(input IosAzureAdSingleSignOnExtension) bool {

	if p.EnableSharedDeviceMode != nil && (input.EnableSharedDeviceMode == nil || *p.EnableSharedDeviceMode != *input.EnableSharedDeviceMode) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
