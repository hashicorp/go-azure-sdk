package authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient: %+v", err)
	}

	return &AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient{
		Client: client,
	}, nil
}
