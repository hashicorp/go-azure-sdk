package joinedteamprimarychannelfilesfoldercontent

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

type DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions() DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions {
	return DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions{}
}

func (o DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteJoinedTeamPrimaryChannelFilesFolderContent - Delete content for the navigation property filesFolder in users.
// The content stream, if the item represents a file.
func (c JoinedTeamPrimaryChannelFilesFolderContentClient) DeleteJoinedTeamPrimaryChannelFilesFolderContent(ctx context.Context, id stable.UserIdJoinedTeamId, options DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) (result DeleteJoinedTeamPrimaryChannelFilesFolderContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/primaryChannel/filesFolder/content", id.ID()),
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
