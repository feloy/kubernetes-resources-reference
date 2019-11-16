package api

type ApiKind string

func (k ApiKind) String() string {
	return string(k)
}
