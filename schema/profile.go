package schema

type (
	Profile struct {
		Result struct {
			IsAnonymous     *bool           `json:"is_anonymous"`
			AllowExplicit   bool            `json:"allow_explicit"`
			Birthday        int64           `json:"birthday"`
			Created         int64           `json:"created"`
			Email           string          `json:"email"`
			ExternalProfile ExternalProfile `json:"external_profile"`
			Gender          string          `json:"gender"`
			HideBirthday    bool            `json:"hide_birthday"`
			HideGender      bool            `json:"hide_gender"`
			ID              int             `json:"id"`
			Image           Image           `json:"image"`
			IsActive        bool            `json:"is_active"`
			IsAgreement     bool            `json:"is_agreement"`
			IsEditor        bool            `json:"is_editor"`
			IsRegistered    bool            `json:"is_registered"`
			Name            string          `json:"name"`
			Phone           string          `json:"phone"`
			Registered      int64           `json:"registered"`
			Token           string          `json:"token"`
			Username        string          `json:"username"`
		} `json:"result"`
	}

	ExternalProfile struct {
		Birthday   int64  `json:"birthday"`
		Email      string `json:"email"`
		ExternalID string `json:"external_id"`
		FirstName  string `json:"first_name"`
		Gender     string `json:"gender"`
		LastName   string `json:"last_name"`
		MiddleName string `json:"middle_name"`
		Phone      string `json:"phone"`
		Type       string `json:"type"`
	}
)
