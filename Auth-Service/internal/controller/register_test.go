package controller

import (
	"Auth-Service/internal/controller/mocks"
	"Auth-Service/internal/dtos"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/parser/factory"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/logger"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewRegisterController(t *testing.T) {
	log := logger.NewLogger()

	type args struct {
		logger  logger.ILogger
		service service.IAuthService
		parsers parser.IFactory
	}
	tests := []struct {
		name string
		args args
		want *registerController
	}{
		{
			name: "Test NewRegisterController",
			args: args{
				logger:  log,
				service: mocks.NewServiceMock(false),
				parsers: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegisterController(tt.args.logger, tt.args.service, tt.args.parsers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registerController_controller(t *testing.T) {
	log := logger.NewLogger()

	parsers := factory.NewParserFactory()
	_ = parsers.Set(parser.UserDtoToUserDomainParser, parser.NewUserDtoToUserDomainParser())

	type fields struct {
		logger  logger.ILogger
		service service.IAuthService
		parsers parser.IFactory
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantHttp int
		wantRes  *dtos.RegisterResponse
		testBody bool
	}{
		{
			name: "Test registerController success",
			fields: fields{
				logger:  log,
				service: mocks.NewServiceMock(false),
				parsers: parsers,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"/success",
					bytes.NewBufferString(`{"user":{"id":"1","name":"John Doe","email":"john@gmail.com", "password":"password"}}`),
				),
			},
			wantHttp: http.StatusOK,
			wantRes: &dtos.RegisterResponse{
				Code:    "200",
				Message: "Success",
			},
			testBody: true,
		},
		{
			name: "Test registerController invalid method",
			fields: fields{
				logger:  log,
				service: mocks.NewServiceMock(false),
				parsers: parsers,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodGet,
					"/invalid-method",
					nil,
				),
			},
			wantHttp: http.StatusMethodNotAllowed,
			testBody: false,
			wantRes:  nil,
		},
		{
			name: "Test registerController invalid body",
			fields: fields{
				logger:  log,
				service: mocks.NewServiceMock(false),
				parsers: parsers,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"/invalid-body",
					bytes.NewBufferString(`{"user":{"id":"1","name":"John Doe","email":"john@gmail.com", "password":"password"`),
				),
			},
			wantHttp: http.StatusBadRequest,
			testBody: false,
			wantRes:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &registerController{
				logger:  tt.fields.logger,
				service: tt.fields.service,
				parsers: tt.fields.parsers,
			}
			c.controller(tt.args.w, tt.args.r)
			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			defer res.Body.Close()
			if res.StatusCode != tt.wantHttp {
				t.Errorf("registerController() http status = %v, want %v", res.StatusCode, tt.wantHttp)
			}
			if tt.testBody {
				var gotRes dtos.RegisterResponse
				err := json.NewDecoder(res.Body).Decode(&gotRes)
				if err != nil {
					t.Errorf("registerController() error decoding response body: %v", err)
				}
				if !reflect.DeepEqual(&gotRes, tt.wantRes) {
					t.Errorf("registerController() response body = %v, want %v", gotRes, tt.wantRes)
				}
			}
		})
	}
}
