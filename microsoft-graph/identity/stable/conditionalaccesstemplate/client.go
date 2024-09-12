package conditionalaccesstemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessTemplateClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesstemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessTemplateClient: %+v", err)
	}

	return &ConditionalAccessTemplateClient{
		Client: client,
	}, nil
}
