package models

type AutoBuild struct {
	GeneralWithDeleted `xorm:"extends"`
	AppId              string `xorm:"varchar(32) notnull unique comment('目标App')" binding:"required"`
	CmsSourceApp       string `xorm:"varchar(32) comment('cms来源App')"`
	CmsExecTime        string `xorm:"varchar(32) comment('cms执行时间')"`
	Mqtt               int    `xorm:"smallint(1) default(0) comment('Mqtt状态:0未配置1已配置')"`
	Callback           string `xorm:"varchar(32) notnull default('none') comment('callback状态:默认空/rtoy/tongmi')"`
	UpgradeName        string `xorm:"default('') comment('升级配置名称')"`
	UpgradeVcode       int    `xorm:"default(0) comment('升级配置号')"`
	UpgradeVname       string `xorm:"default('') comment('升级配置展示名称')"`
	AlbumList          string `xorm:"default(0) comment('预置歌单列表,分隔')"`
}

func (c AutoBuild) TableName() string {
	return "auto_build"
}
