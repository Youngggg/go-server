package cfg

// 各模块日志配置
type LogConfig struct {
	BaseConfig     *BaseLogConfig `mapstructure:"base_config"`
	LinkedinConfig *BaseLogConfig `mapstructure:"linkedin_config"`
	FbConfig       *BaseLogConfig `mapstructure:"fb_config"`
	BinanceConfig  *BaseLogConfig `mapstructure:"binance_config"`
	AmazonConfig   *BaseLogConfig `mapstructure:"amazon_config"`
	InsConfig      *BaseLogConfig `mapstructure:"ins_config"`
	TwitterConfig  *BaseLogConfig `mapstructure:"twitter_config"`
	ZaloConfig     *BaseLogConfig `mapstructure:"zalo_config"`
}

// 当业务应用不直接调用vlog打日志
// 而是封装一层后在打印,这时候需要配置caller_skip=1,其他情况下均不需要设置此值
type BaseLogConfig struct {
	LogPath     string `mapstructure:"log_path"`     // log路径
	LogLevel    string `mapstructure:"log_level"`    // log级别 zap-log级别 -1:debug/0:info/1:warn/2:error/3:dpanic/4:panic/5:fatal
	CallerSkip  int    `mapstructure:"caller_skip"`  // 跳过调用者的调用堆栈深度
	ServiceName string `mapstructure:"service_name"` // 日志里额外参数:表明哪个服务的日志
	//ModuleName  string `yaml:"module_name"`  // 模块名称:此处用来区分各个筛料模块的日志 不同模块创建不同文件夹存储对应的info/error等log
}

// elk日志配置 暂不使用elasticsearch kafka
type LoggerConfig struct {
	EnableKafka bool
	Base        *BaseLogConfig
	Kafka       *KafkaLoggerConfig
}

// 日志kafka配置 暂不使用
type KafkaLoggerConfig struct {
	NameServer []string
	InfoTopic  string
	ErrorTopic string
	WarnTopic  string
	Filter     []string
}

type LogCfgOptFunc func(opt *LoggerConfig)

// 默认日志配置
func GetDefaultLogConfig(serviceName string, opts ...LogCfgOptFunc) *LoggerConfig {
	loggerConfig := &LoggerConfig{
		EnableKafka: false,
		Base: &BaseLogConfig{
			LogPath:     "/tmp/log",
			LogLevel:    "debug",
			ServiceName: serviceName,
			CallerSkip:  0,
		},
		Kafka: nil,
	}
	for _, opt := range opts {
		opt(loggerConfig)
	}
	return loggerConfig
}

func WithLogPath(logPath string) LogCfgOptFunc {
	return func(opt *LoggerConfig) {
		if logPath != "" {
			opt.Base.LogPath = logPath
		}
	}
}
