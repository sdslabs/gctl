package openapi

// CreatedDatabase struct for CreatedDatabase
type CreateRepository struct {
	// Name of the Repository
	Name string `json:"user,omitempty"`
	//Path to project directory
	Path string `json:"user,omitempty"`
}
