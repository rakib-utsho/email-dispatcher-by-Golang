package main

type Recipint struct {
	Name  string
	Email string
}

func main() {
	loadRecipients("./emails.csv")
}
