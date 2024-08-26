package ports

import (
	"context"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/filters"
	"github.com/billowdev/document-system-field-manager/pkg/helpers/pagination"
	"github.com/billowdev/document-system-field-manager/pkg/utils"
)

type IConfigSystemMasterFileFieldRepository interface {
	GetConfigSystemMasterFileField(ctx context.Context, id uint) (*models.ConfigSystemMasterFileField, error)
	GetConfigSystemMasterFileFields(ctx context.Context, p pagination.PaginationParams[filters.ConfigSystemMasterFileFieldFilter]) (*pagination.Pagination[[]models.ConfigSystemMasterFileField], error)
	CreateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) error
	UpdateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) error
	DeleteConfigSystemMasterFileField(ctx context.Context, id uint) error
}

type IConfigSystemMasterFileFieldService interface {
	GetConfigSystemMasterFileField(ctx context.Context, id uint) utils.APIResponse
	GetConfigSystemMasterFileFields(ctx context.Context, p pagination.PaginationParams[filters.ConfigSystemMasterFileFieldFilter]) pagination.Pagination[[]models.ConfigSystemMasterFileField]
	CreateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) utils.APIResponse
	UpdateConfigSystemMasterFileField(ctx context.Context, payload *models.ConfigSystemMasterFileField) utils.APIResponse
	DeleteConfigSystemMasterFileField(ctx context.Context, id uint) utils.APIResponse
}
