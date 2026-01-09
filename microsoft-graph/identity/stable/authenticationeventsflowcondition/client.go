package authenticationeventsflowcondition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlowConditionClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventsFlowConditionClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventsFlowConditionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventsflowcondition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventsFlowConditionClient: %+v", err)
	}

	return &AuthenticationEventsFlowConditionClient{
		Client: client,
	}, nil
}
