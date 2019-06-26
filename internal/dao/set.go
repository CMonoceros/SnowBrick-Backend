package dao

import (
	"context"

	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/internal/model/entity"
)

func (d *Dao) RawBrickSet(c context.Context, page *entity.BasePage) (sets []*entity.BrickSet, err error) {
	if err = d.db.ReadOnlyTable(new(entity.BrickSet).TableName()).
		Limit(page.Limit).Offset(page.GetOffset()).
		Preload("Type").Preload("Theme").Preload("SubTheme").
		Preload("Brand").Preload("Packaging").
		Preload("Tags").Preload("Media").
		Find(&sets).Error; err != nil {
		log.Error("RawBrickSet d.db.ReadOnlyTable error(%v)", err)
		return
	}

	return
}
