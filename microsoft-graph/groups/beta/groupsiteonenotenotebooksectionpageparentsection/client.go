package groupsiteonenotenotebooksectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteNotebookSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteNotebookSectionPageParentSectionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteNotebookSectionPageParentSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotenotebooksectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteNotebookSectionPageParentSectionClient: %+v", err)
	}

	return &GroupSiteOnenoteNotebookSectionPageParentSectionClient{
		Client: client,
	}, nil
}
