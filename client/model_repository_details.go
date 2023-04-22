package openapi

// CreateRepository struct for CreateRepository
type CreateRepository struct {
	// Name of the Repository
	Name string `json:"name,omitempty"`
}

type DeleteRepository struct {
	// GitURL of the Repository
	GitURL string `json:"giturl" bson:"giturl"`
}
