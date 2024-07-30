package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightSummary struct {
	ActiveUsers               *int64  `json:"activeUsers,omitempty"`
	AppId                     *string `json:"appId,omitempty"`
	AuthenticationCompletions *int64  `json:"authenticationCompletions,omitempty"`
	AuthenticationRequests    *int64  `json:"authenticationRequests,omitempty"`
	FactDate                  *string `json:"factDate,omitempty"`
	Id                        *string `json:"id,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	Os                        *string `json:"os,omitempty"`
	SecurityTextCompletions   *int64  `json:"securityTextCompletions,omitempty"`
	SecurityTextRequests      *int64  `json:"securityTextRequests,omitempty"`
	SecurityVoiceCompletions  *int64  `json:"securityVoiceCompletions,omitempty"`
	SecurityVoiceRequests     *int64  `json:"securityVoiceRequests,omitempty"`
}
