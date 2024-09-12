package teamchannelfilesfoldercontentstream

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

type DeleteTeamChannelFilesFolderContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteTeamChannelFilesFolderContentStreamOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteTeamChannelFilesFolderContentStreamOperationOptions() DeleteTeamChannelFilesFolderContentStreamOperationOptions {
	return DeleteTeamChannelFilesFolderContentStreamOperationOptions{}
}

func (o DeleteTeamChannelFilesFolderContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteTeamChannelFilesFolderContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteTeamChannelFilesFolderContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteTeamChannelFilesFolderContentStream - Delete contentStream for the navigation property filesFolder in groups.
// The content stream, if the item represents a file.
func (c TeamChannelFilesFolderContentStreamClient) DeleteTeamChannelFilesFolderContentStream(ctx context.Context, id beta.GroupIdTeamChannelId, options DeleteTeamChannelFilesFolderContentStreamOperationOptions) (result DeleteTeamChannelFilesFolderContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/filesFolder/contentStream", id.ID()),
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
