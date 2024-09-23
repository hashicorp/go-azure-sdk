package onenotesectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
