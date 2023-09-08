package meonlinemeetingregistrationcustomquestion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingRegistrationCustomQuestionClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingRegistrationCustomQuestionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingregistrationcustomquestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingRegistrationCustomQuestionClient: %+v", err)
	}

	return &MeOnlineMeetingRegistrationCustomQuestionClient{
		Client: client,
	}, nil
}
