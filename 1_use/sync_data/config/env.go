package config

var POPdb dbConfg = dbConfg{
    Name:    "",
    IP:      "",
    Port:    "",
    Pwd:     "",
    User:    "",
    Charset: "utf8",
}

var Vpsdb dbConfg = dbConfg{

    Name:    "",
    IP:      "",
    Port:    "",
    Pwd:     "",
    User:    "",
    Charset: "utf8",
}

var redisEnv RedisConf = RedisConf{
    IP:         []string{""},
    MasterName: "",
    Pwd:        "",
    DbIndex:    0,
}

// var ApiConf = map[string]string{
//     "url":   "http://tmdmapi.staff.xdf.cn",
//     "token": "ef4e3a2b-ae21-4cad-ada8-86ec46a8a83g",
//     "appid": "90101",
// }

var ApiConf = map[string]string{
    "url":   "",
    "token": "",
    "appid": "",
}

var ZstuMsgType = map[string]int{
    "h": 1,
    "m": 2,
    "r": 3,
}

func GetRedisDb(id int) RedisConf {
    var tmp RedisConf = redisEnv
    tmp.DbIndex = id
    return tmp
}

type dbConfg struct {
    Name    string
    IP      string
    Port    string
    Pwd     string
    User    string
    Charset string
}

type RedisConf struct {
    IP         []string
    MasterName string
    Pwd        string
    DbIndex    int
}

const BASETIME = "2006-01-02 15:04:05"
