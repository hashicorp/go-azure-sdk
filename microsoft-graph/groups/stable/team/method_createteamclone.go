package team

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamCloneOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamCloneOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamCloneOperationOptions() CreateTeamCloneOperationOptions {
	return CreateTeamCloneOperationOptions{}
}

func (o CreateTeamCloneOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamCloneOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamCloneOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamClone - Invoke action clone. Create a copy of a team. This operation also creates a copy of the
// corresponding group. You can specify which parts of the team to clone: When tabs are cloned, they aren't configured.
// The tabs are displayed on the tab bar in Microsoft Teams, and the first time a user opens them, they must go through
// the configuration screen. If the user who opens the tab doesn't have permission to configure apps, they see a message
// that says that the tab isn't configured. Cloning is a long-running operation. After the POST clone returns, you need
// to GET the operation returned by the Location: header to see if it's running, succeeded, or failed. You should
// continue to GET until the status isn't running. The recommended delay between GETs is 5 seconds.
func (c TeamClient) CreateTeamClone(ctx context.Context, id stable.GroupId, input CreateTeamCloneRequest, options CreateTeamCloneOperationOptions) (result CreateTeamCloneOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/team/clone", id.ID()),
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
