package enum

type ReactTypeValue string

type reactType struct {
	Clap  ReactTypeValue
	Like  ReactTypeValue
	Heart ReactTypeValue
}

var ReactType = &reactType{
	Clap:  "CLAP",
	Like:  "LIKE",
	Heart: "HEART",
}
