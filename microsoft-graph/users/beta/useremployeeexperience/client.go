package useremployeeexperience

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEmployeeExperienceClient struct {
	Client *msgraph.Client
}

func NewUserEmployeeExperienceClientWithBaseURI(api sdkEnv.Api) (*UserEmployeeExperienceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useremployeeexperience", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEmployeeExperienceClient: %+v", err)
	}

	return &UserEmployeeExperienceClient{
		Client: client,
	}, nil
}
