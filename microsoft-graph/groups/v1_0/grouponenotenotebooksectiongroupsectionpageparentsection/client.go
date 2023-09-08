package grouponenotenotebooksectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteNotebookSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteNotebookSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotenotebooksectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteNotebookSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &GroupOnenoteNotebookSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
