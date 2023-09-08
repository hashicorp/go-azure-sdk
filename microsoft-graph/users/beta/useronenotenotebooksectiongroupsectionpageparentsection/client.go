package useronenotenotebooksectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteNotebookSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteNotebookSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotenotebooksectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteNotebookSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &UserOnenoteNotebookSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
