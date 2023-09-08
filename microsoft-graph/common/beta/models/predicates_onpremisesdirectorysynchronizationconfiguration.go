package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronizationConfigurationOperationPredicate struct {
	AnchorAttribute                          *string
	ApplicationId                            *string
	CustomerRequestedSynchronizationInterval *string
	ODataType                                *string
	SynchronizationClientVersion             *string
	SynchronizationInterval                  *string
}

func (p OnPremisesDirectorySynchronizationConfigurationOperationPredicate) Matches(input OnPremisesDirectorySynchronizationConfiguration) bool {

	if p.AnchorAttribute != nil && (input.AnchorAttribute == nil || *p.AnchorAttribute != *input.AnchorAttribute) {
		return false
	}

	if p.ApplicationId != nil && (input.ApplicationId == nil || *p.ApplicationId != *input.ApplicationId) {
		return false
	}

	if p.CustomerRequestedSynchronizationInterval != nil && (input.CustomerRequestedSynchronizationInterval == nil || *p.CustomerRequestedSynchronizationInterval != *input.CustomerRequestedSynchronizationInterval) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SynchronizationClientVersion != nil && (input.SynchronizationClientVersion == nil || *p.SynchronizationClientVersion != *input.SynchronizationClientVersion) {
		return false
	}

	if p.SynchronizationInterval != nil && (input.SynchronizationInterval == nil || *p.SynchronizationInterval != *input.SynchronizationInterval) {
		return false
	}

	return true
}
