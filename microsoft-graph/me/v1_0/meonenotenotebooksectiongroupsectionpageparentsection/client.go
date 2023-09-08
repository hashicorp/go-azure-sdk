package meonenotenotebooksectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteNotebookSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteNotebookSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotenotebooksectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteNotebookSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &MeOnenoteNotebookSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
