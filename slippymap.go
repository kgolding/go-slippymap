package slippymap

import (
	"fmt"
	"math"
)

// LatLng2TileZoomXY takes a latitude, longtitude and zoom level and returns a slippy map x,y tile cord
func LatLng2TileZoomXY(lat float64, lng float64, zoom int) (int, int, error) {
	if zoom < 0 || zoom > 30 {
		return 0, 0, fmt.Errorf("invalid zoom %d", zoom)
	}
	if lat < -90 || lat > 90 {
		return 0, 0, fmt.Errorf("invalid latitude %f", lat)
	}
	if lng < -180 || lng > 180 {
		return 0, 0, fmt.Errorf("invalid longtitude %f", lng)
	}
	x := math.Floor((lng + 180.0) / 360.0 * (math.Exp2(float64(zoom))))
	y := math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(zoom))))

	return int(x), int(y), nil
}
