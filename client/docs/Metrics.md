# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Document&#39;s identifier assigned by MongoDB | [optional] 
**Name** | **string** | Name of the application | [optional] 
**HostIp** | **string** | IPv4 address of the node where the application is deployed | [optional] 
**Alive** | **bool** | Is the application up and running? | [optional] 
**OnlineCpus** | **int32** | The number of available CPUs in the node | [optional] 
**CpuUsage** | **float64** | Fraction of total CPUs utilized by the application | [optional] 
**MemoryUsage** | **float64** | Fraction of memory utilized by the application | [optional] 
**MaxMemoryUsage** | **float64** | Highest fraction of memory utilized by the application in its lifetime | [optional] 
**MemoryLimit** | **float32** | Memory (in GigaBytes) alloted to the application | [optional] 
**Timestamp** | **int64** | Unix timestamp of the metrics document | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


