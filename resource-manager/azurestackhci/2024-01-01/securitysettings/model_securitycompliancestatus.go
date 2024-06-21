package securitysettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityComplianceStatus struct {
	DataAtRestEncrypted    *ComplianceStatus `json:"dataAtRestEncrypted,omitempty"`
	DataInTransitProtected *ComplianceStatus `json:"dataInTransitProtected,omitempty"`
	LastUpdated            *string           `json:"lastUpdated,omitempty"`
	SecuredCoreCompliance  *ComplianceStatus `json:"securedCoreCompliance,omitempty"`
	WdacCompliance         *ComplianceStatus `json:"wdacCompliance,omitempty"`
}
