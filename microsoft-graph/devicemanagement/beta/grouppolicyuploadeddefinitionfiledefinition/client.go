package grouppolicyuploadeddefinitionfiledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFileDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyUploadedDefinitionFileDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyUploadedDefinitionFileDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyuploadeddefinitionfiledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyUploadedDefinitionFileDefinitionClient: %+v", err)
	}

	return &GroupPolicyUploadedDefinitionFileDefinitionClient{
		Client: client,
	}, nil
}
