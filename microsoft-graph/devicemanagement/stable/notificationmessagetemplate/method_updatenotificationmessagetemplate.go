package notificationmessagetemplate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateNotificationMessageTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateNotificationMessageTemplateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateNotificationMessageTemplateOperationOptions() UpdateNotificationMessageTemplateOperationOptions {
	return UpdateNotificationMessageTemplateOperationOptions{}
}

func (o UpdateNotificationMessageTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateNotificationMessageTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateNotificationMessageTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateNotificationMessageTemplate - Update notificationMessageTemplate. Update the properties of a
// notificationMessageTemplate object.
func (c NotificationMessageTemplateClient) UpdateNotificationMessageTemplate(ctx context.Context, id stable.DeviceManagementNotificationMessageTemplateId, input stable.NotificationMessageTemplate, options UpdateNotificationMessageTemplateOperationOptions) (result UpdateNotificationMessageTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
