package resourceaccessprofile

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

type AssignResourceAccessProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileAssignment
}

type AssignResourceAccessProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileAssignment
}

type AssignResourceAccessProfilesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultAssignResourceAccessProfilesOperationOptions() AssignResourceAccessProfilesOperationOptions {
	return AssignResourceAccessProfilesOperationOptions{}
}

func (o AssignResourceAccessProfilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignResourceAccessProfilesOperationOptions) ToOData() *odata.Query {
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

func (o AssignResourceAccessProfilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignResourceAccessProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignResourceAccessProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignResourceAccessProfiles - Invoke action assign
func (c ResourceAccessProfileClient) AssignResourceAccessProfiles(ctx context.Context, id beta.DeviceManagementResourceAccessProfileId, input AssignResourceAccessProfilesRequest, options AssignResourceAccessProfilesOperationOptions) (result AssignResourceAccessProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignResourceAccessProfilesCustomPager{},
		Path:          fmt.Sprintf("%s/assign", id.ID()),
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
		Values *[]beta.DeviceManagementResourceAccessProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignResourceAccessProfilesComplete retrieves all the results into a single object
func (c ResourceAccessProfileClient) AssignResourceAccessProfilesComplete(ctx context.Context, id beta.DeviceManagementResourceAccessProfileId, input AssignResourceAccessProfilesRequest, options AssignResourceAccessProfilesOperationOptions) (AssignResourceAccessProfilesCompleteResult, error) {
	return c.AssignResourceAccessProfilesCompleteMatchingPredicate(ctx, id, input, options, DeviceManagementResourceAccessProfileAssignmentOperationPredicate{})
}

// AssignResourceAccessProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileClient) AssignResourceAccessProfilesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementResourceAccessProfileId, input AssignResourceAccessProfilesRequest, options AssignResourceAccessProfilesOperationOptions, predicate DeviceManagementResourceAccessProfileAssignmentOperationPredicate) (result AssignResourceAccessProfilesCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileAssignment, 0)

	resp, err := c.AssignResourceAccessProfiles(ctx, id, input, options)
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

	result = AssignResourceAccessProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
