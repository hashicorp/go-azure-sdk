package grouponenotesectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
