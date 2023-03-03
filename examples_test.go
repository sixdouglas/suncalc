// get today's sunlight times for London
package suncalc_test

import (
	"fmt"
	"math"
	"time"

	"github.com/sixdouglas/suncalc"
)

func ExampleGetTimes() {
	var now = time.Date(2005, 6, 1, 12, 0, 0, 0, time.UTC)

	// Location to query for
	lat, long := 51.5, -0.1

	// get the times for today, latitude, longitude, height below or above the
	// horizon, and in timezone
	var times = suncalc.GetTimes(now, lat, long)

	for _, timeOfDay := range suncalc.DayTimeNames {
		oneTime := times[timeOfDay]
		if !oneTime.Value.IsZero() {
			fmt.Printf("%-13s %s\n", string(oneTime.Name),
				oneTime.Value.Format("2006-01-02 15:04:05"))
		}
	}

	// format sunrise time from the Date object
	fmt.Printf("Sunrise / Sunset time: %s / %s\n",
		times[suncalc.Sunrise].Value.Format("15:04:05"),
		times[suncalc.Sunset].Value.Format("15:04:05"),
	)

	// get position of the sun (azimuth and altitude) at today's sunrise
	var sunrisePos = suncalc.GetPosition(times[suncalc.Sunrise].Value, lat, long)

	// get sunrise azimuth in degrees
	var sunriseAzimuth = sunrisePos.Azimuth * 180 / math.Pi
	fmt.Printf("Sunrise Azimuth: %f deg\n", sunriseAzimuth)

	// get position of the sun (azimuth and altitude) at today's sunset
	var sunsetPos = suncalc.GetPosition(times[suncalc.Sunset].Value, lat, long)

	// get the sunset azimuth in degrees
	var sunsetAzimuth = sunsetPos.Azimuth * 180 / math.Pi
	fmt.Printf("Sunset Azimuth: %f deg\n", sunsetAzimuth)

	// get current position of the sun (azimuth and altitude) at today's sunrise
	var sunPos = suncalc.GetPosition(now, lat, long)
	fmt.Printf("Sun Azimuth: %f deg\n", sunPos.Azimuth*180/math.Pi)
	fmt.Printf("Sun Altitude: %f deg\n", sunPos.Altitude*180/math.Pi)
	// Output:
	// nauticalDawn  2005-06-01 01:57:31
	// dawn          2005-06-01 03:04:43
	// sunrise       2005-06-01 03:50:12
	// sunriseEnd    2005-06-01 03:54:33
	// goldenHourEnd 2005-06-01 04:42:52
	// goldenHour    2005-06-01 19:16:35
	// sunsetStart   2005-06-01 20:04:54
	// sunset        2005-06-01 20:09:15
	// dusk          2005-06-01 20:54:44
	// nauticalDusk  2005-06-01 22:01:56
	// Sunrise / Sunset time: 03:50:12 / 20:09:15
	// Sunrise Azimuth: -128.374196 deg
	// Sunset Azimuth: 128.610491 deg
	// Sun Azimuth: 0.351955 deg
	// Sun Altitude: 60.593709 deg
}

func ExampleGetTimesWithObserver() {
	var now = time.Date(2012, 12, 12, 12, 0, 0, 0, time.UTC)
	// Use time.Now() for today

	// Location to query for
	lat, long := 51.5, -0.1

	// get the times for today, latitude, longitude, height below or above the
	// horizon, and in timezone
	var observer = suncalc.Observer{
		Latitude:  lat,
		Longitude: long,
		Height:    0,
		Location:  time.UTC,
	}

	var times = suncalc.GetTimesWithObserver(now, observer)

	for _, timeOfDay := range suncalc.DayTimeNames {
		oneTime := times[timeOfDay]
		if !oneTime.Value.IsZero() {
			fmt.Printf("%-13s %s\n", string(oneTime.Name),
				oneTime.Value.Format("2006-01-02 15:04:05"))
		}
	}

	// format sunrise time from the Date object
	fmt.Printf("Sunrise / Sunset time: %s / %s\n",
		times[suncalc.Sunrise].Value.Format("15:04:05"),
		times[suncalc.Sunset].Value.Format("15:04:05"),
	)

	// get position of the sun (azimuth and altitude) at today's sunrise
	var sunrisePos = suncalc.GetPosition(times[suncalc.Sunrise].Value, lat, long)

	// get sunrise azimuth in degrees
	var sunriseAzimuth = sunrisePos.Azimuth * 180 / math.Pi
	fmt.Printf("Sunrise Azimuth: %f deg\n", sunriseAzimuth)

	// get position of the sun (azimuth and altitude) at today's sunset
	var sunsetPos = suncalc.GetPosition(times[suncalc.Sunset].Value, lat, long)

	// get the sunset azimuth in degrees
	var sunsetAzimuth = sunsetPos.Azimuth * 180 / math.Pi
	fmt.Printf("Sunset Azimuth: %f deg\n", sunsetAzimuth)

	// get current position of the sun (azimuth and altitude) at today's sunrise
	var sunPos = suncalc.GetPosition(now, lat, long)
	fmt.Printf("Sun Azimuth: %f deg\n", sunPos.Azimuth*180/math.Pi)
	fmt.Printf("Sun Altitude: %f deg\n", sunPos.Altitude*180/math.Pi)
	// Output:
	// nightEnd      2012-12-12 05:54:58
	// nauticalDawn  2012-12-12 06:35:38
	// dawn          2012-12-12 07:18:37
	// sunrise       2012-12-12 07:58:39
	// sunriseEnd    2012-12-12 08:03:00
	// goldenHourEnd 2012-12-12 08:59:36
	// goldenHour    2012-12-12 14:51:47
	// sunsetStart   2012-12-12 15:48:23
	// sunset        2012-12-12 15:52:45
	// dusk          2012-12-12 16:32:47
	// nauticalDusk  2012-12-12 17:15:46
	// night         2012-12-12 17:56:26
	// Sunrise / Sunset time: 07:58:39 / 15:52:45
	// Sunrise Azimuth: -52.101764 deg
	// Sunset Azimuth: 52.356679 deg
	// Sun Azimuth: 1.187845 deg
	// Sun Altitude: 15.381733 deg
}
