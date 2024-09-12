package grouppolicyuploadeddefinitionfilegrouppolicyoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyUploadedDefinitionFileGroupPolicyOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyuploadeddefinitionfilegrouppolicyoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient: %+v", err)
	}

	return &GroupPolicyUploadedDefinitionFileGroupPolicyOperationClient{
		Client: client,
	}, nil
}
