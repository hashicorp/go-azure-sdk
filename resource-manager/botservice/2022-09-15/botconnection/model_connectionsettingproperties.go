package botconnection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionSettingProperties struct {
	ClientId                   *string                       `json:"clientId,omitempty"`
	ClientSecret               *string                       `json:"clientSecret,omitempty"`
	Id                         *string                       `json:"id,omitempty"`
	Name                       *string                       `json:"name,omitempty"`
	Parameters                 *[]ConnectionSettingParameter `json:"parameters,omitempty"`
	ProvisioningState          *string                       `json:"provisioningState,omitempty"`
	Scopes                     *string                       `json:"scopes,omitempty"`
	ServiceProviderDisplayName *string                       `json:"serviceProviderDisplayName,omitempty"`
	ServiceProviderId          *string                       `json:"serviceProviderId,omitempty"`
	SettingId                  *string                       `json:"settingId,omitempty"`
}
