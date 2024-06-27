package policy

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAgreementPolicies struct {
	AccountOwnerViewCharges    *EnrollmentAccountOwnerViewCharges    `json:"accountOwnerViewCharges,omitempty"`
	AuthenticationType         *EnrollmentAuthLevelState             `json:"authenticationType,omitempty"`
	DepartmentAdminViewCharges *EnrollmentDepartmentAdminViewCharges `json:"departmentAdminViewCharges,omitempty"`
}
