package authenticationeventlistener

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AuthenticationEventListenerOperationPredicate struct {
}

func (p AuthenticationEventListenerOperationPredicate) Matches(input stable.AuthenticationEventListener) bool {

	return true
}
