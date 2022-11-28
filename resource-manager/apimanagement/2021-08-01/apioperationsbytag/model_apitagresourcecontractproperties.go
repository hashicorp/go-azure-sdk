package apioperationsbytag

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiTagResourceContractProperties struct {
	ApiRevision                   *string                                `json:"apiRevision,omitempty"`
	ApiRevisionDescription        *string                                `json:"apiRevisionDescription,omitempty"`
	ApiVersion                    *string                                `json:"apiVersion,omitempty"`
	ApiVersionDescription         *string                                `json:"apiVersionDescription,omitempty"`
	ApiVersionSetId               *string                                `json:"apiVersionSetId,omitempty"`
	AuthenticationSettings        *AuthenticationSettingsContract        `json:"authenticationSettings"`
	Contact                       *ApiContactInformation                 `json:"contact"`
	Description                   *string                                `json:"description,omitempty"`
	Id                            *string                                `json:"id,omitempty"`
	IsCurrent                     *bool                                  `json:"isCurrent,omitempty"`
	IsOnline                      *bool                                  `json:"isOnline,omitempty"`
	License                       *ApiLicenseInformation                 `json:"license"`
	Name                          *string                                `json:"name,omitempty"`
	Path                          *string                                `json:"path,omitempty"`
	Protocols                     *[]Protocol                            `json:"protocols,omitempty"`
	ServiceUrl                    *string                                `json:"serviceUrl,omitempty"`
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract `json:"subscriptionKeyParameterNames"`
	SubscriptionRequired          *bool                                  `json:"subscriptionRequired,omitempty"`
	TermsOfServiceUrl             *string                                `json:"termsOfServiceUrl,omitempty"`
	Type                          *ApiType                               `json:"type,omitempty"`
}
