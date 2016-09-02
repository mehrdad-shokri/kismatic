package cli

import "github.com/apprenda/kismatic-platform/pkg/install"

type fakePlan struct {
	exists     bool
	plan       *install.Plan
	err        error
	readCalled bool
}

func (fp *fakePlan) PlanExists() bool { return fp.exists }
func (fp *fakePlan) Read() (*install.Plan, error) {
	fp.readCalled = true
	return fp.plan, fp.err
}
func (fp *fakePlan) Write(p *install.Plan) error {
	fp.plan = p
	return fp.err
}

type fakeExecutor struct {
	installCalled bool
	err           error
}

func (fe *fakeExecutor) Install(p *install.Plan) error {
	fe.installCalled = true
	return fe.err
}

type fakePKI struct {
	called bool
	err    error
}

func (fp *fakePKI) GenerateClusterCerts(p *install.Plan) error {
	fp.called = true
	return fp.err
}
