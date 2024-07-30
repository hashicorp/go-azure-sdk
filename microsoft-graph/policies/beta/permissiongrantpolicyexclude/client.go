package permissiongrantpolicyexclude

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantPolicyExcludeClient struct {
	Client *msgraph.Client
}

func NewPermissionGrantPolicyExcludeClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionGrantPolicyExcludeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissiongrantpolicyexclude", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionGrantPolicyExcludeClient: %+v", err)
	}

	return &PermissionGrantPolicyExcludeClient{
		Client: client,
	}, nil
}
