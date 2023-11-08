package models

type NewJob struct {
	Company         Company         `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint            `json:"cid"`
	Name            string          `json:"name" validate:"required"`
	Budget          string          `json:"budget" validate:"required"`
	MinNoticePeriod string          `json:"min_notice_period" validate:"required"`
	MaxNoticePeriod string          `json:"max_notice_period" validate:"required"`
	JobLocation     []Location      `json:"job_location" gorm:"many2many:job_location"`
	Technology      []Technology    `json:"technology" gorm:"many2many:job_technology"`
	Description     string          `json:"description" validate:"required"`
	MinExp          string          `json:"min_exp" validate:"required"`
	MaxExp          string          `json:"max_exp" validate:"required"`
	Qualifications  []Qualification `json:"qualification" gorm:"many2many:job_qualification"`
	Shift           []Shift         `json:"shift" gorm:"many2many:job_shift"`
	JobType         string          `json:"job_type" validate:"required"`
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
