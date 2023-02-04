package jobs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllJobsInSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VMwareJob

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetAllJobsInSiteOperationResponse, error)
}

type GetAllJobsInSiteCompleteResult struct {
	Items []VMwareJob
}

func (r GetAllJobsInSiteOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetAllJobsInSiteOperationResponse) LoadMore(ctx context.Context) (resp GetAllJobsInSiteOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetAllJobsInSite ...
func (c JobsClient) GetAllJobsInSite(ctx context.Context, id VMwareSiteId) (resp GetAllJobsInSiteOperationResponse, err error) {
	req, err := c.preparerForGetAllJobsInSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetAllJobsInSite", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetAllJobsInSite", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetAllJobsInSite(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetAllJobsInSite", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetAllJobsInSite prepares the GetAllJobsInSite request.
func (c JobsClient) preparerForGetAllJobsInSite(ctx context.Context, id VMwareSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/jobs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetAllJobsInSiteWithNextLink prepares the GetAllJobsInSite request with the given nextLink token.
func (c JobsClient) preparerForGetAllJobsInSiteWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAllJobsInSite handles the response to the GetAllJobsInSite request. The method always
// closes the http.Response Body.
func (c JobsClient) responderForGetAllJobsInSite(resp *http.Response) (result GetAllJobsInSiteOperationResponse, err error) {
	type page struct {
		Values   []VMwareJob `json:"value"`
		NextLink *string     `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetAllJobsInSiteOperationResponse, err error) {
			req, err := c.preparerForGetAllJobsInSiteWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetAllJobsInSite", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetAllJobsInSite", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetAllJobsInSite(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetAllJobsInSite", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetAllJobsInSiteComplete retrieves all of the results into a single object
func (c JobsClient) GetAllJobsInSiteComplete(ctx context.Context, id VMwareSiteId) (GetAllJobsInSiteCompleteResult, error) {
	return c.GetAllJobsInSiteCompleteMatchingPredicate(ctx, id, VMwareJobOperationPredicate{})
}

// GetAllJobsInSiteCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c JobsClient) GetAllJobsInSiteCompleteMatchingPredicate(ctx context.Context, id VMwareSiteId, predicate VMwareJobOperationPredicate) (resp GetAllJobsInSiteCompleteResult, err error) {
	items := make([]VMwareJob, 0)

	page, err := c.GetAllJobsInSite(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := GetAllJobsInSiteCompleteResult{
		Items: items,
	}
	return out, nil
}
