package onenotenotebooksectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteNotebookSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteNotebookSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotenotebooksectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteNotebookSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &OnenoteNotebookSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
