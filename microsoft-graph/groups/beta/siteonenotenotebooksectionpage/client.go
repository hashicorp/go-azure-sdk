package siteonenotenotebooksectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteNotebookSectionPageClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteNotebookSectionPageClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteNotebookSectionPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotenotebooksectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteNotebookSectionPageClient: %+v", err)
	}

	return &SiteOnenoteNotebookSectionPageClient{
		Client: client,
	}, nil
}
