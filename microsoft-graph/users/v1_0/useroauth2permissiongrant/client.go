package useroauth2permissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOauth2PermissionGrantClient struct {
	Client *msgraph.Client
}

func NewUserOauth2PermissionGrantClientWithBaseURI(api sdkEnv.Api) (*UserOauth2PermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroauth2permissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOauth2PermissionGrantClient: %+v", err)
	}

	return &UserOauth2PermissionGrantClient{
		Client: client,
	}, nil
}
