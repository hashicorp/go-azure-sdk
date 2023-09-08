package grouppermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewGroupPermissionGrantClientWithBaseURI(api sdkEnv.Api) (*GroupPermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouppermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPermissionGrantClient: %+v", err)
	}

	return &GroupPermissionGrantClient{
		Client: client,
	}, nil
}
