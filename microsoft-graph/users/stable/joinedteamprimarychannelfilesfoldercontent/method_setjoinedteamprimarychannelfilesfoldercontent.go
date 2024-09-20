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

type SetJoinedTeamPrimaryChannelFilesFolderContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions() SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions {
	return SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions{}
}

func (o SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetJoinedTeamPrimaryChannelFilesFolderContent - Update content for the navigation property filesFolder in users. The
// content stream, if the item represents a file.
func (c JoinedTeamPrimaryChannelFilesFolderContentClient) SetJoinedTeamPrimaryChannelFilesFolderContent(ctx context.Context, id stable.UserIdJoinedTeamId, input []byte, options SetJoinedTeamPrimaryChannelFilesFolderContentOperationOptions) (result SetJoinedTeamPrimaryChannelFilesFolderContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/primaryChannel/filesFolder/content", id.ID()),
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
