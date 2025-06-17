package authenticationeventsflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AuthenticationEventsFlowOperationPredicate struct {
}

func (p AuthenticationEventsFlowOperationPredicate) Matches(input stable.AuthenticationEventsFlow) bool {

	return true
}

type ExternalUsersSelfServiceSignUpEventsFlowOperationPredicate struct {
}

func (p ExternalUsersSelfServiceSignUpEventsFlowOperationPredicate) Matches(input stable.ExternalUsersSelfServiceSignUpEventsFlow) bool {

	return true
}
