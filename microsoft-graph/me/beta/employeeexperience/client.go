package employeeexperience

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperienceClient struct {
	Client *msgraph.Client
}

func NewEmployeeExperienceClientWithBaseURI(sdkApi sdkEnv.Api) (*EmployeeExperienceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "employeeexperience", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmployeeExperienceClient: %+v", err)
	}

	return &EmployeeExperienceClient{
		Client: client,
	}, nil
}
