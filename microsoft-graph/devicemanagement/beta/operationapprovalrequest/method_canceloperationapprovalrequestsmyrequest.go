package operationapprovalrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelOperationApprovalRequestsMyRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelOperationApprovalRequestsMyRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelOperationApprovalRequestsMyRequestOperationOptions() CancelOperationApprovalRequestsMyRequestOperationOptions {
	return CancelOperationApprovalRequestsMyRequestOperationOptions{}
}

func (o CancelOperationApprovalRequestsMyRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelOperationApprovalRequestsMyRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelOperationApprovalRequestsMyRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelOperationApprovalRequestsMyRequest - Invoke action cancelMyRequest
func (c OperationApprovalRequestClient) CancelOperationApprovalRequestsMyRequest(ctx context.Context, input CancelOperationApprovalRequestsMyRequestRequest, options CancelOperationApprovalRequestsMyRequestOperationOptions) (result CancelOperationApprovalRequestsMyRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/operationApprovalRequests/cancelMyRequest",
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
