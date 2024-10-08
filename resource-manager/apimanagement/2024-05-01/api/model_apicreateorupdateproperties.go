package api

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiCreateOrUpdateProperties struct {
	ApiRevision                      *string                                  `json:"apiRevision,omitempty"`
	ApiRevisionDescription           *string                                  `json:"apiRevisionDescription,omitempty"`
	ApiType                          *SoapApiType                             `json:"apiType,omitempty"`
	ApiVersion                       *string                                  `json:"apiVersion,omitempty"`
	ApiVersionDescription            *string                                  `json:"apiVersionDescription,omitempty"`
	ApiVersionSet                    *ApiVersionSetContractDetails            `json:"apiVersionSet,omitempty"`
	ApiVersionSetId                  *string                                  `json:"apiVersionSetId,omitempty"`
	AuthenticationSettings           *AuthenticationSettingsContract          `json:"authenticationSettings,omitempty"`
	Contact                          *ApiContactInformation                   `json:"contact,omitempty"`
	Description                      *string                                  `json:"description,omitempty"`
	DisplayName                      *string                                  `json:"displayName,omitempty"`
	Format                           *ContentFormat                           `json:"format,omitempty"`
	IsCurrent                        *bool                                    `json:"isCurrent,omitempty"`
	IsOnline                         *bool                                    `json:"isOnline,omitempty"`
	License                          *ApiLicenseInformation                   `json:"license,omitempty"`
	Path                             string                                   `json:"path"`
	Protocols                        *[]Protocol                              `json:"protocols,omitempty"`
	ProvisioningState                *string                                  `json:"provisioningState,omitempty"`
	ServiceURL                       *string                                  `json:"serviceUrl,omitempty"`
	SourceApiId                      *string                                  `json:"sourceApiId,omitempty"`
	SubscriptionKeyParameterNames    *SubscriptionKeyParameterNamesContract   `json:"subscriptionKeyParameterNames,omitempty"`
	SubscriptionRequired             *bool                                    `json:"subscriptionRequired,omitempty"`
	TermsOfServiceURL                *string                                  `json:"termsOfServiceUrl,omitempty"`
	TranslateRequiredQueryParameters *TranslateRequiredQueryParametersConduct `json:"translateRequiredQueryParameters,omitempty"`
	Type                             *ApiType                                 `json:"type,omitempty"`
	Value                            *string                                  `json:"value,omitempty"`
	WsdlSelector                     *ApiCreateOrUpdatePropertiesWsdlSelector `json:"wsdlSelector,omitempty"`
}
