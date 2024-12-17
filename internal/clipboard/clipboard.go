package clipboard

type Copier interface {
	CopyToSlot(slot int) error
}

type Paster interface {
	PasteFromSlot(slot int) error
}

type Manager interface {
	Copier
	Paster
}
