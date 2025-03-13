package response

import "github.com/gawbsouza/boot-help/httperr"

func (r *Response) Err(err *httperr.HttpError) *Response {
	if r != nil {
		r.err = err
	}
	return r
}

func (r *Response) BadErr(message string) *Response {
	if r != nil {
		r.err = httperr.Bad(message)
	}
	return r
}

func (r *Response) NotFoundErr(message string) *Response {
	if r != nil {
		r.err = httperr.NotFound(message)
	}
	return r
}

func (r *Response) ConflictErr(message string) *Response {
	if r != nil {
		r.err = httperr.Conflict(message)
	}
	return r
}

func (r *Response) ConditionErr(message string) *Response {
	if r != nil {
		r.err = httperr.Condition(message)
	}
	return r
}

func (r *Response) UnprocessableErr(message string) *Response {
	if r != nil {
		r.err = httperr.Unprocessable(message)
	}
	return r
}

func (r *Response) InternalErr(message string) *Response {
	if r != nil {
		r.err = httperr.Internal(message)
	}
	return r
}
