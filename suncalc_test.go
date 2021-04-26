package suncalc

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTimes(t *testing.T) {
	type args struct {
		date   time.Time
		lat    float64
		lng    float64
		height float64
	}
	tests := []struct {
		name string
		args args
		want map[DayTimeName]DayTime
	}{
		{
			"2020-05-17 15:05:16.414278 +0700",
			args{
				date:   time.Date(2020, 5, 17, 15, 05, 16, 414278, time.FixedZone("Novosibirsk", 7)),
				lat:    55.007379,
				lng:    82.956132,
				height: 0,
			},
			map[DayTimeName]DayTime{
				Dawn:          {Dawn, NullTime{Time: time.Date(2020, 5, 16, 21, 29, 45, 597757952, time.UTC), Valid: true}},
				Dusk:          {Dusk, NullTime{Time: time.Date(2020, 5, 17, 15, 22, 12, 177163520, time.UTC), Valid: true}},
				GoldenHour:    {GoldenHour, NullTime{Time: time.Date(2020, 5, 17, 13, 37, 57, 784171520, time.UTC), Valid: true}},
				GoldenHourEnd: {GoldenHourEnd, NullTime{Time: time.Date(2020, 5, 16, 23, 13, 59, 990749952, time.UTC), Valid: true}},
				Nadir:         {Nadir, NullTime{Time: time.Date(2020, 5, 16, 18, 25, 58, 887460864, time.UTC), Valid: true}},
				NauticalDawn:  {NauticalDawn, NullTime{Time: time.Date(2020, 5, 16, 20, 16, 15, 785748736, time.UTC), Valid: true}},
				NauticalDusk:  {NauticalDusk, NullTime{Time: time.Date(2020, 5, 17, 16, 35, 41, 989172992, time.UTC), Valid: true}},

				Night:    {Night, NullTime{Time: time.Date(1677, 9, 21, 0, 12, 43, 145224192, time.UTC), Valid: false}},
				NightEnd: {NightEnd, NullTime{Time: time.Date(1677, 9, 21, 0, 12, 43, 145224192, time.UTC), Valid: false}},

				SolarNoon:   {SolarNoon, NullTime{Time: time.Date(2020, 5, 17, 6, 25, 58, 887460864, time.UTC), Valid: true}},
				Sunrise:     {Sunrise, NullTime{Time: time.Date(2020, 5, 16, 22, 18, 13, 487034624, time.UTC), Valid: true}},
				SunriseEnd:  {SunriseEnd, NullTime{Time: time.Date(2020, 5, 16, 22, 22, 50, 376259072, time.UTC), Valid: true}},
				Sunset:      {Sunset, NullTime{Time: time.Date(2020, 5, 17, 14, 33, 44, 287886848, time.UTC), Valid: true}},
				SunsetStart: {SunsetStart, NullTime{Time: time.Date(2020, 5, 17, 14, 29, 07, 398662656, time.UTC), Valid: true}},
			},
		},
		{
			"2020-05-17 15:05:16.414278 +0100",
			args{
				date:   time.Date(2020, 5, 17, 15, 05, 16, 414278, time.FixedZone("Paris", 1)),
				lat:    50.700078,
				lng:    2.891449,
				height: 0,
			},
			map[DayTimeName]DayTime{
				Dawn:          {Dawn, NullTime{Time: time.Date(2020, 5, 17, 3, 17, 8, 991038976, time.UTC), Valid: true}},
				Dusk:          {Dusk, NullTime{Time: time.Date(2020, 5, 17, 20, 15, 20, 984436480, time.UTC), Valid: true}},
				GoldenHour:    {GoldenHour, NullTime{Time: time.Date(2020, 5, 17, 18, 45, 30, 187927296, time.UTC), Valid: true}},
				GoldenHourEnd: {GoldenHourEnd, NullTime{Time: time.Date(2020, 5, 17, 4, 46, 59, 787548672, time.UTC), Valid: true}},
				Nadir:         {Nadir, NullTime{Time: time.Date(2020, 5, 16, 23, 46, 14, 987737856, time.UTC), Valid: true}},
				NauticalDawn:  {NauticalDawn, NullTime{Time: time.Date(2020, 5, 17, 2, 21, 40, 781048832, time.UTC), Valid: true}},
				NauticalDusk:  {NauticalDusk, NullTime{Time: time.Date(2020, 5, 17, 21, 10, 49, 194426880, time.UTC), Valid: true}},

				Night:    {Night, NullTime{Time: time.Date(2020, 5, 17, 22, 31, 59, 413715456, time.UTC), Valid: true}},
				NightEnd: {NightEnd, NullTime{Time: time.Date(2020, 05, 17, 1, 0, 30, 561760000, time.UTC), Valid: true}},

				SolarNoon:   {SolarNoon, NullTime{Time: time.Date(2020, 5, 17, 11, 46, 14, 987737856, time.UTC), Valid: true}},
				Sunrise:     {Sunrise, NullTime{Time: time.Date(2020, 5, 17, 3, 57, 59, 507258624, time.UTC), Valid: true}},
				SunriseEnd:  {SunriseEnd, NullTime{Time: time.Date(2020, 5, 17, 4, 1, 58, 918955264, time.UTC), Valid: true}},
				Sunset:      {Sunset, NullTime{Time: time.Date(2020, 5, 17, 19, 34, 30, 468217088, time.UTC), Valid: true}},
				SunsetStart: {SunsetStart, NullTime{Time: time.Date(2020, 5, 17, 19, 30, 31, 56520192, time.UTC), Valid: true}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimes(tt.args.date, tt.args.lat, tt.args.lng, tt.args.height, time.UTC); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
