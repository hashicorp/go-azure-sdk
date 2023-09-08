package usercalendargroupcalendareventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventInstanceAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventInstanceAttachmentClient{
		Client: client,
	}, nil
}
