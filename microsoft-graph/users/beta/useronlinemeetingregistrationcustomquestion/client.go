package useronlinemeetingregistrationcustomquestion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingRegistrationCustomQuestionClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingRegistrationCustomQuestionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingregistrationcustomquestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingRegistrationCustomQuestionClient: %+v", err)
	}

	return &UserOnlineMeetingRegistrationCustomQuestionClient{
		Client: client,
	}, nil
}
