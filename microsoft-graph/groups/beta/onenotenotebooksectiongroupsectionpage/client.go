package onenotenotebooksectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupSectionPageClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
