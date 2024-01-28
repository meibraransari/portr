package service

import (
	"context"

	db "github.com/amalshaji/portr/internal/server/db/models"
)

func (s *Service) ListSettings(ctx context.Context) db.GlobalSetting {
	settings, _ := s.db.Queries.GetGlobalSettings(ctx)
	return settings
}

func (s *Service) UpdateEmailSettings(ctx context.Context, updateSettingsInput UpdateEmailSettingsInput) (db.GlobalSetting, error) {
	err := s.db.Queries.UpdateGlobalSettings(ctx, db.UpdateGlobalSettingsParams{
		SmtpEnabled:            updateSettingsInput.SmtpEnabled,
		SmtpHost:               updateSettingsInput.SmtpHost,
		SmtpPort:               updateSettingsInput.SmtpPort,
		SmtpUsername:           updateSettingsInput.SmtpUsername,
		SmtpPassword:           updateSettingsInput.SmtpPassword,
		FromAddress:            updateSettingsInput.FromAddress,
		AddMemberEmailSubject:  updateSettingsInput.AddMemberEmailSubject,
		AddMemberEmailTemplate: updateSettingsInput.AddMemberEmailTemplate,
	})
	if err != nil {
		return db.GlobalSetting{}, err
	}
	return s.ListSettings(ctx), nil
}
