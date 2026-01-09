package exchangeroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewExchangeRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &ExchangeRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
