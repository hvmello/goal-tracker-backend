package goals

import "errors"

// Validation constants
const (
	MinProgress = 0
	MaxProgress = 100
)

// ValidationErrors
var (
	ErrTitleRequired   = errors.New("title is required")
	ErrInvalidProgress = errors.New("progress must be between 0 and 100")
)

// ValidateGoal centralizes all goal validations
func ValidateGoal(goal *Goal) error {
	if goal.Title == "" {
		return ErrTitleRequired
	}
	if goal.Progress < MinProgress || goal.Progress > MaxProgress {
		return ErrInvalidProgress
	}
	return nil
}
