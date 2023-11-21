package main

type APIClient interface {
	GetData(query string) (Response, error)
}

type Response struct {
	Text       string
	StatusCode int
}

type MockAPIClient interface {
	APIClient
	SetResponse(res Response, err error)
}

type Mock struct {
	Resp Response
	Err  error
}

func (m *Mock) GetData(query string) (Response, error) {
	return m.Resp, m.Err
}

func (m *Mock) SetResponse(res Response, err error) {
	m.Resp = res
	m.Err = err
}
