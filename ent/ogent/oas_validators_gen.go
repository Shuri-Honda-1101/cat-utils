// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *CatCreate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CatCreateSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *CatList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CatListSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *CatRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CatReadSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *CatToiletsList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CatToiletsListType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *CatUpdate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CatUpdateSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *CreateCatReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateCatReqSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *CreateToiletReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateToiletReqType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s ListCatOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ListCatToiletsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ListToiletOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ListUserCatsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ListUserOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s *ToiletCatRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ToiletCatReadSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *ToiletCreate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ToiletCreateType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *ToiletList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ToiletListType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *ToiletRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ToiletReadType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *ToiletUpdate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ToiletUpdateType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *UpdateCatReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Sex.Set {
			if err := func() error {
				if err := s.Sex.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UpdateCatReqSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *UpdateToiletReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Type.Set {
			if err := func() error {
				if err := s.Type.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UpdateToiletReqType) Validate() error {
	switch s {
	case "pee":
		return nil
	case "poo":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s *UserCatsList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Sex.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sex",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserCatsListSex) Validate() error {
	switch s {
	case "male":
		return nil
	case "female":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}