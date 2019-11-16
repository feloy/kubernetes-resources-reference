package api

type ApiGroup string

func (g ApiGroup) String() string {
	return string(g)
}

func (this ApiGroup) LessThan(that ApiGroup) bool {
	// admissionregistration > apiextensions
	if this.String() == "apiextensions" && that.String() == "admissionregistration" {
		return true
	}
	if that.String() == "apiextensions" && this.String() == "admissionregistration" {
		return false
	}

	// policy > extensions
	if this.String() == "extensions" && that.String() == "policy" {
		return true
	}
	if that.String() == "extensions" && this.String() == "policy" {
		return false
	}
	return this.String() < that.String()
}
