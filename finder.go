package traverse

type Finder interface {
	Find(identifier string) (string, error)
}
