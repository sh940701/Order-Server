package util

func PanicHandler (err error) {
	if err != nil {
		panic(err)
	}
}