package directoryfeaturerolloutpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryFeatureRolloutPolicyClient struct {
	Client *msgraph.Client
}

func NewDirectoryFeatureRolloutPolicyClientWithBaseURI(api sdkEnv.Api) (*DirectoryFeatureRolloutPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryfeaturerolloutpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryFeatureRolloutPolicyClient: %+v", err)
	}

	return &DirectoryFeatureRolloutPolicyClient{
		Client: client,
	}, nil
}
