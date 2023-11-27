package geoquery

import (
	"encoding/json"
	"net/http"

	"github.com/Fancyyy21/geoquery/gq"
	"github.com/Fancyyy21/geoquery/helper"
	"github.com/Fancyyy21/geoquery/models"
)

func ReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// func PostGeoIntersects(w http.ResponseWriter, r *http.Request) {
// 	var msg model.IteungMessage
// 	var resp atmessage.Response
// 	json.NewDecoder(r.Body).Decode(&msg)
// 	if r.Header.Get("Secret") == config.EndpointSecret {
// 		resp.Response = gq.GeoIntersects(config.Mongocon, msg.Longitude, msg.Latitude)
// 	} else {
// 		resp.Response = "Secret Salah"
// 	}
// 	fmt.Fprintf(w, helper.Json(resp))
// }

func PostGeoIntersects(mongoenv, dbname string, r *http.Request) string {
	var longlat models.LongLat
	var response models.Pesan
	response.Status = false
	mconn := helper.SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&longlat)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		response.Message = gq.GeoIntersects(mconn, longlat.Longitude, longlat.Latitude)
	}
	return ReturnStruct(response)
}
