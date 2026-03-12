package csmpublishingcredentialspoliciesentityoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CsmPublishingCredentialsPoliciesEntityOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewCsmPublishingCredentialsPoliciesEntityOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*CsmPublishingCredentialsPoliciesEntityOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "csmpublishingcredentialspoliciesentityoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CsmPublishingCredentialsPoliciesEntityOperationGroupClient: %+v", err)
	}

	return &CsmPublishingCredentialsPoliciesEntityOperationGroupClient{
		Client: client,
	}, nil
}
