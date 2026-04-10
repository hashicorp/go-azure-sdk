package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnnotationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Annotation
}

type AnnotationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Annotation
}

type AnnotationsListOperationOptions struct {
	End   *string
	Start *string
}

func DefaultAnnotationsListOperationOptions() AnnotationsListOperationOptions {
	return AnnotationsListOperationOptions{}
}

func (o AnnotationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AnnotationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AnnotationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.End != nil {
		out.Append("end", fmt.Sprintf("%v", *o.End))
	}
	if o.Start != nil {
		out.Append("start", fmt.Sprintf("%v", *o.Start))
	}
	return &out
}

type AnnotationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AnnotationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AnnotationsList ...
func (c OpenapisClient) AnnotationsList(ctx context.Context, id ComponentId, options AnnotationsListOperationOptions) (result AnnotationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AnnotationsListCustomPager{},
		Path:          fmt.Sprintf("%s/annotations", id.ID()),
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
		Values *[]Annotation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AnnotationsListComplete retrieves all the results into a single object
func (c OpenapisClient) AnnotationsListComplete(ctx context.Context, id ComponentId, options AnnotationsListOperationOptions) (AnnotationsListCompleteResult, error) {
	return c.AnnotationsListCompleteMatchingPredicate(ctx, id, options, AnnotationOperationPredicate{})
}

// AnnotationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) AnnotationsListCompleteMatchingPredicate(ctx context.Context, id ComponentId, options AnnotationsListOperationOptions, predicate AnnotationOperationPredicate) (result AnnotationsListCompleteResult, err error) {
	items := make([]Annotation, 0)

	resp, err := c.AnnotationsList(ctx, id, options)
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

	result = AnnotationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
