package internal

import "gitlab.com/Spide_IT/spide_it/email/internal/model"

type EmailsRepository interface {
	CreateEmail(el model.Email) (model.Email,error)
}