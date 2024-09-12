package pendingaccessreviewinstancecontactedreviewer

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

type DeletePendingAccessReviewInstanceContactedReviewerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePendingAccessReviewInstanceContactedReviewerOperationOptions struct {
	IfMatch *string
}

func DefaultDeletePendingAccessReviewInstanceContactedReviewerOperationOptions() DeletePendingAccessReviewInstanceContactedReviewerOperationOptions {
	return DeletePendingAccessReviewInstanceContactedReviewerOperationOptions{}
}

func (o DeletePendingAccessReviewInstanceContactedReviewerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePendingAccessReviewInstanceContactedReviewerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeletePendingAccessReviewInstanceContactedReviewerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePendingAccessReviewInstanceContactedReviewer - Delete navigation property contactedReviewers for me
func (c PendingAccessReviewInstanceContactedReviewerClient) DeletePendingAccessReviewInstanceContactedReviewer(ctx context.Context, id beta.MePendingAccessReviewInstanceIdContactedReviewerId, options DeletePendingAccessReviewInstanceContactedReviewerOperationOptions) (result DeletePendingAccessReviewInstanceContactedReviewerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
