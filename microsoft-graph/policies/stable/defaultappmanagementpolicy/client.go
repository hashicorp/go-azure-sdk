package defaultappmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultAppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewDefaultAppManagementPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*DefaultAppManagementPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defaultappmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefaultAppManagementPolicyClient: %+v", err)
	}

	return &DefaultAppManagementPolicyClient{
		Client: client,
	}, nil
}
