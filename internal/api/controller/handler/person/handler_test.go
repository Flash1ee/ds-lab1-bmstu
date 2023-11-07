package person_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"crud-app/internal/api/controller/handler/person"
	"crud-app/internal/domain"
)

func TestPatchHandler(t *testing.T) {
	type fields struct {
		storage *Mockstorage
	}
	testCases := map[string]struct {
		request          string
		expectedStatus   int
		expectedResponse string
		fields           func(ctrl *gomock.Controller) fields
	}{
		"ok": {
			request:          `{"name": "John Doe", "age": 30, "address": "123 Main St", "work": "Software Developer"}`,
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":1,"name":"John Doe","age":30,"address":"123 Main St","work":"Software Developer"}`,
			fields: func(ctrl *gomock.Controller) fields {
				st := NewMockstorage(ctrl)
				st.EXPECT().Patch(&domain.Person{
					ID:      1,
					Name:    "John Doe",
					Age:     30,
					Address: "123 Main St",
					Work:    "Software Developer",
				}).Return(&domain.Person{
					ID:      1,
					Name:    "John Doe",
					Age:     30,
					Address: "123 Main St",
					Work:    "Software Developer",
				}, nil)
				return fields{
					storage: st,
				}
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/api/v1/person/:id", strings.NewReader(tc.request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("1")

			h := person.New(e.Group("api/v1"), tc.fields(ctrl).storage)

			err := h.Patch(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			if tc.expectedStatus == http.StatusOK {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestAddHandler(t *testing.T) {
	type fields struct {
		storage *Mockstorage
	}
	testCases := map[string]struct {
		request          string
		expectedStatus   int
		expectedResponse string
		fields           func(ctrl *gomock.Controller) fields
	}{
		"ok": {
			request:          `{"id": 1, "name": "John Doe", "age": 30, "address": "123 Main St", "work": "Software Developer"}`,
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{}`,
			fields: func(ctrl *gomock.Controller) fields {
				st := NewMockstorage(ctrl)
				st.EXPECT().Create(&domain.Person{
					ID:      1,
					Name:    "John Doe",
					Age:     30,
					Address: "123 Main St",
					Work:    "Software Developer",
				}).Return(nil)
				return fields{
					storage: st,
				}
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/persons", strings.NewReader(tc.request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := person.New(e.Group("api/v1"), tc.fields(ctrl).storage)

			err := h.Add(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)

			if tc.expectedStatus == http.StatusCreated {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	type fields struct {
		storage *Mockstorage
	}
	testCases := map[string]struct {
		request          string
		expectedStatus   int
		expectedResponse string
		wantErr          bool
		fields           func(ctrl *gomock.Controller) fields
	}{
		"ok": {
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":1,"name":"John Doe","age":30,"address":"123 Main St","work":"Software Developer"}`,
			fields: func(ctrl *gomock.Controller) fields {
				st := NewMockstorage(ctrl)
				st.EXPECT().Find(1).Return(&domain.Person{
					ID:      1,
					Name:    "John Doe",
					Age:     30,
					Address: "123 Main St",
					Work:    "Software Developer",
				}, nil)
				return fields{
					storage: st,
				}
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/persons/1", strings.NewReader(tc.request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := person.New(e.Group("api/v1"), tc.fields(ctrl).storage)
			c.SetParamNames("id")
			c.SetParamValues("1")

			err := h.Get(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			if !tc.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetAllHandler(t *testing.T) {
	type fields struct {
		storage *Mockstorage
	}
	testCases := map[string]struct {
		request          string
		expectedStatus   int
		expectedResponse string
		wantErr          bool
		fields           func(ctrl *gomock.Controller) fields
	}{
		"ok": {
			expectedStatus:   http.StatusOK,
			expectedResponse: `[{"address":"123 Main St", "age":30, "id":1, "name":"John Doe", "work":"Software Developer"},{"address":"123 Main St", "age":30, "id":2, "name":"John Doe", "work":"Software Developer"}]`,
			fields: func(ctrl *gomock.Controller) fields {
				st := NewMockstorage(ctrl)
				st.EXPECT().GetAll().Return([]domain.Person{{
					ID:      1,
					Name:    "John Doe",
					Age:     30,
					Address: "123 Main St",
					Work:    "Software Developer",
				}, {
					ID:      2,
					Name:    "John Doe",
					Age:     30,
					Address: "123 Main St",
					Work:    "Software Developer",
				}})
				return fields{
					storage: st,
				}
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/persons", strings.NewReader(tc.request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := person.New(e.Group("api/v1"), tc.fields(ctrl).storage)

			err := h.GetAll(c)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())

			if !tc.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
