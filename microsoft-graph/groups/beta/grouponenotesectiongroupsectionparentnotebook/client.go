package grouponenotesectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
