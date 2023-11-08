package models

type NewJob struct {
	Company         Company `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint    `json:"cid"`
	Name            string  `json:"name" validate:"required"`
	Budget          string  `json:"budget" validate:"required"`
	MinNoticePeriod string  `json:"minNoticePeriod" validate:"required"`
	MaxNoticePeriod string  `json:"maxNoticePeriod" validate:"required"`
	JobLocation     []uint  `json:"jobLocation" gorm:"many2many:job_location"`
	Technology      []uint  `json:"technology" gorm:"many2many:job_technology"`
	Description     string  `json:"description" validate:"required"`
	MinExp          string  `json:"minExp" validate:"required"`
	MaxExp          string  `json:"maxExp" validate:"required"`
	Qualifications  []uint  `json:"qualification" gorm:"many2many:job_qualification"`
	Shift           []uint  `json:"shift" gorm:"many2many:job_shift"`
	JobType         string  `json:"jobType" validate:"required"`
}

type ResponseJob struct {
	Id uint `json:"responseId"`
}

// type NewLocation struct {
// 	PlaceId uint `json:"id"`
// }

// type NewTechnology struct {
// 	TechnologyId uint `json:"id"`
// }

// type NewQualification struct {
// 	QualificationId uint `json:"id"`
// }

// type NewShift struct {
// 	ShiftId uint `json:"id"`
// }
