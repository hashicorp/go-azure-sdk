package authenticationmethoduserregistrationdetail

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type UserRegistrationDetailsOperationPredicate struct {
}

func (p UserRegistrationDetailsOperationPredicate) Matches(input beta.UserRegistrationDetails) bool {

	return true
}
