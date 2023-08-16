package platforms

import "strings"

type Platform string

const Web Platform = "WEB"

var m = map[string]Platform{
	string(Web): Web,
}

func Parse(plt string) Platform {
	return m[strings.ToUpper(plt)]
}

func (p Platform) String() string {
	return string(p)
}
