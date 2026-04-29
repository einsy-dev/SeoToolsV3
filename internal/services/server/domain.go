package server

type Domain struct{}

func (d *Domain) Fetch(domain string) (string, error) {
	return Fetch("https://edsail.com/api/domain/"+domain, FetchConfig{Token: "147852", Method: "GET"})
}
