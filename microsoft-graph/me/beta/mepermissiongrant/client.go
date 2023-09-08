package mepermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePermissionGrantClient struct {
	Client *msgraph.Client
}

func NewMePermissionGrantClientWithBaseURI(api sdkEnv.Api) (*MePermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mepermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePermissionGrantClient: %+v", err)
	}

	return &MePermissionGrantClient{
		Client: client,
	}, nil
}
