package mediator

// Mediator : 相談役
type Mediator interface {
	createColleagues()
	colleagueChanged()
}

// Colleague : メンバー
type Colleague interface {
	setMediator(Mediator)
	setColleagueEnabled(bool)
}

type mediator struct {
	ColleagueA *colleagueA
	ColleagueB *colleagueB
}

func NewMediator() *mediator {
	obj := &mediator{}
	// NOTE : 生成パターンはいろいろありそう
	obj.createColleagues() // メンバーをセットする

	return obj
}

func (m *mediator) createColleagues() {
	m.ColleagueA = &colleagueA{}
	m.ColleagueB = &colleagueB{}

	m.ColleagueA.setMediator(m)
	m.ColleagueB.setMediator(m)
}

// colleagueChanged : ここで調整する
func (m *mediator) colleagueChanged() {
	if m.ColleagueA.enabled {
		m.ColleagueB.setColleagueEnabled(true)
	} else {
		m.ColleagueB.setColleagueEnabled(false)
	}
}

type colleagueA struct {
	enabled  bool
	mediator Mediator
}

func (c *colleagueA) setMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *colleagueA) setColleagueEnabled(enabled bool) {
	c.enabled = enabled
}

func (c *colleagueA) SetEnabled(enabled bool) {
	c.enabled = enabled

	// NOTE : 状態を変更したら相談役に報告する
	defer c.mediator.colleagueChanged()
}

type colleagueB struct {
	enabled  bool
	mediator Mediator
}

func (c *colleagueB) setMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *colleagueB) setColleagueEnabled(enabled bool) {
	c.enabled = enabled
}

func (c *colleagueB) Enabled() bool {
	return c.enabled
}
