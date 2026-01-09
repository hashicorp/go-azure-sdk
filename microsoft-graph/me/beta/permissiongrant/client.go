package permissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantClient struct {
	Client *msgraph.Client
}

func NewPermissionGrantClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionGrantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionGrantClient: %+v", err)
	}

	return &PermissionGrantClient{
		Client: client,
	}, nil
}
