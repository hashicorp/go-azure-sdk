package cloudpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetCloudPCReviewStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetCloudPCReviewStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetCloudPCReviewStatusOperationOptions() SetCloudPCReviewStatusOperationOptions {
	return SetCloudPCReviewStatusOperationOptions{}
}

func (o SetCloudPCReviewStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetCloudPCReviewStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetCloudPCReviewStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetCloudPCReviewStatus - Invoke action setReviewStatus. Set the review status of a specific Cloud PC device using the
// Cloud PC ID. Use this API to set the review status of a Cloud PC to in review if you consider a Cloud PC suspicious.
// After the review is completed, use this API again to set the Cloud PC back to a normal state.
func (c CloudPCClient) SetCloudPCReviewStatus(ctx context.Context, id beta.MeCloudPCId, input SetCloudPCReviewStatusRequest, options SetCloudPCReviewStatusOperationOptions) (result SetCloudPCReviewStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/setReviewStatus", id.ID()),
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
