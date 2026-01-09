package exchangeroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewExchangeRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &ExchangeRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}
