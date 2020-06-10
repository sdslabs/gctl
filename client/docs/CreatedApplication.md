# CreatedApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | ID of the application&#39;s docker container | [optional] 
**ContainerPort** | **int64** | Port assigned by the node to the application&#39;s docker container | [optional] 
**DockerImage** | **string** | Docker image used in building the application&#39;s container | [optional] 
**AppUrl** | **string** | The domain name of the application (DNS entry is managed by GenDNS 💡) | [optional] 
**HostIp** | **string** | IPv4 address of the node | [optional] 
**NameServers** | **[]string** | The DNS NameServers used by the application&#39;s docker container | [optional] 
**InstanceType** | **string** | The kind of instance this application belongs to | [optional] 
**Language** | **string** | The programming language in which the application is written | [optional] 
**Owner** | **string** | Owner of the application | [optional] 
**SshCmd** | **string** | Command to SSH into the application&#39;s docker container | [optional] 
**Id** | **string** | Application&#39;s identifier assigned by MongoDB | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


