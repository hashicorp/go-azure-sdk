package teamprimarychannelallmember

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveTeamPrimaryChannelAllMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type RemoveTeamPrimaryChannelAllMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type RemoveTeamPrimaryChannelAllMembersOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultRemoveTeamPrimaryChannelAllMembersOperationOptions() RemoveTeamPrimaryChannelAllMembersOperationOptions {
	return RemoveTeamPrimaryChannelAllMembersOperationOptions{}
}

func (o RemoveTeamPrimaryChannelAllMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveTeamPrimaryChannelAllMembersOperationOptions) ToOData() *odata.Query {
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

func (o RemoveTeamPrimaryChannelAllMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type RemoveTeamPrimaryChannelAllMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *RemoveTeamPrimaryChannelAllMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// RemoveTeamPrimaryChannelAllMembers - Invoke action remove. Remove multiple members from a team in a single request.
// The response provides details about which memberships could and couldn't be removed.
func (c TeamPrimaryChannelAllMemberClient) RemoveTeamPrimaryChannelAllMembers(ctx context.Context, id beta.GroupId, input RemoveTeamPrimaryChannelAllMembersRequest, options RemoveTeamPrimaryChannelAllMembersOperationOptions) (result RemoveTeamPrimaryChannelAllMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &RemoveTeamPrimaryChannelAllMembersCustomPager{},
		Path:          fmt.Sprintf("%s/team/primaryChannel/allMembers/remove", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// RemoveTeamPrimaryChannelAllMembersComplete retrieves all the results into a single object
func (c TeamPrimaryChannelAllMemberClient) RemoveTeamPrimaryChannelAllMembersComplete(ctx context.Context, id beta.GroupId, input RemoveTeamPrimaryChannelAllMembersRequest, options RemoveTeamPrimaryChannelAllMembersOperationOptions) (RemoveTeamPrimaryChannelAllMembersCompleteResult, error) {
	return c.RemoveTeamPrimaryChannelAllMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// RemoveTeamPrimaryChannelAllMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelAllMemberClient) RemoveTeamPrimaryChannelAllMembersCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, input RemoveTeamPrimaryChannelAllMembersRequest, options RemoveTeamPrimaryChannelAllMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result RemoveTeamPrimaryChannelAllMembersCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.RemoveTeamPrimaryChannelAllMembers(ctx, id, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = RemoveTeamPrimaryChannelAllMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
