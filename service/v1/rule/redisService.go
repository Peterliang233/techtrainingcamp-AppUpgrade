package rule

import (
	"log"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/rule/ruleconfig"
	"github.com/robfig/cron/v3"
)

// CacheData 将数据库里面的缓存的数据都加入到redis里面
func CacheData() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("redis的协程错误")
		}
	}()

	c := cron.New()

	_, err := c.AddFunc("0 */12 * * *", func() {
		if err := CacheRules(); err != nil {
			log.Println("对规则进行重新缓存错误")
		}
		if err := CacheWhiteList(); err != nil {
			log.Println("对白名单进行重新缓存错误")
		}
	})

	if err != nil {
		log.Println("定时器处理错误")
	}
	c.Start()
}

// CacheRules 缓存所有的规则
func CacheRules() error {
	rules, err := ruleconfig.GetAllRules()
	if err != nil {
		log.Printf("规则配置部分的redis错误, %v\n", err)
		return err
	}
	for _, rule := range rules {
		if rule.Status {
			ruleconfig.CacheOnline(strconv.Itoa(rule.ID))
		}
		ruleconfig.CacheBasicInfo(rule.Platform, rule.ChannelNumber, rule.CPUArch, rule.AppID, rule.ID)
		ruleconfig.CacheOsApi(rule.MinOSApi, rule.MaxOSApi, rule.ID)
		ruleconfig.CacheUpdateVersionCode(rule.MinUpdateVersionCode, rule.MaxUpdateVersionCode, rule.ID)
	}
	return nil
}

// CacheWhiteList 对白名单进行缓存
func CacheWhiteList() error {
	devices, err := ruleconfig.GetDeviceWhiteList()
	if err != nil {
		log.Fatalf("白名单部分的redis错误, %v\n", err)
		return err
	}
	for _, device := range devices {
		CacheDeviceID(&device)
	}
	return nil
}
