package pagination

import (
	"fmt"
	"math"
	"server/app/models"

	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, p *models.PaginationData, tableName string) func(db *gorm.DB) *gorm.DB {
	dbClone := db.Session(&gorm.Session{})

	var count int32
	dbClone.Select(fmt.Sprintf("COUNT(DISTINCT %s.id)", tableName)).Scan(&count)

	p.Metadata.Total = count
	pages := int(math.Ceil(float64(p.Metadata.Total) / float64(p.Metadata.PerPage)))

	if pages == 0 {
		pages = 1
	}

	p.Metadata.Pages = pages
	last := pages

	if p.Metadata.Page > last {
		p.Metadata.Page = last
	}

	if p.Metadata.Page > 1 {
		p.Metadata.Prev = p.Metadata.Page - 1
	}

	if p.Metadata.Page == last {
		p.Metadata.Next = 0
		p.Metadata.Count = int(count) - ((p.Metadata.Page - 1) * p.Metadata.PerPage)
	} else {
		p.Metadata.Next = p.Metadata.Page + 1
		p.Metadata.Count = p.Metadata.PerPage
	}

	offset := p.Metadata.PerPage * (p.Metadata.Page - 1)

	if count == 0 {
		p.Metadata.From = 0
		p.Metadata.To = 0
	} else {
		p.Metadata.From = offset + 1
		p.Metadata.To = int(math.Min(float64(offset+p.Metadata.Count), float64(count)))
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(p.Metadata.PerPage)
	}
}
