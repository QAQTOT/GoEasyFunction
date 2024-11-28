package carbon

func (this Carbon) AddDays(days int) Carbon {
	this.Time.AddDate(0, 0, days)
	return this
}

func (this Carbon) AddYears(years int) Carbon {
	this.Time.AddDate(years, 0, 0)
	return this
}

func (this Carbon) AddMonths(months int) Carbon {
	this.Time.AddDate(0, months, 0)
	return this
}
