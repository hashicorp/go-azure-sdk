package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditLogRoot struct {
	CustomSecurityAttributeAudits *[]CustomSecurityAttributeAudit `json:"customSecurityAttributeAudits,omitempty"`
	DirectoryAudits               *[]DirectoryAudit               `json:"directoryAudits,omitempty"`
	DirectoryProvisioning         *[]ProvisioningObjectSummary    `json:"directoryProvisioning,omitempty"`
	ODataType                     *string                         `json:"@odata.type,omitempty"`
	Provisioning                  *[]ProvisioningObjectSummary    `json:"provisioning,omitempty"`
	SignIns                       *[]SignIn                       `json:"signIns,omitempty"`
}
