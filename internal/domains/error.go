package domains

const (
	Validation ErrorKind = iota
	Policy
	NotFound
	Internal
)

type (
	ErrorKind   uint8
	DomainError struct {
		Kind    ErrorKind
		Message string
	}
)

func (d *DomainError) IsValidation() bool {
	return d.Kind == Validation
}

func (d *DomainError) IsPolicy() bool {
	return d.Kind == Policy
}

func (d *DomainError) IsNotFound() bool {
	return d.Kind == NotFound
}

func (d *DomainError) IsInternal() bool {
	return d.Kind == Internal
}

func (d *DomainError) Error() string {
	return d.Message
}
