package useronenotenotebooksectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
