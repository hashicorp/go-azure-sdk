package policypermissiongrantpolicyinclude

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyPermissionGrantPolicyIncludeClient struct {
	Client *msgraph.Client
}

func NewPolicyPermissionGrantPolicyIncludeClientWithBaseURI(api sdkEnv.Api) (*PolicyPermissionGrantPolicyIncludeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policypermissiongrantpolicyinclude", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyPermissionGrantPolicyIncludeClient: %+v", err)
	}

	return &PolicyPermissionGrantPolicyIncludeClient{
		Client: client,
	}, nil
}
