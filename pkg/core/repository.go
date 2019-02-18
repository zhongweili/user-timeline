package core

// Repository defines the API repository implementation should follow.
type Repository interface {
	Find(id string) (*Following, error)
	FindAll(selector map[string]interface{}) ([]*Following, error)
	Delete(Following *Following) error
	Update(Following *Following) error
	Create(Following ...*Following) error
	Count() (int, error)
}
