package exchangeroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewExchangeRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeRoleAssignmentClient: %+v", err)
	}

	return &ExchangeRoleAssignmentClient{
		Client: client,
	}, nil
}
