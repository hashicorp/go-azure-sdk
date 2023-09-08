package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingSingleSignOn struct {
	KerberosSignOnSettings *KerberosSignOnSettings                           `json:"kerberosSignOnSettings,omitempty"`
	ODataType              *string                                           `json:"@odata.type,omitempty"`
	SingleSignOnMode       *OnPremisesPublishingSingleSignOnSingleSignOnMode `json:"singleSignOnMode,omitempty"`
}
