package impactedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactedResourceClient struct {
	Client *msgraph.Client
}

func NewImpactedResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*ImpactedResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "impactedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ImpactedResourceClient: %+v", err)
	}

	return &ImpactedResourceClient{
		Client: client,
	}, nil
}
