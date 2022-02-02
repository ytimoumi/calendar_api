package models

import "time"

type Calendar struct {
	Id    int
	Start time.Time
	End   time.Time
}

func NewCalendar(id int, start, end time.Time) *Calendar {
	return &Calendar{
		Id:    id,
		Start: start,
		End:   end,
	}
}

type ToReserve struct {
	Start time.Time
	End   time.Time
	Title string
	Email string
}

func NewToReserve(start, end time.Time, title, email string) ToReserve {
	return ToReserve{
		Start: start,
		End:   end,
		Title: title,
		Email: email,
	}
}

type Response struct {
	Code    int 	`json:"code"`
	Message string 	`json:"message"`
}
func NewResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

type Availability struct {
	Start time.Time
	End   time.Time
}

func NewAvailability(start, end time.Time) Availability {
	return Availability{
		Start: start,
		End:   end,
	}
}
