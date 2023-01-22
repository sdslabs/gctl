package openapi

// CreateRepository struct for CreateRepository
type CreateRepository struct {
	// Name of the Repository
	Name string `json:"name,omitempty"`
	//Path to project directory
	Path string `json:"path,omitempty"`
}
