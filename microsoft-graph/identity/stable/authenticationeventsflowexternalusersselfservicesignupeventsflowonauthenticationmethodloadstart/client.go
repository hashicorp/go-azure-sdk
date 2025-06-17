package authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventsflowexternalusersselfservicesignupeventsflowonauthenticationmethodloadstart", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient: %+v", err)
	}

	return &AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAuthenticationMethodLoadStartClient{
		Client: client,
	}, nil
}
