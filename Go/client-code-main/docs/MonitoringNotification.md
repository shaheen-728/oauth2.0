# MonitoringNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [***Ts29122CommonDataYamlcomponentsschemasLink**](./TS29122_CommonData.yaml#/components/schemas/Link.md) |  | [default to null]
**ConfigResults** | [**[]Ts29122CommonDataYamlcomponentsschemasConfigResult**](./TS29122_CommonData.yaml#/components/schemas/ConfigResult.md) | Each element identifies a notification of grouping configuration result. | [optional] [default to null]
**MonitoringEventReports** | [**[]MonitoringEventReport**](MonitoringEventReport.md) | Monitoring event reports. | [optional] [default to null]
**CancelInd** | **bool** | Indicates whether to request to cancel the corresponding monitoring subscription. Set to false or omitted otherwise. | [optional] [default to null]
**AppliedParam** | [***AppliedParameterConfiguration**](AppliedParameterConfiguration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

