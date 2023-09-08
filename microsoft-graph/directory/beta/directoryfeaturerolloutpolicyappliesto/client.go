package directoryfeaturerolloutpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryFeatureRolloutPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewDirectoryFeatureRolloutPolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*DirectoryFeatureRolloutPolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryfeaturerolloutpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryFeatureRolloutPolicyAppliesToClient: %+v", err)
	}

	return &DirectoryFeatureRolloutPolicyAppliesToClient{
		Client: client,
	}, nil
}
