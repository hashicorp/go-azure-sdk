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

type GetJoinedTeamChannelFilesFolderContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetJoinedTeamChannelFilesFolderContentOperationOptions struct {
	Format    *odata.Format
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetJoinedTeamChannelFilesFolderContentOperationOptions() GetJoinedTeamChannelFilesFolderContentOperationOptions {
	return GetJoinedTeamChannelFilesFolderContentOperationOptions{}
}

func (o GetJoinedTeamChannelFilesFolderContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamChannelFilesFolderContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Format != nil {
		out.Format = *o.Format
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetJoinedTeamChannelFilesFolderContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamChannelFilesFolderContent - Get content for the navigation property filesFolder from me. The content
// stream, if the item represents a file.
func (c JoinedTeamChannelFilesFolderContentClient) GetJoinedTeamChannelFilesFolderContent(ctx context.Context, id stable.MeJoinedTeamIdChannelId, options GetJoinedTeamChannelFilesFolderContentOperationOptions) (result GetJoinedTeamChannelFilesFolderContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
