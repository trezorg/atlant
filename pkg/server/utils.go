package server

import (
	"github.com/trezorg/atlant/pkg/db"
	pb "github.com/trezorg/atlant/pkg/proto"
)

func SetPageDefaults(page *pb.Page) *pb.Page {
	if page.Limit == 0 {
		page.Limit = db.DefaultLimit
	}
	if page.Limit > db.MaxLimit {
		page.Limit = db.MaxLimit
	}
	if page.Limit < db.MinLimit {
		page.Limit = db.MinLimit
	}
	if page.Sorting == nil {
		page.Sorting = &pb.Sorting{
			Field: pb.SortingField_NAME,
			Order: pb.SortingOrder_ASC,
		}
	}
	if page.Cursor == nil {
		page.Cursor = &pb.Cursor{}
	}
	return page
}
