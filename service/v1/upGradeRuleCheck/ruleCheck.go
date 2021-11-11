package upGradeRuleCheck

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	"strconv"
	"strings"

)

func compareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		x, y := 0, 0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}


func checkVersion(userVersion,minVersion,maxversion string) bool{
	 comp1:=compareVersion(userVersion,minVersion)
	 if comp1==-1 {
		 return false
	 }
	 comp2:=compareVersion(userVersion,maxversion)
	 if comp2==1 {
		 return false
	 }
	 return true
}

//检查设备号是否在白名单中
func checkDeviceID(DID *string) bool{

}

//比较顺序：业务id > platform > 渠道 > 设备白名单 > 【其他条件计算顺序均可】
func checkSingleRule(rule *model.Rule,info *model.Info) bool{
	if rule.AID!=info.AID {
		return false
	}
	if rule.Platform!=info.DevicePlatform {
		return false
	}
	if rule.ChannelNumber!=info.ChannelNumber {
		return false
	}
	if !checkDeviceID(&info.DeviceID) {
		return false
	}
	if rule.CPUArch!=info.CPUArch{
		return false
	}

}




//检查客户端传来参数Info是否符合Rule规则
func check(rules *[]model.Rule,info *model.Info) (model.Rule,bool){
	for i:=0;i<len(*rules);i++{

	}
}
