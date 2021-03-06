package bdd

// Matcher can check if a passed-in value matches the matcher's expectations.
// Depending on the matcher arguments are required for the matching.
type Matcher interface {

	// Apply applies the matcher to the passed-in data and returns a Result.
	Apply(obtained interface{}, args []interface{}) Result
}

// MatcherArgsRequierer can require a certain amount of arguments for a matcher.
type MatcherArgsRequierer interface {

	// MinArgs is the minimum number of arguments for a matcher.
	MinArgs() int
}

// MatcherArgsLimiter can limit the amount of arguments for a matcher.
type MatcherArgsLimiter interface {

	// MinArgs is the maximum number of arguments for a matcher.
	MaxArgs() int
}

// namedMatcher can give a matcher a descriptive name.
type namedMatcher interface {

	// Name is the descriptive name of a matcher.
	Name() string
}

// Result is the result from applying a Matcher.
type Result struct {

	// Success is whether a match was successful.
	Success bool

	// Error is an error that occurred during matching.
	Error error

	// FailureMessage is the message for an unsuccessful match.
	FailureMessage string

	// NegatedFailureMessage is the negated message for an unsuccessful match.
	NegatedFailureMessage string
}

type matcher struct {
	minArgs int
	maxArgs int
	name    string
	apply   func(obtained interface{}, args []interface{}) Result
}

func (m *matcher) Name() string {
	return m.name
}

func (m *matcher) MinArgs() int {
	return m.minArgs
}

func (m *matcher) MaxArgs() int {
	return m.maxArgs
}

func (m *matcher) Apply(obtained interface{}, args []interface{}) Result {
	return m.apply(obtained, args)
}

type MatcherBuilder struct {
	m *matcher
}

func NewMatcherBuilder() *MatcherBuilder {
	return &MatcherBuilder{&matcher{}}
}

func (mb *MatcherBuilder) SetName(name string) *MatcherBuilder {
	mb.m.name = name
	return mb
}

func (mb *MatcherBuilder) SetMaxArgs(max int) *MatcherBuilder {
	mb.m.maxArgs = max
	return mb
}

func (mb *MatcherBuilder) SetMinArgs(min int) *MatcherBuilder {
	mb.m.minArgs = min
	return mb
}

func (mb *MatcherBuilder) SetApply(apply func(obtained interface{}, args []interface{}) Result) *MatcherBuilder {
	mb.m.apply = apply
	return mb
}

func (mb *MatcherBuilder) Build() Matcher {
	return mb.m
}
