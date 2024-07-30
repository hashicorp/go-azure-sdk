package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNowConnection struct {
	AuthenticationMethod       *ServiceNowAuthenticationMethod                 `json:"authenticationMethod,omitempty"`
	CreatedDateTime            *string                                         `json:"createdDateTime,omitempty"`
	Id                         *string                                         `json:"id,omitempty"`
	IncidentApiUrl             *string                                         `json:"incidentApiUrl,omitempty"`
	InstanceUrl                *string                                         `json:"instanceUrl,omitempty"`
	LastModifiedDateTime       *string                                         `json:"lastModifiedDateTime,omitempty"`
	LastQueriedDateTime        *string                                         `json:"lastQueriedDateTime,omitempty"`
	ODataType                  *string                                         `json:"@odata.type,omitempty"`
	ServiceNowConnectionStatus *ServiceNowConnectionServiceNowConnectionStatus `json:"serviceNowConnectionStatus,omitempty"`
}
