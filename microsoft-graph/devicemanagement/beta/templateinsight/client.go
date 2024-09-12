package templateinsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateInsightClient struct {
	Client *msgraph.Client
}

func NewTemplateInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templateinsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateInsightClient: %+v", err)
	}

	return &TemplateInsightClient{
		Client: client,
	}, nil
}
