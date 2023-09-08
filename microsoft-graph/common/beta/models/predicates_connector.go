package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorOperationPredicate struct {
	ExternalIp  *string
	Id          *string
	MachineName *string
	ODataType   *string
	Version     *string
}

func (p ConnectorOperationPredicate) Matches(input Connector) bool {

	if p.ExternalIp != nil && (input.ExternalIp == nil || *p.ExternalIp != *input.ExternalIp) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MachineName != nil && (input.MachineName == nil || *p.MachineName != *input.MachineName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
