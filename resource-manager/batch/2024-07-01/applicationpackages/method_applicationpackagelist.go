package applicationpackages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPackageListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationPackage
}

type ApplicationPackageListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationPackage
}

type ApplicationPackageListOperationOptions struct {
	Maxresults *int64
}

func DefaultApplicationPackageListOperationOptions() ApplicationPackageListOperationOptions {
	return ApplicationPackageListOperationOptions{}
}

func (o ApplicationPackageListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ApplicationPackageListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ApplicationPackageListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	return &out
}

type ApplicationPackageListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ApplicationPackageListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ApplicationPackageList ...
func (c ApplicationPackagesClient) ApplicationPackageList(ctx context.Context, id ApplicationId, options ApplicationPackageListOperationOptions) (result ApplicationPackageListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ApplicationPackageListCustomPager{},
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]ApplicationPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ApplicationPackageListComplete retrieves all the results into a single object
func (c ApplicationPackagesClient) ApplicationPackageListComplete(ctx context.Context, id ApplicationId, options ApplicationPackageListOperationOptions) (ApplicationPackageListCompleteResult, error) {
	return c.ApplicationPackageListCompleteMatchingPredicate(ctx, id, options, ApplicationPackageOperationPredicate{})
}

// ApplicationPackageListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationPackagesClient) ApplicationPackageListCompleteMatchingPredicate(ctx context.Context, id ApplicationId, options ApplicationPackageListOperationOptions, predicate ApplicationPackageOperationPredicate) (result ApplicationPackageListCompleteResult, err error) {
	items := make([]ApplicationPackage, 0)

	resp, err := c.ApplicationPackageList(ctx, id, options)
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

	result = ApplicationPackageListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
