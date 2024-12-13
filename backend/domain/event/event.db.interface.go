package event

type EventPersistenceInterface interface {
	Save(event EventInterface) error
}