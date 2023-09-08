package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSExtensionsConfigurationOperationPredicate struct {
	CreatedDateTime                 *string
	Description                     *string
	DisplayName                     *string
	Id                              *string
	KernelExtensionOverridesAllowed *bool
	LastModifiedDateTime            *string
	ODataType                       *string
	SupportsScopeTags               *bool
	SystemExtensionsBlockOverride   *bool
	Version                         *int64
}

func (p MacOSExtensionsConfigurationOperationPredicate) Matches(input MacOSExtensionsConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.KernelExtensionOverridesAllowed != nil && (input.KernelExtensionOverridesAllowed == nil || *p.KernelExtensionOverridesAllowed != *input.KernelExtensionOverridesAllowed) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.SystemExtensionsBlockOverride != nil && (input.SystemExtensionsBlockOverride == nil || *p.SystemExtensionsBlockOverride != *input.SystemExtensionsBlockOverride) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
