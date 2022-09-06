package define

type EmailConfig struct {
	ServerName string `ini:"server"`
	Port       int    `ini:"port"`
	From       string `ini:"from"`
	Password   string `ini:"password"`
	Second     int    `ini:"second"`
}

type RedisConfig struct {
	Server      string `ini:"server"`
	Port        int    `ini:"port"`
	RedisPrefix string `ini:"prefix"`
}

type MysqlConfig struct {
	Server   string `ini:"server"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
}

type CosConfig struct {
	Server    string `ini:"server"`
	SecretID  string `ini:"secret_id"`
	SecretKey string `ini:"secret_key"`
}

// 验证码长度
var Code_length = 6

// 验证码有效时间
var Code_expire = 300

// 默认分页大小
var PageSize = 20
