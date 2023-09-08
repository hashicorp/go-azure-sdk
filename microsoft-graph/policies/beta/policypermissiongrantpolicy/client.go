package policypermissiongrantpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyPermissionGrantPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyPermissionGrantPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyPermissionGrantPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policypermissiongrantpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyPermissionGrantPolicyClient: %+v", err)
	}

	return &PolicyPermissionGrantPolicyClient{
		Client: client,
	}, nil
}
