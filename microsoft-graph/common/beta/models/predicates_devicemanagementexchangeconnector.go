package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorOperationPredicate struct {
	ConnectorServerName  *string
	ExchangeAlias        *string
	ExchangeOrganization *string
	Id                   *string
	LastSyncDateTime     *string
	ODataType            *string
	PrimarySmtpAddress   *string
	ServerName           *string
	Version              *string
}

func (p DeviceManagementExchangeConnectorOperationPredicate) Matches(input DeviceManagementExchangeConnector) bool {

	if p.ConnectorServerName != nil && (input.ConnectorServerName == nil || *p.ConnectorServerName != *input.ConnectorServerName) {
		return false
	}

	if p.ExchangeAlias != nil && (input.ExchangeAlias == nil || *p.ExchangeAlias != *input.ExchangeAlias) {
		return false
	}

	if p.ExchangeOrganization != nil && (input.ExchangeOrganization == nil || *p.ExchangeOrganization != *input.ExchangeOrganization) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrimarySmtpAddress != nil && (input.PrimarySmtpAddress == nil || *p.PrimarySmtpAddress != *input.PrimarySmtpAddress) {
		return false
	}

	if p.ServerName != nil && (input.ServerName == nil || *p.ServerName != *input.ServerName) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
