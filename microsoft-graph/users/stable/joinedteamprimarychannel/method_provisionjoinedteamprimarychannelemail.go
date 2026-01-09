package joinedteamprimarychannel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionJoinedTeamPrimaryChannelEmailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ProvisionChannelEmailResult
}

type ProvisionJoinedTeamPrimaryChannelEmailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultProvisionJoinedTeamPrimaryChannelEmailOperationOptions() ProvisionJoinedTeamPrimaryChannelEmailOperationOptions {
	return ProvisionJoinedTeamPrimaryChannelEmailOperationOptions{}
}

func (o ProvisionJoinedTeamPrimaryChannelEmailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProvisionJoinedTeamPrimaryChannelEmailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ProvisionJoinedTeamPrimaryChannelEmailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ProvisionJoinedTeamPrimaryChannelEmail - Invoke action provisionEmail. Provision an email address for a channel.
// Microsoft Teams doesn't automatically provision an email address for a channel by default. To have Teams provision an
// email address, you can call provisionEmail, or through the Teams user interface, select Get email address, which
// triggers Teams to generate an email address if it has not already provisioned one. To remove the email address of a
// channel, use the removeEmail method.
func (c JoinedTeamPrimaryChannelClient) ProvisionJoinedTeamPrimaryChannelEmail(ctx context.Context, id stable.UserIdJoinedTeamId, options ProvisionJoinedTeamPrimaryChannelEmailOperationOptions) (result ProvisionJoinedTeamPrimaryChannelEmailOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/primaryChannel/provisionEmail", id.ID()),
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

	var model stable.ProvisionChannelEmailResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
