package permissionsanalyticawfinding

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticAwFindingClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticAwFindingClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticAwFindingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticawfinding", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticAwFindingClient: %+v", err)
	}

	return &PermissionsAnalyticAwFindingClient{
		Client: client,
	}, nil
}
