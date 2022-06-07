package tokencall

type ApiService interface {
	Configure(ak string, as string, url string, conf map[string]string) ApiService
	Leak() (map[string]string, error)
}

type GeneralService struct {
	Ak  string
	As  string
	Url string
}
