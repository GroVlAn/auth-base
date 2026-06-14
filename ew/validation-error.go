package ew

type ValidationField struct {
	Field  string
	Reason string
}

type ErrValidation struct {
	msg    string
	fields []ValidationField
}

func NewErrValidation(msg string) *ErrValidation {
	return &ErrValidation{
		msg:    msg,
		fields: make([]ValidationField, 0),
	}
}

func (ev *ErrValidation) AddField(field, reason string) {
	ev.fields = append(ev.fields, ValidationField{
		Field:  field,
		Reason: reason,
	})
}

func (ev *ErrValidation) Error() string {
	return ev.msg
}

func (ev *ErrValidation) Fields() []ValidationField {
	return ev.fields
}

func (ev *ErrValidation) IsEmpty() bool {
	return len(ev.fields) == 0
}
