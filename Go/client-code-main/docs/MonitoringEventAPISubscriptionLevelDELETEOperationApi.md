# {{classname}}

All URIs are relative to *{apiRoot}/3gpp-monitoring-event/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScsAsIdSubscriptionsSubscriptionIdDelete**](MonitoringEventAPISubscriptionLevelDELETEOperationApi.md#ScsAsIdSubscriptionsSubscriptionIdDelete) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Deletes an already existing monitoring event subscription

# **ScsAsIdSubscriptionsSubscriptionIdDelete**
> []MonitoringEventReport ScsAsIdSubscriptionsSubscriptionIdDelete(ctx, scsAsId, subscriptionId)
Deletes an already existing monitoring event subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scsAsId** | **string**| Identifier of the SCS/AS | 
  **subscriptionId** | **string**| Identifier of the subscription resource | 

### Return type

[**[]MonitoringEventReport**](MonitoringEventReport.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

