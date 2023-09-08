package meoauth2permissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOauth2PermissionGrantClient struct {
	Client *msgraph.Client
}

func NewMeOauth2PermissionGrantClientWithBaseURI(api sdkEnv.Api) (*MeOauth2PermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoauth2permissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOauth2PermissionGrantClient: %+v", err)
	}

	return &MeOauth2PermissionGrantClient{
		Client: client,
	}, nil
}
