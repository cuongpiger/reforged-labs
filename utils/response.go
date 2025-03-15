package utils

type ResponseBuildProcess interface {
	SetData(interface{}) ResponseBuildProcess
	SetStringMessage(string) ResponseBuildProcess
	SetStatus(ResponseStatus) ResponseBuildProcess

	GetResponse() ResponseBase
}

type (
	ResponseBase struct {
		Status  ResponseStatus `json:"status"`
		Message string         `json:"message"`
		Data    interface{}    `json:"data,omitempty"`
	}

	ResponseStatus string
)

const (
	ResponseStatusSuccess ResponseStatus = "success"
	ResponseStatusFailed  ResponseStatus = "failed"
)

func NewResponse() ResponseBuildProcess {
	return &ResponseBase{
		Status: ResponseStatusSuccess,
	}
}

func (s *ResponseBase) SetStringMessage(pmsg string) ResponseBuildProcess {
	s.Message = pmsg
	return s
}

func (s *ResponseBase) SetStatus(pstatus ResponseStatus) ResponseBuildProcess {
	s.Status = pstatus
	return s
}

func (s *ResponseBase) SetData(pdata interface{}) ResponseBuildProcess {
	s.Data = &pdata
	return s
}

func (s *ResponseBase) GetResponse() ResponseBase {
	return *s
}
