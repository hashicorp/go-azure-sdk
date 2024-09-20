package teamchannelmember

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddTeamChannelMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type AddTeamChannelMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type AddTeamChannelMembersOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultAddTeamChannelMembersOperationOptions() AddTeamChannelMembersOperationOptions {
	return AddTeamChannelMembersOperationOptions{}
}

func (o AddTeamChannelMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddTeamChannelMembersOperationOptions) ToOData() *odata.Query {
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

func (o AddTeamChannelMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AddTeamChannelMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AddTeamChannelMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AddTeamChannelMembers - Invoke action add. Add multiple members in a single request to a team. The response provides
// details about which memberships could and couldn't be created.
func (c TeamChannelMemberClient) AddTeamChannelMembers(ctx context.Context, id beta.GroupIdTeamChannelId, input AddTeamChannelMembersRequest, options AddTeamChannelMembersOperationOptions) (result AddTeamChannelMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AddTeamChannelMembersCustomPager{},
		Path:          fmt.Sprintf("%s/members/add", id.ID()),
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

// AddTeamChannelMembersComplete retrieves all the results into a single object
func (c TeamChannelMemberClient) AddTeamChannelMembersComplete(ctx context.Context, id beta.GroupIdTeamChannelId, input AddTeamChannelMembersRequest, options AddTeamChannelMembersOperationOptions) (AddTeamChannelMembersCompleteResult, error) {
	return c.AddTeamChannelMembersCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// AddTeamChannelMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelMemberClient) AddTeamChannelMembersCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelId, input AddTeamChannelMembersRequest, options AddTeamChannelMembersOperationOptions, predicate ActionResultPartOperationPredicate) (result AddTeamChannelMembersCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.AddTeamChannelMembers(ctx, id, input, options)
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

	result = AddTeamChannelMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
