package carbon

import (
	"time"
)

type Carbon struct {
	Time time.Time
}

func Now() Carbon {
	this := Carbon{}
	this.Time = time.Now()
	this.SetLocale(DefaultLocale)
	return this
}

func (this Carbon) SetLocale(local string) Carbon {
	location, err := time.LoadLocation(local)
	if err != nil {
		panic(err)
	}
	this.Time = this.Time.In(location)
	return this
}

func (this Carbon) GetDateTimeString() string {
	dateTime := this.Time.Format(TimeFormatDateTime)
	return dateTime
}

func (this Carbon) GetDateString() string {
	date := this.Time.Format(TimeFormatDate)
	return date
}
func (this Carbon) GetTimeString() string {
	format := this.Time.Format(TimeFormatTime)
	return format
}

func (this Carbon) GetUnixTimeStamp() int64 {
	unix := this.Time.Unix()
	return unix
}

func (this Carbon) GetUnixMicroTimeStamp() int64 {
	micro := this.Time.UnixMicro()
	return micro
}

func (this Carbon) GetUnixNanoTimeStamp() int64 {
	nano := this.Time.UnixNano()
	return nano
}
