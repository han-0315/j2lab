package gitlabx

import (
	"fmt"
	"net/http"
	"time"

	gitlab "github.com/xanzy/go-gitlab"
)

type CreateEpicOptions struct {
	Title            *string         `url:"title,omitempty" json:"title,omitempty"`
	Description      *string         `url:"description,omitempty" json:"description,omitempty"`
	Labels           *gitlab.Labels  `url:"labels,comma,omitempty" json:"labels,omitempty"`
	StartDateIsFixed *bool           `url:"start_date_is_fixed,omitempty" json:"start_date_is_fixed,omitempty"`
	StartDateFixed   *gitlab.ISOTime `url:"start_date_fixed,omitempty" json:"start_date_fixed,omitempty"`
	DueDateIsFixed   *bool           `url:"due_date_is_fixed,omitempty" json:"due_date_is_fixed,omitempty"`
	DueDateFixed     *gitlab.ISOTime `url:"due_date_fixed,omitempty" json:"due_date_fixed,omitempty"`

	//* 라이브러리에서 지원하지 않는 추가 옵션
	Color        *string    `url:"color,omitempty" json:"color,omitempty"`
	Confidential *bool      `url:"confidential,omitempty" json:"confidential,omitempty"`
	CreatedAt    *time.Time `url:"created_at,omitempty" json:"created_at,omitempty"`
	// ParentID ...
}

func CreateEpic(gl *gitlab.Client, gid interface{}, opt *CreateEpicOptions) (*gitlab.Epic, *gitlab.Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/epics", gitlab.PathEscape(group))

	req, err := gl.NewRequest(http.MethodPost, u, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	e := new(gitlab.Epic)
	resp, err := gl.Do(req, e)
	if err != nil {
		return nil, resp, err
	}

	return e, resp, nil
}
