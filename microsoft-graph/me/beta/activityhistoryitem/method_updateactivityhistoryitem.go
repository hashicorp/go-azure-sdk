package activityhistoryitem

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateActivityHistoryItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateActivityHistoryItemOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateActivityHistoryItemOperationOptions() UpdateActivityHistoryItemOperationOptions {
	return UpdateActivityHistoryItemOperationOptions{}
}

func (o UpdateActivityHistoryItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateActivityHistoryItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateActivityHistoryItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateActivityHistoryItem - Delete an activityHistoryItem. Delete an existing history item for an existing user
// activity.
func (c ActivityHistoryItemClient) UpdateActivityHistoryItem(ctx context.Context, id beta.MeActivityIdHistoryItemId, input beta.ActivityHistoryItem, options UpdateActivityHistoryItemOperationOptions) (result UpdateActivityHistoryItemOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
