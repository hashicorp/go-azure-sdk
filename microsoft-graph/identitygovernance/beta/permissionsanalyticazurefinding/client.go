package permissionsanalyticazurefinding

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticAzureFindingClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticAzureFindingClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticAzureFindingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticazurefinding", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticAzureFindingClient: %+v", err)
	}

	return &PermissionsAnalyticAzureFindingClient{
		Client: client,
	}, nil
}
