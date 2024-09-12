package authenticationeventsflow

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlowClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventsFlowClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventsFlowClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventsflow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventsFlowClient: %+v", err)
	}

	return &AuthenticationEventsFlowClient{
		Client: client,
	}, nil
}
