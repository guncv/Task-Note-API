package utils

import (
	"strings"
	"time"

	constants "github.com/guncv/tech-exam-software-engineering/constant"
	"github.com/guncv/tech-exam-software-engineering/entities"
)

type FieldError map[string]string

func ValidateCreateTaskInput(input entities.CreateTaskRequest) interface{} {
	var errs []FieldError

	if isEmpty(input.Title) {
		errs = append(errs, newFieldError("title", "Title is required"))
	} else if exceedsMaxLength(input.Title, 100) {
		errs = append(errs, newFieldError("title", "Title must not exceed 100 characters"))
	}

	if isEmpty(string(input.Status)) {
		errs = append(errs, newFieldError("status", "Status is required"))
	} else if isInvalidStatus(input.Status) {
		errs = append(errs, newFieldError("status", "Status must be IN_PROGRESS or COMPLETED"))
	}

	if isZeroTime(input.Date) {
		errs = append(errs, newFieldError("date", "Date is required and must be RFC3339 format"))
	}

	return returnIfErrors(errs)
}

func ValidateUpdateTaskInput(input entities.UpdateTaskRequest) interface{} {
	var errs []FieldError

	if !isEmpty(input.Title) && exceedsMaxLength(input.Title, 100) {
		errs = append(errs, newFieldError("title", "Title must not exceed 100 characters"))
	}

	if !isEmpty(string(input.Status)) && isInvalidStatus(input.Status) {
		errs = append(errs, newFieldError("status", "Status must be IN_PROGRESS or COMPLETED"))
	}

	return returnIfErrors(errs)
}

func ValidateLoginInput(input entities.LoginRequest) interface{} {
	var errs []FieldError

	if isEmpty(input.Email) {
		errs = append(errs, newFieldError("email", "Email is required"))
	} else if isInvalidEmail(input.Email) {
		errs = append(errs, newFieldError("email", "Email is invalid"))
	}

	if isEmpty(input.Password) {
		errs = append(errs, newFieldError("password", "Password is required"))
	} else if belowMinLength(input.Password, 8) {
		errs = append(errs, newFieldError("password", "Password must be at least 8 characters"))
	}

	return returnIfErrors(errs)
}

func ValidateRegisterInput(input entities.RegisterRequest) interface{} {
	var errs []FieldError

	if isEmpty(input.FirstName) {
		errs = append(errs, newFieldError("first_name", "First name is required"))
	}

	if isEmpty(input.LastName) {
		errs = append(errs, newFieldError("last_name", "Last name is required"))
	}

	if isEmpty(input.Email) {
		errs = append(errs, newFieldError("email", "Email is required"))
	} else if isInvalidEmail(input.Email) {
		errs = append(errs, newFieldError("email", "Email is invalid"))
	}

	if isEmpty(input.Password) {
		errs = append(errs, newFieldError("password", "Password is required"))
	} else if belowMinLength(input.Password, 8) {
		errs = append(errs, newFieldError("password", "Password must be at least 8 characters"))
	}

	return returnIfErrors(errs)
}

func ValidateGetAllTasksInput(input entities.GetAllTasksRequest) interface{} {
	var errs []FieldError

	if exceedsMaxLength(input.Search, 100) {
		errs = append(errs, newFieldError("search", "Search must not exceed 100 characters"))
	}

	if isInvalidOrder(input.Order) {
		errs = append(errs, newFieldError("order", "Order must be asc or desc"))
	}

	if isInvalidSortBy(input.SortBy) {
		errs = append(errs, newFieldError("sort_by", "Sort by must be title, created_at, or status"))
	}

	if input.Limit < 1 {
		errs = append(errs, newFieldError("limit", "Limit must be greater than 0"))
	}

	if input.Offset < 1 {
		errs = append(errs, newFieldError("offset", "Offset must be greater than 0"))
	}

	if input.Limit > 100 {
		errs = append(errs, newFieldError("limit", "Limit must not exceed 100"))
	}

	return returnIfErrors(errs)
}

func newFieldError(field, message string) FieldError {
	return FieldError{"field": field, "message": message}
}

func returnIfErrors(errs []FieldError) interface{} {
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func isInvalidEmail(s string) bool {
	return !strings.Contains(s, "@")
}

func exceedsMaxLength(s string, max int) bool {
	return len(s) > max
}

func belowMinLength(s string, min int) bool {
	return len(s) < min
}

func isInvalidStatus(s constants.TaskStatus) bool {
	return s != constants.TaskStatusPending && s != constants.TaskStatusCompleted
}

func isZeroTime(t time.Time) bool {
	return t.IsZero()
}

func isInvalidOrder(s string) bool {
	if s == "" {
		return false
	}
	return s != "asc" && s != "desc"
}

func isInvalidSortBy(s string) bool {
	if s == "" {
		return false
	}
	return s != "title" && s != "created_at" && s != "status"
}
