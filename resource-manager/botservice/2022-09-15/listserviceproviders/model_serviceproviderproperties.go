package listserviceproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProviderProperties struct {
	DevPortalURL        *string                     `json:"devPortalUrl,omitempty"`
	DisplayName         *string                     `json:"displayName,omitempty"`
	IconURL             *string                     `json:"iconUrl,omitempty"`
	Id                  *string                     `json:"id,omitempty"`
	Parameters          *[]ServiceProviderParameter `json:"parameters,omitempty"`
	ServiceProviderName *string                     `json:"serviceProviderName,omitempty"`
}
