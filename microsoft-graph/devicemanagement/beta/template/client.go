package template

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateClient struct {
	Client *msgraph.Client
}

func NewTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "template", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateClient: %+v", err)
	}

	return &TemplateClient{
		Client: client,
	}, nil
}
