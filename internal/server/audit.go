package server

import (
	"context"

	audit "github.com/AndrewMislyuk/go-shop-logger/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (audit *audit.Empty, err error) {
	err = h.service.Insert(ctx, req)

	return audit, err
}
