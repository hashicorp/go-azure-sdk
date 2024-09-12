package managedappregistration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ManagedAppRegistrationOperationPredicate struct {
}

func (p ManagedAppRegistrationOperationPredicate) Matches(input beta.ManagedAppRegistration) bool {

	return true
}
