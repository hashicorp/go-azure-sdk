package pendingaccessreviewinstance

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

type StopPendingAccessReviewInstanceApplyDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StopPendingAccessReviewInstanceApplyDecisionsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultStopPendingAccessReviewInstanceApplyDecisionsOperationOptions() StopPendingAccessReviewInstanceApplyDecisionsOperationOptions {
	return StopPendingAccessReviewInstanceApplyDecisionsOperationOptions{}
}

func (o StopPendingAccessReviewInstanceApplyDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StopPendingAccessReviewInstanceApplyDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StopPendingAccessReviewInstanceApplyDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StopPendingAccessReviewInstanceApplyDecisions - Invoke action stopApplyDecisions
func (c PendingAccessReviewInstanceClient) StopPendingAccessReviewInstanceApplyDecisions(ctx context.Context, id beta.MePendingAccessReviewInstanceId, options StopPendingAccessReviewInstanceApplyDecisionsOperationOptions) (result StopPendingAccessReviewInstanceApplyDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/stopApplyDecisions", id.ID()),
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
