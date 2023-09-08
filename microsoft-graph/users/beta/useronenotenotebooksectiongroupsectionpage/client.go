package useronenotenotebooksectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupSectionPageClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
