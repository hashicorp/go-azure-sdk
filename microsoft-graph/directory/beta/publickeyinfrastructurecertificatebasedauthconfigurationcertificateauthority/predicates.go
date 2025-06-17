package publickeyinfrastructurecertificatebasedauthconfigurationcertificateauthority

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CertificateAuthorityDetailOperationPredicate struct {
}

func (p CertificateAuthorityDetailOperationPredicate) Matches(input beta.CertificateAuthorityDetail) bool {

	return true
}
