package model

// User 后台登录用户信息
type User struct {
	Username string `json:"username" label:"用户名" validate:"min=6,max=10"`
	Password string `json:"password" label:"密码" validate:"min=6,max=12"`
}

// Rule 配置的新版本更新规则
type Rule struct {
	ID                   int    `json:"id"`
	AppID                int    `json:"app_id" label:"app的唯一标识"`
	Platform             string `json:"platform" label:"平台"`
	DownloadURL          string `json:"download_url" label:"包的下载链接"`
	UpdateVersionCode    string `json:"update_version_code" label:"当前包的版本号"`
	Md5                  string `json:"md5" label:"包的MD5"`
	MaxUpdateVersionCode string `json:"max_update_version_code" label:"可升级的最大版本号"`
	MinUpdateVersionCode string `json:"min_update_version_code" label:"可升级的最小版本号"`
	MaxOSApi             int    `json:"max_os_api" label:"支持的最大操作系统版本"`
	MinOSApi             int    `json:"min_os_api" label:"支持的最小操作系统版本"`
	CPUArch              int    `json:"cpu_arch" label:"设备的CPU架构"`
	ChannelNumber        string `json:"channel_number" label:"渠道号"`
	Title                string `json:"title" label:"弹窗标题"`
	UpdateTips           string `json:"update_tips" label:"弹窗的更新文本"`
	Status               bool   `json:"status" label:"规则上线或者下线情况，0表示下线，1表示上线"`
}

// Info 客户端上报的参数信息
type Info struct {
	Version           string `json:"version" label:"请求api版本"`
	DevicePlatform    string `json:"device_platform" label:"设备平台"`
	DeviceID          string `json:"device_id" label:"设备ID"`
	OSApi             int    `json:"os_api" label:"安卓的系统版本"`
	ChannelNumber     string `json:"channel_number" label:"渠道，标识包的类型"`
	VersionCode       string `json:"version_code" label:"应用的大版本"`
	UpdateVersionCode string `json:"update_version_code" label:"应用的小版本"`
	AppID             int    `json:"app_id" label:"app的唯一标识"`
	CPUArch           int    `json:"cpu_arch" label:"设备的CPU架构"`
}

// Device 设备ID白名单
type Device struct {
	RuleID   int    `json:"rule_id" label:"这个白名单对应的规则的id"`
	DeviceID string `json:"device_id" label:"设备ID"`
}
