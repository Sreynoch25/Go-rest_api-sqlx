package notification_repository

import (
	"fmt"
	notification_model "marketing/src/models/notification"

	"github.com/jmoiron/sqlx"
)

type NotificationRepository interface {
	Create(req *notification_model.CreateNotificationRequest) (*notification_model.Notification, error)
	Update(id int, req *notification_model.UpdateNotificationRequest) (*notification_model.Notification, error)
	Show(page, perPage int) (*notification_model.NotificationsResponse, error)
	Delete(id, deletedBy int) error
}

type notificationRepository struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (repo *notificationRepository) Create(req *notification_model.CreateNotificationRequest) (*notification_model.Notification, error) {
	query := `
        INSERT INTO tbl_users_notifications (
            user_id, context, subject, description, icon_id, 
            notification_type_id, status_id, "order", created_by, created_at 
        ) VALUES (
            :user_id, :context, :subject, :description, :icon_id,
            :notification_type_id, :status_id, :order, :created_by, NOW()
        ) RETURNING *`

	notification := &notification_model.Notification{}
	rows, err := repo.db.NamedQuery(query, req)
	if err != nil {
		return nil, fmt.Errorf("error creating notification: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(notification); err != nil {
			return nil, fmt.Errorf("error scanning notification: %w", err)
		}
	}

	return notification, nil
}

func (repo *notificationRepository) Update(id int, req *notification_model.UpdateNotificationRequest) (*notification_model.Notification, error) {
	query := `
        UPDATE tbl_users_notifications SET
            user_id = :user_id,
            context = :context,
            subject = :subject,
            description = :description,
            icon_id = :icon_id,
            notification_type_id = :notification_type_id,
            status_id = :status_id,
            "order" = :order,
            updated_by = :updated_by,
            updated_at = NOW()
        WHERE id = :id
        RETURNING *`

	notification := &notification_model.Notification{}
	rows, err := repo.db.NamedQuery(query, req)
	if err != nil {
		return nil, fmt.Errorf("error updating notification: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(notification); err != nil {
			return nil, fmt.Errorf("error scanning notification: %w", err)
		}
	}

	return notification, nil
}

func (r *notificationRepository) Delete(id, deletedBy int) error {
	query := `
    UPDATE tbl_users_notifications SET
        deleted_by = $1,
        deleted_at = NOW()
    WHERE id = $2`

	// Use positional parameters in the Exec method
	_, err := r.db.Exec(query, deletedBy, id)
	if err != nil {
		return fmt.Errorf("error deleting notification: %w", err)
	}

	// if err != nil {
	// 	return fmt.Errorf("error deleting notification: %w", err)
	// }

	return nil
}

// func (repo *notificationRepository) Show(page, perPage int) (*notification_model.NotificationsResponse, error) {
//     query := `
//         SELECT * FROM tbl_users_notifications
//         LIMIT :limit OFFSET :offset`

//     notifications := []notification_model.Notification{}
//     err := repo.db.Select(&notifications, query, map[string]interface{}{
//         "limit":  perPage,
//         "offset": (page - 1) * perPage,
//     })
//     if err != nil {
//         return nil, fmt.Errorf("error fetching notifications: %w", err)
//     }

//     total := len(notifications)
//     return &notification_model.NotificationsResponse{
//         Notifications: notifications,
//         Total:         total,
//     }, nil
// }

func (repo *notificationRepository) Show(page, perPage int) (*notification_model.NotificationsResponse, error) {
	offset := (page - 1) * perPage
	var total int

	countQuery := ` SELECT COUNT(*) FROM tbl_users_notifications WHERE deleted_at IS NULL`
	err := repo.db.Get(&total, countQuery)
	if err != nil {
		return nil, fmt.Errorf("error counting users: %v", err)
	}

	query := `
    SELECT *FROM tbl_users_notifications
    WHERE deleted_at IS NULL
    ORDER BY created_at DESC
    LIMIT $1 OFFSET $2
    `
	var notifications []notification_model.Notification

	err = repo.db.Select(&notifications, query, perPage, offset)

	if err != nil {
		return nil, fmt.Errorf("error fetching notifications: %v", err)
	}

	return &notification_model.NotificationsResponse{
		Notifications: notifications,
		Total:         total,
	}, nil

}

func (repo *notificationRepository) ShowOne(id int) (*notification_model.Notification, error) {
	query := `
        SELECT * FROM tbl_users_notifications
        WHERE id = :id`

	notification := &notification_model.Notification{}
	err := repo.db.Get(notification, query, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching notification: %w", err)
	}

	return notification, nil
}
