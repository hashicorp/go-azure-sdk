package informationprotectionpolicylabel

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type InformationProtectionActionOperationPredicate struct {
}

func (p InformationProtectionActionOperationPredicate) Matches(input beta.InformationProtectionAction) bool {

	return true
}

type InformationProtectionLabelOperationPredicate struct {
}

func (p InformationProtectionLabelOperationPredicate) Matches(input beta.InformationProtectionLabel) bool {

	return true
}
