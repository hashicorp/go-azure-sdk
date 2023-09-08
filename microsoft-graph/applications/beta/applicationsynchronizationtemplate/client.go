package applicationsynchronizationtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationTemplateClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationTemplateClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationTemplateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronizationtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationTemplateClient: %+v", err)
	}

	return &ApplicationSynchronizationTemplateClient{
		Client: client,
	}, nil
}
