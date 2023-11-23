package masking

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}

type Service struct {
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) *Service {
	return &Service{
		prod: prod,
		pres: pres,
	}
}

func (s *Service) Run() error {
	data, err := s.prod.produce()
	if err != nil {
		return err
	}

	maskedData := make([]string, 0, len(data))
	for _, d := range data {
		masked := MaskLinks(d)
		maskedData = append(maskedData, masked)
	}

	if err := s.pres.present(maskedData); err != nil {
		return err
	}

	return nil
}
