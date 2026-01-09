package onenotesectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
