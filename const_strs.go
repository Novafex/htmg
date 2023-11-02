package htmg

var (
	CHAR_SPACE     = []byte{' '}
	CHAR_EQ        = []byte{'='}
	CHAR_QT        = []byte{'"'}
	CHAR_LT        = []byte{'<'}
	CHAR_GT        = []byte{'>'}
	STR_EQ_QT      = []byte{'=', '"'}
	STR_CLOSE      = []byte{'<', '/'}
	STR_SELF_CLOSE = []byte{' ', '/', '>'}
)
