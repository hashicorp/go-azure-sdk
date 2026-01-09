package pendingaccessreviewinstancedefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPendingAccessReviewInstanceDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewScheduleDefinition
}

type GetPendingAccessReviewInstanceDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPendingAccessReviewInstanceDefinitionOperationOptions() GetPendingAccessReviewInstanceDefinitionOperationOptions {
	return GetPendingAccessReviewInstanceDefinitionOperationOptions{}
}

func (o GetPendingAccessReviewInstanceDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPendingAccessReviewInstanceDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPendingAccessReviewInstanceDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPendingAccessReviewInstanceDefinition - Get definition from me. There's exactly one accessReviewScheduleDefinition
// associated with each instance. It's the parent schedule for the instance, where instances are created for each
// recurrence of a review definition and each group selected to review by the definition.
func (c PendingAccessReviewInstanceDefinitionClient) GetPendingAccessReviewInstanceDefinition(ctx context.Context, id beta.MePendingAccessReviewInstanceId, options GetPendingAccessReviewInstanceDefinitionOperationOptions) (result GetPendingAccessReviewInstanceDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/definition", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.AccessReviewScheduleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
