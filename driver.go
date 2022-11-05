func getCh() []byte {
	t, _ := term.Open("/dev/tty")
	_ = term.RawMode(t)
	b := make([]byte, 3)
	numRead, _ := t.Read(b)
	_ = t.Restore()
	_ = t.Close()
	return b[0:numRead]
}

func getKey() string {
	fmt.Print("Enter a 32-bit cipher key.\n[ 0] >: ")

	var password string
	var star string

	for {

		password = ""
		star = ""

		for {

			var x = getCh()

			if bytes.Compare(x, []byte{13}) == 0 { // 13  == Enter Key was Pressed
				break

			} else if bytes.Compare(x, []byte{127}) == 0 { // 127 == Backspace Key was Pressed
				password = password[:(len(password) - 1)]
				star = star[:(len(star) - 1)]

			} else if bytes.Compare(x, []byte{32}) == 1 && bytes.Compare(x, []byte{127}) == -1 {
				password = password + string(x)
				star = star + "*"

			} else {
				continue
			}

			fmt.Printf("\r[%2d] >: %s", len(password), star)
		}

		if len(password) == 32 {
			break
		} else {
			fmt.Print("\n\nValue was not a 32-bit key. Try again!\n>: ")
		}
	}

	fmt.Println()

	return password
}
