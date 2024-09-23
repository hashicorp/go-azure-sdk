package synchronizationtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTemplateClient struct {
	Client *msgraph.Client
}

func NewSynchronizationTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*SynchronizationTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "synchronizationtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynchronizationTemplateClient: %+v", err)
	}

	return &SynchronizationTemplateClient{
		Client: client,
	}, nil
}
