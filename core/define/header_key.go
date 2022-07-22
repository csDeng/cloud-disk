package define

type Header_Key struct {
	Uid          string
	UserIdentity string
	UserName     string
}

var HKOBJ = Header_Key{
	Uid:          "UserId",
	UserIdentity: "UserIdentity",
	UserName:     "UserName",
}
