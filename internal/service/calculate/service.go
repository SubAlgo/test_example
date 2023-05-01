package calculate

import (
	"labTest/internal/ports"
)

type service struct {
	numberRepo ports.NumberRepository
}

func New(numberRepo ports.NumberRepository) ports.CalculateService {
	return &service{
		numberRepo: numberRepo,
	}
}

func (s *service) Cal() int {

	var (
		chN1 = make(chan int)
		chN2 = make(chan int)
	)

	go func() {
		chN1 <- s.numberRepo.GetN1()
	}()

	go func() {
		chN2 <- s.numberRepo.GetN2()
	}()

	return <-chN1 + <-chN2
	// n1 := s.numberRepo.GetN1()
	// fmt.Println("N1 >>>>>>", n1)
	// n2 := s.numberRepo.GetN2()
	// return n1 + n2
}
