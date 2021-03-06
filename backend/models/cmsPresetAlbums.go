package models

type CmsPresetAlbums struct {
	General   `xorm:"extends"`
	DeletedAt int    `xorm:"INT(11) deleted default(0) comment('软删除时间')"`
	AppId     string `xorm:"varchar(255) notnull default('') 'appId'"`
	AlbumId   int    `xorm:"BIGINT(20) notnull default(0)"`
}

func (c CmsPresetAlbums) TableName() string {
	return "cms_preset_albums"
}
