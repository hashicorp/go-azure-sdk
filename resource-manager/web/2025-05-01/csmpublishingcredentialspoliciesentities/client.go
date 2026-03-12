package csmpublishingcredentialspoliciesentities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CsmPublishingCredentialsPoliciesEntitiesClient struct {
	Client *resourcemanager.Client
}

func NewCsmPublishingCredentialsPoliciesEntitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*CsmPublishingCredentialsPoliciesEntitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "csmpublishingcredentialspoliciesentities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CsmPublishingCredentialsPoliciesEntitiesClient: %+v", err)
	}

	return &CsmPublishingCredentialsPoliciesEntitiesClient{
		Client: client,
	}, nil
}
