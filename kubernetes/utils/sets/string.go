package sets

type String map[string]Empty

func NewString(items ...string) String {
	return String(New[string](items...))
}

func (s String) Has(item string) bool {
	return cast(s).Has(item)
}
