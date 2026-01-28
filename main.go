package main

type Recipint struct {
	Name  string
	Email string
}

func main() {

	recipentChannel := make(chan Recipint)
	loadRecipients("./emails.csv", recipentChannel)
}
