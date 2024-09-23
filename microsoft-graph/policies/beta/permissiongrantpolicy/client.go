package permissiongrantpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantPolicyClient struct {
	Client *msgraph.Client
}

func NewPermissionGrantPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionGrantPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissiongrantpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionGrantPolicyClient: %+v", err)
	}

	return &PermissionGrantPolicyClient{
		Client: client,
	}, nil
}
