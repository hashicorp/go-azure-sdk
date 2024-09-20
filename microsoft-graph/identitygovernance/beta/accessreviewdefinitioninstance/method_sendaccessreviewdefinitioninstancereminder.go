package accessreviewdefinitioninstance

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

type SendAccessReviewDefinitionInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendAccessReviewDefinitionInstanceReminderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSendAccessReviewDefinitionInstanceReminderOperationOptions() SendAccessReviewDefinitionInstanceReminderOperationOptions {
	return SendAccessReviewDefinitionInstanceReminderOperationOptions{}
}

func (o SendAccessReviewDefinitionInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendAccessReviewDefinitionInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendAccessReviewDefinitionInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendAccessReviewDefinitionInstanceReminder - Invoke action sendReminder. Send a reminder to the reviewers of a
// currently active accessReviewInstance.
func (c AccessReviewDefinitionInstanceClient) SendAccessReviewDefinitionInstanceReminder(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options SendAccessReviewDefinitionInstanceReminderOperationOptions) (result SendAccessReviewDefinitionInstanceReminderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sendReminder", id.ID()),
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
