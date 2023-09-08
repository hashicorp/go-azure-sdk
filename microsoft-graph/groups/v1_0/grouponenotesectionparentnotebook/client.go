package grouponenotesectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionParentNotebookClient: %+v", err)
	}

	return &GroupOnenoteSectionParentNotebookClient{
		Client: client,
	}, nil
}
