package policypermissiongrantpolicyexclude

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyPermissionGrantPolicyExcludeClient struct {
	Client *msgraph.Client
}

func NewPolicyPermissionGrantPolicyExcludeClientWithBaseURI(api sdkEnv.Api) (*PolicyPermissionGrantPolicyExcludeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policypermissiongrantpolicyexclude", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyPermissionGrantPolicyExcludeClient: %+v", err)
	}

	return &PolicyPermissionGrantPolicyExcludeClient{
		Client: client,
	}, nil
}
