package useronenotesectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionParentNotebookClient: %+v", err)
	}

	return &UserOnenoteSectionParentNotebookClient{
		Client: client,
	}, nil
}
