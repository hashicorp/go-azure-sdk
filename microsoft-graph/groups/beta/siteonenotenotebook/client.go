package siteonenotenotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookClient: %+v", err)
	}

	return &SiteOnenoteNotebookClient{
		Client: client,
	}, nil
}
