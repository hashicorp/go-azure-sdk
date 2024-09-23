package joinedteamchannelfilesfoldercontent

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

type DeleteJoinedTeamChannelFilesFolderContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteJoinedTeamChannelFilesFolderContentOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteJoinedTeamChannelFilesFolderContentOperationOptions() DeleteJoinedTeamChannelFilesFolderContentOperationOptions {
	return DeleteJoinedTeamChannelFilesFolderContentOperationOptions{}
}

func (o DeleteJoinedTeamChannelFilesFolderContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteJoinedTeamChannelFilesFolderContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteJoinedTeamChannelFilesFolderContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteJoinedTeamChannelFilesFolderContent - Delete content for the navigation property filesFolder in users. The
// content stream, if the item represents a file.
func (c JoinedTeamChannelFilesFolderContentClient) DeleteJoinedTeamChannelFilesFolderContent(ctx context.Context, id stable.UserIdJoinedTeamIdChannelId, options DeleteJoinedTeamChannelFilesFolderContentOperationOptions) (result DeleteJoinedTeamChannelFilesFolderContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/filesFolder/content", id.ID()),
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
