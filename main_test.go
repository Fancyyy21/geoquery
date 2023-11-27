package geoquery

import (
	"fmt"
	"testing"

	"github.com/Fancyyy21/geoquery/gq"
	"github.com/petapedia/geoquery/helper"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("mongodb+srv://Fancy21:Acermaulana21@maulana.uiq9nmy.mongodb.net/?retryWrites=true&w=majority", "gis")
	datagedung := gq.GeoIntersects(mconn, 108.22501803948785, -6.835460702759789)
	fmt.Println(datagedung)
}
