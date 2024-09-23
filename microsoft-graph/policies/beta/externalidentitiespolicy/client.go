package externalidentitiespolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalIdentitiesPolicyClient struct {
	Client *msgraph.Client
}

func NewExternalIdentitiesPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ExternalIdentitiesPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "externalidentitiespolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExternalIdentitiesPolicyClient: %+v", err)
	}

	return &ExternalIdentitiesPolicyClient{
		Client: client,
	}, nil
}
