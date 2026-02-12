package iotcentrals

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DestinationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Destination
}

type DestinationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Destination
}

type DestinationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DestinationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DestinationsList ...
func (c IotcentralsClient) DestinationsList(ctx context.Context) (result DestinationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DestinationsListCustomPager{},
		Path:       "/dataExport/destinations",
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

	temp := make([]Destination, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalDestinationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for Destination (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// DestinationsListComplete retrieves all the results into a single object
func (c IotcentralsClient) DestinationsListComplete(ctx context.Context) (DestinationsListCompleteResult, error) {
	return c.DestinationsListCompleteMatchingPredicate(ctx, DestinationOperationPredicate{})
}

// DestinationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DestinationsListCompleteMatchingPredicate(ctx context.Context, predicate DestinationOperationPredicate) (result DestinationsListCompleteResult, err error) {
	items := make([]Destination, 0)

	resp, err := c.DestinationsList(ctx)
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

	result = DestinationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
