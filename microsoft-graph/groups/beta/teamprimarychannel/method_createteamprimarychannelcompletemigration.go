package teamprimarychannel

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

type CreateTeamPrimaryChannelCompleteMigrationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamPrimaryChannelCompleteMigrationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamPrimaryChannelCompleteMigrationOperationOptions() CreateTeamPrimaryChannelCompleteMigrationOperationOptions {
	return CreateTeamPrimaryChannelCompleteMigrationOperationOptions{}
}

func (o CreateTeamPrimaryChannelCompleteMigrationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamPrimaryChannelCompleteMigrationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamPrimaryChannelCompleteMigrationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamPrimaryChannelCompleteMigration - Invoke action completeMigration. Complete the message migration process
// by removing migration mode from a channel in a team. Migration mode is a special state that prevents certain
// operations, like sending messages and adding members, during the data migration process. After a completeMigration
// request is made, you can't import more messages into the team. You can add members to the team after the request
// returns a successful response.
func (c TeamPrimaryChannelClient) CreateTeamPrimaryChannelCompleteMigration(ctx context.Context, id beta.GroupId, options CreateTeamPrimaryChannelCompleteMigrationOperationOptions) (result CreateTeamPrimaryChannelCompleteMigrationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/primaryChannel/completeMigration", id.ID()),
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
