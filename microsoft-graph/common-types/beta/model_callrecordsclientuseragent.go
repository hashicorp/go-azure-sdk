package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsClientUserAgent struct {
	ApplicationVersion     *string                                  `json:"applicationVersion,omitempty"`
	AzureADAppId           *string                                  `json:"azureADAppId,omitempty"`
	CommunicationServiceId *string                                  `json:"communicationServiceId,omitempty"`
	HeaderValue            *string                                  `json:"headerValue,omitempty"`
	ODataType              *string                                  `json:"@odata.type,omitempty"`
	Platform               *CallRecordsClientUserAgentPlatform      `json:"platform,omitempty"`
	ProductFamily          *CallRecordsClientUserAgentProductFamily `json:"productFamily,omitempty"`
}
