# MonitoringEventReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImeiChange** | [***AssociationType**](AssociationType.md) |  | [optional] [default to null]
**ExternalId** | [***Ts29122CommonDataYamlcomponentsschemasExternalId**](./TS29122_CommonData.yaml#/components/schemas/ExternalId.md) |  | [optional] [default to null]
**IdleStatusInfo** | [***IdleStatusInfo**](IdleStatusInfo.md) |  | [optional] [default to null]
**LocationInfo** | [***LocationInfo**](LocationInfo.md) |  | [optional] [default to null]
**LossOfConnectReason** | **int32** | If \&quot;monitoringType\&quot; is \&quot;LOSS_OF_CONNECTIVITY\&quot;, this parameter shall be included if available to identify the reason why loss of connectivity is reported. Refer to 3GPPÂ TSÂ 29.336Â [11] SubclauseÂ 8.4.58. | [optional] [default to null]
**MaxUEAvailabilityTime** | [***Ts29122CommonDataYamlcomponentsschemasDateTime**](./TS29122_CommonData.yaml#/components/schemas/DateTime.md) |  | [optional] [default to null]
**Msisdn** | [***Ts29122CommonDataYamlcomponentsschemasMsisdn**](./TS29122_CommonData.yaml#/components/schemas/Msisdn.md) |  | [optional] [default to null]
**MonitoringType** | [***MonitoringType**](MonitoringType.md) |  | [default to null]
**UePerLocationReport** | [***UePerLocationReport**](UePerLocationReport.md) |  | [optional] [default to null]
**PlmnId** | [***Ts29122CommonDataYamlcomponentsschemasPlmnId**](./TS29122_CommonData.yaml#/components/schemas/PlmnId.md) |  | [optional] [default to null]
**ReachabilityType** | [***ReachabilityType**](ReachabilityType.md) |  | [optional] [default to null]
**RoamingStatus** | **bool** | If \&quot;monitoringType\&quot; is \&quot;ROAMING_STATUS\&quot;, this parameter shall be set to \&quot;true\&quot; if the UE is on roaming status. Set to false or omitted otherwise. | [optional] [default to null]
**FailureCause** | [***FailureCause**](FailureCause.md) |  | [optional] [default to null]
**EventTime** | [***Ts29122CommonDataYamlcomponentsschemasDateTime**](./TS29122_CommonData.yaml#/components/schemas/DateTime.md) |  | [optional] [default to null]
**PdnConnInfo** | [***PdnConnectionInformation**](PdnConnectionInformation.md) |  | [optional] [default to null]
**DddStatus** | [***Ts29571CommonDataYamlcomponentsschemasDlDataDeliveryStatus**](./TS29571_CommonData.yaml#/components/schemas/DlDataDeliveryStatus.md) |  | [optional] [default to null]
**DddTrafDescriptor** | [***Ts29571CommonDataYamlcomponentsschemasDddTrafficDescriptor**](./TS29571_CommonData.yaml#/components/schemas/DddTrafficDescriptor.md) |  | [optional] [default to null]
**MaxWaitTime** | [***Ts29122CommonDataYamlcomponentsschemasDateTime**](./TS29122_CommonData.yaml#/components/schemas/DateTime.md) |  | [optional] [default to null]
**ApiCaps** | [**[]ApiCapabilityInfo**](ApiCapabilityInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

