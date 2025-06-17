package authenticationeventsflowexternalusersselfservicesignupeventsflowcondition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventsflowexternalusersselfservicesignupeventsflowcondition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient: %+v", err)
	}

	return &AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowConditionClient{
		Client: client,
	}, nil
}
