package securityinformationprotectionsensitivitylabel

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type SecurityInformationProtectionActionOperationPredicate struct {
}

func (p SecurityInformationProtectionActionOperationPredicate) Matches(input beta.SecurityInformationProtectionAction) bool {

	return true
}

type SecuritySensitivityLabelOperationPredicate struct {
}

func (p SecuritySensitivityLabelOperationPredicate) Matches(input beta.SecuritySensitivityLabel) bool {

	return true
}
