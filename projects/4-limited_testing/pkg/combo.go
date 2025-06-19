package pkg

func Use() error {
	_, err := Open("test.ffmap")
	if err != nil {
		return err
	}
	_, _, _, err = PopulateData()
	return err
}
