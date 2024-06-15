package todo

type InternalError struct{ Message string }
type AuthorizationError struct{ Message string }
type DBError struct{ Message string }
type BadFormatError struct{ Message string }
type NoDataFound struct{ Message string }

func (err InternalError) Error() string      { return err.Message }
func (err AuthorizationError) Error() string { return err.Message }
func (err DBError) Error() string            { return err.Message }
func (err BadFormatError) Error() string     { return err.Message }
func (err NoDataFound) Error() string        { return err.Message }
