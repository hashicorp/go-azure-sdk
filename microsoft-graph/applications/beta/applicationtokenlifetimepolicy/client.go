package applicationtokenlifetimepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTokenLifetimePolicyClient struct {
	Client *msgraph.Client
}

func NewApplicationTokenLifetimePolicyClientWithBaseURI(api sdkEnv.Api) (*ApplicationTokenLifetimePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationtokenlifetimepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationTokenLifetimePolicyClient: %+v", err)
	}

	return &ApplicationTokenLifetimePolicyClient{
		Client: client,
	}, nil
}
