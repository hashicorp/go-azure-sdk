package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAuditEvent struct {
	Activity          *string `json:"activity,omitempty"`
	ActivityDateTime  *string `json:"activityDateTime,omitempty"`
	ActivityId        *string `json:"activityId,omitempty"`
	Category          *string `json:"category,omitempty"`
	HttpVerb          *string `json:"httpVerb,omitempty"`
	Id                *string `json:"id,omitempty"`
	InitiatedByAppId  *string `json:"initiatedByAppId,omitempty"`
	InitiatedByUpn    *string `json:"initiatedByUpn,omitempty"`
	InitiatedByUserId *string `json:"initiatedByUserId,omitempty"`
	IpAddress         *string `json:"ipAddress,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	RequestBody       *string `json:"requestBody,omitempty"`
	RequestUrl        *string `json:"requestUrl,omitempty"`
	TenantIds         *string `json:"tenantIds,omitempty"`
	TenantNames       *string `json:"tenantNames,omitempty"`
}
