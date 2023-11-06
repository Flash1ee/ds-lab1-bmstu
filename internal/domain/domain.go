package domain

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"not null"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

type PersonDTO struct {
	ID      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age,omitempty" validate:"numeric"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

func (*Person) TableName() string {
	return "person"
}

func (obj *Person) ToDTO() (*PersonDTO, error) {
	dto := &PersonDTO{}
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonStr, dto); err != nil {
		return nil, err
	}
	return dto, nil
}

func (obj *Person) Validate() error {
	if obj.Name == "" {
		return InvalidPerson
	} else if obj.Age < 0 {
		return InvalidPerson
	} else {
		return nil
	}
}

func (*Person) ArrToDTO(src []Person) ([]PersonDTO, error) {
	dst := make([]PersonDTO, len(src))
	for k, v := range src {
		data, err := v.ToDTO()
		if err != nil {
			return nil, err
		}
		dst[k] = *data
	}
	return dst, nil
}

func (obj *PersonDTO) ToModel() (*Person, error) {
	model := &Person{}

	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonStr, model); err != nil {
		return nil, err
	}
	return model, nil
}

func (obj *PersonDTO) Validate() error {
	if obj.Name == "" {
		return InvalidRequest
	} else if obj.Age < 0 {
		return InvalidRequest
	} else {
		return nil
	}
}
