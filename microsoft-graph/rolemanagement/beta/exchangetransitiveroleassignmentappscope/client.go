package exchangetransitiveroleassignmentappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeTransitiveRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewExchangeTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeTransitiveRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangetransitiveroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeTransitiveRoleAssignmentAppScopeClient: %+v", err)
	}

	return &ExchangeTransitiveRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}
