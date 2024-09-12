package exchangetransitiveroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeTransitiveRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewExchangeTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeTransitiveRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangetransitiveroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeTransitiveRoleAssignmentPrincipalClient: %+v", err)
	}

	return &ExchangeTransitiveRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
