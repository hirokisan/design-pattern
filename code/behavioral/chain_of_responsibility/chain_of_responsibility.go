package chain_of_responsibility

type IssueLevel string

const (
	IssueLevelStaff   = IssueLevel("staff")
	IssueLevelManager = IssueLevel("manager")
	IssueLevelUnknown = IssueLevel("unknown")
)

type Issue struct {
	IssueLevel IssueLevel
	Content    string
}

func NewIssue(level IssueLevel, content string) *Issue {
	// NOTE : クライアントがIssueを作成するのはAPIとして不便なので、実際にはクライアントからの要請を受けてIssueを作成するオブジェクトがいるはず
	return &Issue{
		IssueLevel: level,
		Content:    content,
	}
}

type Role interface {
	solve(issue Issue) bool
	Handle(issue Issue) string
}

type Staff struct {
	next Role
}

func NewStaff() Role {
	// NOTE : nextはstaffが決めるわけではなくてcompany的なfactoryが注入する形が良さそう
	return &Staff{
		next: NewManager(),
	}
}

func (s *Staff) solve(issue Issue) bool {
	return true
}

func (s *Staff) done() string {
	return "solved"
}

func (s *Staff) fail() string {
	return "not solved"
}

// Handle : 場合によっては再帰的な処理になるかも
func (s *Staff) Handle(issue Issue) string {
	if issue.IssueLevel == IssueLevelStaff {
		if s.solve(issue) {
			return s.done()
		}
		return s.fail()
	}
	if s.next != nil {
		return s.next.Handle(issue)
	}

	return s.fail()
}

type Manager struct {
}

func NewManager() Role {
	return &Manager{}
}

func (m *Manager) solve(issue Issue) bool {
	return true
}

func (s *Manager) done() string {
	return "solved"
}

func (s *Manager) fail() string {
	return "not solved"
}

func (m *Manager) Handle(issue Issue) string {
	if issue.IssueLevel == IssueLevelManager {
		if m.solve(issue) {
			return m.done()
		}
	}

	return m.fail()
}
