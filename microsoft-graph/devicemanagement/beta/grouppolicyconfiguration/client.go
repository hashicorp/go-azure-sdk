package grouppolicyconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyConfigurationClient: %+v", err)
	}

	return &GroupPolicyConfigurationClient{
		Client: client,
	}, nil
}
