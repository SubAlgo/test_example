package calculate_test

import (
	"labTest/internal/ports"
	"labTest/internal/ports/mocks"
	"labTest/internal/service/calculate"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testModule struct {
	numberRepo *mocks.NumberRepository
	svc        ports.CalculateService
}

func newTestModule(t *testing.T) *testModule {
	numberRepo := mocks.NewNumberRepository(t)
	return &testModule{
		numberRepo: numberRepo,
		svc:        calculate.New(numberRepo),
	}
}

type test struct {
	name     string
	args     []interface{}
	mockFn   func(*testModule)
	assertFn func()
}

func TestCal(t *testing.T) {
	var res int
	tests := []test{
		{
			name: "case1",
			args: []interface{}{},
			mockFn: func(mod *testModule) {
				mod.numberRepo.On("GetN1").Return(3)
				mod.numberRepo.On("GetN2").Return(2)
			},
			assertFn: func() {
				assert.Equal(t, 5, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testModule := newTestModule(t)
			if tt.mockFn != nil {
				tt.mockFn(testModule)
			}
			res = testModule.svc.Cal()
			tt.assertFn()
		})
	}
}
