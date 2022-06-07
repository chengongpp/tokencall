package tokencall

type AliyunService struct {
	GeneralService
}

func (s *AliyunService) Configure(ak string, as string, url string, conf map[string]string) ApiService {
	s.Ak = ak
	s.As = as
	if url != "" {
		s.Url = url
	}
	return s
}

func (s *AliyunService) Leak() (map[string]string, error) {
	panic("unimplemented!")
}
