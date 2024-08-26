package repository

import (
	"coinkeeper/errs"
	"errors"
	"github.com/jinzhu/gorm"
)

func translateError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	return err
}
