package notification_model

import "time"

type Notification struct {
	ID                 int        `json:"id" db:"id"`
	UserID             int        `json:"user_id" db:"user_id"`
	Context            string     `json:"context" db:"context"`
	Subject            string     `json:"subject" db:"subject"`
	Description        string     `json:"description" db:"description"`
	IconID             int        `json:"icon_id" db:"icon_id"`
	NotificationtypeID int        `json:"notification_type_id" db:"notification_type_id"`
	StatusID           int        `json:"status_id" db:"status_id"`
	Order              int        `json:"order" db:"order"`
	CreatedBy          int        `json:"created_by" db:"created_by"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy          *int        `json:"updated_by" db:"updated_by"`
	UpdatedAt          *time.Time  `json:"updated_at" db:"updated_at"`
	DeletedBy          *int       `json:"deleted_by" db:"deleted_by"`
	DeletedAt          *time.Time `json:"deleted_at" db:"deleted_at"`
}

type CreateNotificationRequest struct {
	UserID             int    `json:"user_id" db:"user_id"`
	Context            string `json:"context" db:"context"`
	Subject            string `json:"subject" db:"subject"`
	Description        string `json:"description" db:"description"`
	IconID             int    `json:"icon_id" db:"icon_id"`
	NotificationtypeID int    `json:"notification_type_id" db:"notification_type_id"`
	StatusID           int    `json:"status_id" db:"status_id"`
	Order              int    `json:"order" db:"order"`
	CreatedBy          int    `json:"created_by" db:"created_by"`
}

type UpdateNotificationRequest struct {
	UserID             int    `json:"user_id" db:"user_id"`
	Context            string `json:"context" db:"context"`
	Subject            string `json:"subject" db:"subject"`
	Description        string `json:"description" db:"description"`
	IconID             int    `json:"icon_id" db:"icon_id"`
	NotificationtypeID int    `json:"notification_type_id" db:"notification_type_id"`
	StatusID           int    `json:"status_id" db:"status_id"`
	Order              int    `json:"order" db:"order"`
	UpdatedBy          int    `json:"updated_by" db:"updated_by"`
}

type NotificationsResponse struct {
	Notifications []Notification `json:"notifications"`
	Total         int            `json:"total"`
}

type NotificationResponse struct {
	Notification *Notification `json:"notification"`
}
