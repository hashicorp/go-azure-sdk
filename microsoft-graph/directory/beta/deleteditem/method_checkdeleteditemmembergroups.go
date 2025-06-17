package deleteditem

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

type CheckDeletedItemMemberGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type CheckDeletedItemMemberGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type CheckDeletedItemMemberGroupsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultCheckDeletedItemMemberGroupsOperationOptions() CheckDeletedItemMemberGroupsOperationOptions {
	return CheckDeletedItemMemberGroupsOperationOptions{}
}

func (o CheckDeletedItemMemberGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckDeletedItemMemberGroupsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o CheckDeletedItemMemberGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type CheckDeletedItemMemberGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *CheckDeletedItemMemberGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CheckDeletedItemMemberGroups - Invoke action checkMemberGroups. Check for membership in a specified list of group
// IDs, and return from that list the IDs of groups where a specified object is a member. The specified object can be of
// one of the following types: - user - group - service principal - organizational contact - device - directory object
// This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all
// groups provisioned in Microsoft Entra ID. Because Microsoft 365 groups cannot contain other groups, membership in a
// Microsoft 365 group is always direct.
func (c DeletedItemClient) CheckDeletedItemMemberGroups(ctx context.Context, id beta.DirectoryDeletedItemId, input CheckDeletedItemMemberGroupsRequest, options CheckDeletedItemMemberGroupsOperationOptions) (result CheckDeletedItemMemberGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &CheckDeletedItemMemberGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/checkMemberGroups", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CheckDeletedItemMemberGroupsComplete retrieves all the results into a single object
func (c DeletedItemClient) CheckDeletedItemMemberGroupsComplete(ctx context.Context, id beta.DirectoryDeletedItemId, input CheckDeletedItemMemberGroupsRequest, options CheckDeletedItemMemberGroupsOperationOptions) (result CheckDeletedItemMemberGroupsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.CheckDeletedItemMemberGroups(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = CheckDeletedItemMemberGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
