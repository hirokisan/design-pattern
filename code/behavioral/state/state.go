package state

// state
// 状態を入れ替え可能にするというコンセプトは理解できる
// 入れ替え処理を状態に任せるのは気持ち悪い（状態は他の状態を知らないといけない？）

var hState = newHappyState()
var aState = newAngryState()

type State interface {
	expressFeeling() string
	setFeeling(*person, Feeling)
}

type happyState struct {
}

func newHappyState() *happyState {
	return &happyState{}
}

func getHappyState() *happyState {
	return hState
}

func (s *happyState) expressFeeling() string {
	return "happy"
}

func (s *happyState) setFeeling(person *person, feeling Feeling) {
	person.state = setFeeling(feeling)
}

type angryState struct {
}

func newAngryState() *angryState {
	return &angryState{}
}

func getAngryState() *angryState {
	return aState
}

func (s *angryState) expressFeeling() string {
	return "angry"
}

// NOTE : stateがstateを設定するのはなんだか気持ち悪さがある
func (s *angryState) setFeeling(person *person, feeling Feeling) {
	person.state = setFeeling(feeling)
}

func setFeeling(feeling Feeling) State {
	switch feeling {
	case FeelingHappy:
		return getHappyState()
	case FeelingAngry:
		return getAngryState()
	default:
		panic("not supported feeling")
	}
}

type Feeling string

const (
	FeelingHappy = Feeling("happy")
	FeelingAngry = Feeling("angry")
	FeelingCry   = Feeling("cry")
)

type Context interface {
	SetFeeling(Feeling)
	ExpressFeeling() string
}

type person struct {
	state State
}

func NewPerson() *person {
	// デフォルトは happyState
	state := getHappyState()

	return &person{
		state: state,
	}
}

func (p *person) SetFeeling(feeling Feeling) {
	p.state.setFeeling(p, feeling)
}

func (p *person) ExpressFeeling() string {
	return p.state.expressFeeling()
}
