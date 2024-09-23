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

type CreateTeamCompleteMigrationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamCompleteMigrationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamCompleteMigrationOperationOptions() CreateTeamCompleteMigrationOperationOptions {
	return CreateTeamCompleteMigrationOperationOptions{}
}

func (o CreateTeamCompleteMigrationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamCompleteMigrationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamCompleteMigrationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamCompleteMigration - Invoke action completeMigration. Complete the message migration process by removing
// migration mode from a team. Migration mode is a special state where certain operations are barred, like message POST
// and membership operations during the data migration process. After a completeMigration request is made, you can't
// import additional messages into the team. You can add members to the team after the request returns a successful
// response.
func (c TeamClient) CreateTeamCompleteMigration(ctx context.Context, id stable.GroupId, options CreateTeamCompleteMigrationOperationOptions) (result CreateTeamCompleteMigrationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/completeMigration", id.ID()),
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

	return
}
