package grouppolicyuploadeddefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyUploadedDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyUploadedDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyuploadeddefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyUploadedDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyUploadedDefinitionFileClient{
		Client: client,
	}, nil
}
