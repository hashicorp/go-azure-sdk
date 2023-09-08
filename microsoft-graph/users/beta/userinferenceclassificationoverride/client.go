package userinferenceclassificationoverride

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInferenceClassificationOverrideClient struct {
	Client *msgraph.Client
}

func NewUserInferenceClassificationOverrideClientWithBaseURI(api sdkEnv.Api) (*UserInferenceClassificationOverrideClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinferenceclassificationoverride", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInferenceClassificationOverrideClient: %+v", err)
	}

	return &UserInferenceClassificationOverrideClient{
		Client: client,
	}, nil
}
