package exchangeroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewExchangeRoleAssignmentPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeRoleAssignmentPrincipalClient: %+v", err)
	}

	return &ExchangeRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
