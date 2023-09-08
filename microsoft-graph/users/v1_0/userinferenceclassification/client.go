package userinferenceclassification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInferenceClassificationClient struct {
	Client *msgraph.Client
}

func NewUserInferenceClassificationClientWithBaseURI(api sdkEnv.Api) (*UserInferenceClassificationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinferenceclassification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInferenceClassificationClient: %+v", err)
	}

	return &UserInferenceClassificationClient{
		Client: client,
	}, nil
}
