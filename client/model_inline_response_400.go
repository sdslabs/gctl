/*
 * Gasper Master API
 *
 * Handles authentication, creation/management of applications, databases, users and also provides a superuser API. <br><br> Only a superuser can avail the superuser API. A superuser can **grant/revoke** superuser privileges to other users. A default  superuser is created every time a Gasper instance is launched whose credentials are defined in the `admin` section of `config.toml`, the main configuration file. A sample configuration file is available [here](https://github.com/sdslabs/gasper/blob/develop/config.sample.toml#L37).<br><br> **Note:-** Normally the applications and databases can only be managed by their owners but the superuser can bypass that check.<br><br> **PS:-** If you want to programmatically generate a client for this API, you can find the corresponding OpenAPI specifications [here](https://github.com/sdslabs/gasper/tree/develop/docs/content/api/specs). We recommend using [OpenAPI-Generator](https://openapi-generator.tech/) for generating clients.
 *
 * API version: 1.0
 * Contact: contact@sdslabs.co.in
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {
	Success bool `json:"success,omitempty"`
	Error string `json:"error,omitempty"`
}

// InlineResponse401 struct for InlineResponse401
type InlineResponse401 struct {
	Success bool `json:"success,omitempty"`
	Error string `json:"error,omitempty"`
}

// InlineResponse500 struct for InlineResponse500
type InlineResponse500 struct {
	Success bool `json:"success,omitempty"`
	Error string `json:"error,omitempty"`
}
