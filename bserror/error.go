package bserror

type InsufficientStockError struct {
	Msg string
}

func (e *InsufficientStockError) Error() string {
	return e.Msg
}

// -------------------------------------------------------- //

type BadParameterError struct {
	Msg string
}

func (e *BadParameterError) Error() string {
	return e.Msg
}

// -------------------------------------------------------- //

type DataVersionError struct {
	Msg string
}

func (e *DataVersionError) Error() string {
	return e.Msg
}

// -------------------------------------------------------- //

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}
