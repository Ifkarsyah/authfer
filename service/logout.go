package service

import (
	"github.com/Ifkarsyah/authfer/model"
)

type LogoutParams struct {
	AccessDetails *model.AccessDetails
}

func (h *Service) Logout(u *LogoutParams) error {
	deleted, err := h.Redis.RedisDeleteAuth(u.AccessDetails.AccessUuid)
	if err != nil || deleted == 0 { //if any goes wrong
		return err
	}
	return nil
}
