package joinedteamprimarychannel

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

type CreateJoinedTeamPrimaryChannelCompleteMigrationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateJoinedTeamPrimaryChannelCompleteMigration - Invoke action completeMigration. Complete the message migration
// process by removing migration mode from a channel in a team. Migration mode is a special state that prevents certain
// operations, like sending messages and adding members, during the data migration process. After a completeMigration
// request is made, you can't import additional messages into the team. You can add members to the team after the
// request returns a successful response.
func (c JoinedTeamPrimaryChannelClient) CreateJoinedTeamPrimaryChannelCompleteMigration(ctx context.Context, id stable.UserIdJoinedTeamId) (result CreateJoinedTeamPrimaryChannelCompleteMigrationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/primaryChannel/completeMigration", id.ID()),
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
