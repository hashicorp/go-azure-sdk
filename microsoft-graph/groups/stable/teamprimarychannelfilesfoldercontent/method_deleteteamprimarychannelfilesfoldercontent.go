package teamprimarychannelfilesfoldercontent

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

type DeleteTeamPrimaryChannelFilesFolderContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteTeamPrimaryChannelFilesFolderContentOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteTeamPrimaryChannelFilesFolderContentOperationOptions() DeleteTeamPrimaryChannelFilesFolderContentOperationOptions {
	return DeleteTeamPrimaryChannelFilesFolderContentOperationOptions{}
}

func (o DeleteTeamPrimaryChannelFilesFolderContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteTeamPrimaryChannelFilesFolderContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteTeamPrimaryChannelFilesFolderContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteTeamPrimaryChannelFilesFolderContent - Delete content for the navigation property filesFolder in groups. The
// content stream, if the item represents a file.
func (c TeamPrimaryChannelFilesFolderContentClient) DeleteTeamPrimaryChannelFilesFolderContent(ctx context.Context, id stable.GroupId, options DeleteTeamPrimaryChannelFilesFolderContentOperationOptions) (result DeleteTeamPrimaryChannelFilesFolderContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/primaryChannel/filesFolder/content", id.ID()),
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
