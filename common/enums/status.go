package enums

// STATUS enum
type Status string

const (
	StatusPending  Status = "PENDING"
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
)

// STATUS_INVITATION enum
type StatusInvitation string

const (
	StatusInvitationPending  StatusInvitation = "PENDING"
	StatusInvitationAccepted StatusInvitation = "ACCEPTED"
	StatusInvitationRejected StatusInvitation = "REJECTED"
)

// STATUS_EXPORTABLE_DOCUMENT enum
type StatusExportableDocument string

const (
	StatusExportableDocumentPending         StatusExportableDocument = "PENDING"
	StatusExportableDocumentApproved        StatusExportableDocument = "APPROVED"
	StatusExportableDocumentRejected        StatusExportableDocument = "REJECTED"
	StatusExportableDocumentChangesRequired StatusExportableDocument = "CHANGES_REQUIRED"
	StatusExportableDocumentChangesMade     StatusExportableDocument = "CHANGES_MADE"
)

// ORDER_FILE_STATUS enum
type OrderFileStatus string

const (
	OrderFileStatusPending  OrderFileStatus = "PENDING"
	OrderFileStatusUploaded OrderFileStatus = "UPLOADED"
	OrderFileStatusModified OrderFileStatus = "MODIFIED"
)
