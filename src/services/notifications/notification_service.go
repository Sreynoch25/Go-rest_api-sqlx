package notifications_service

import (
	notification_model "marketing/src/models/notification"
	notification_repository "marketing/src/repositeries/notification"
)

type NotificationService interface{
	Create(req *notification_model.CreateNotificationRequest) (*notification_model.Notification, error) 
	Update(id int, req *notification_model.UpdateNotificationRequest) (*notification_model.Notification, error)
	Show(page, pageSize int) (*notification_model.NotificationsResponse , error) 
}

type notificationService struct {
	// repo notification_repositery.NotificationRepository
	repo notification_repository.NotificationRepository
} 

func NewNotificationService(repo notification_repository.NotificationRepository) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) Create(req *notification_model.CreateNotificationRequest) (*notification_model.Notification, error) {
	return s.repo.Create(req)
}

func (s *notificationService) Update(id int, req *notification_model.UpdateNotificationRequest) (*notification_model.Notification, error) {
	return s.repo.Update(id, req)
}

func(s *notificationService) Show(page, pageSize int) (*notification_model.NotificationsResponse , error) {
	offset := (page -1 ) * pageSize
	return s.repo.Show(offset, pageSize)
} 