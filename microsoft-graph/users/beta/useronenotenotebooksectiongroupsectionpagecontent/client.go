package useronenotenotebooksectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupSectionPageContentClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
