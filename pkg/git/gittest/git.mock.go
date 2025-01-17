// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pipe-cd/pipecd/pkg/git (interfaces: Repo)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod --package=gittest --destination=pkg/git/gittest/git.mock.go github.com/pipe-cd/pipecd/pkg/git Repo
//

// Package gittest is a generated GoMock package.
package gittest

import (
	context "context"
	reflect "reflect"

	git "github.com/pipe-cd/pipecd/pkg/git"
	gomock "go.uber.org/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
	isgomock struct{}
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// ChangedFiles mocks base method.
func (m *MockRepo) ChangedFiles(ctx context.Context, from, to string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangedFiles", ctx, from, to)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangedFiles indicates an expected call of ChangedFiles.
func (mr *MockRepoMockRecorder) ChangedFiles(ctx, from, to any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangedFiles", reflect.TypeOf((*MockRepo)(nil).ChangedFiles), ctx, from, to)
}

// Checkout mocks base method.
func (m *MockRepo) Checkout(ctx context.Context, commitish string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkout", ctx, commitish)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checkout indicates an expected call of Checkout.
func (mr *MockRepoMockRecorder) Checkout(ctx, commitish any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkout", reflect.TypeOf((*MockRepo)(nil).Checkout), ctx, commitish)
}

// CheckoutPullRequest mocks base method.
func (m *MockRepo) CheckoutPullRequest(ctx context.Context, number int, branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckoutPullRequest", ctx, number, branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckoutPullRequest indicates an expected call of CheckoutPullRequest.
func (mr *MockRepoMockRecorder) CheckoutPullRequest(ctx, number, branch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckoutPullRequest", reflect.TypeOf((*MockRepo)(nil).CheckoutPullRequest), ctx, number, branch)
}

// Clean mocks base method.
func (m *MockRepo) Clean() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clean")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clean indicates an expected call of Clean.
func (mr *MockRepoMockRecorder) Clean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockRepo)(nil).Clean))
}

// CleanPath mocks base method.
func (m *MockRepo) CleanPath(ctx context.Context, relativePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanPath", ctx, relativePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanPath indicates an expected call of CleanPath.
func (mr *MockRepoMockRecorder) CleanPath(ctx, relativePath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanPath", reflect.TypeOf((*MockRepo)(nil).CleanPath), ctx, relativePath)
}

// CommitChanges mocks base method.
func (m *MockRepo) CommitChanges(ctx context.Context, branch, message string, newBranch bool, changes map[string][]byte, trailers map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitChanges", ctx, branch, message, newBranch, changes, trailers)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitChanges indicates an expected call of CommitChanges.
func (mr *MockRepoMockRecorder) CommitChanges(ctx, branch, message, newBranch, changes, trailers any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitChanges", reflect.TypeOf((*MockRepo)(nil).CommitChanges), ctx, branch, message, newBranch, changes, trailers)
}

// Copy mocks base method.
func (m *MockRepo) Copy(dest string) (git.Worktree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", dest)
	ret0, _ := ret[0].(git.Worktree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Copy indicates an expected call of Copy.
func (mr *MockRepoMockRecorder) Copy(dest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockRepo)(nil).Copy), dest)
}

// CopyToModify mocks base method.
func (m *MockRepo) CopyToModify(dest string) (git.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyToModify", dest)
	ret0, _ := ret[0].(git.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyToModify indicates an expected call of CopyToModify.
func (mr *MockRepoMockRecorder) CopyToModify(dest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyToModify", reflect.TypeOf((*MockRepo)(nil).CopyToModify), dest)
}

// GetClonedBranch mocks base method.
func (m *MockRepo) GetClonedBranch() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClonedBranch")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClonedBranch indicates an expected call of GetClonedBranch.
func (mr *MockRepoMockRecorder) GetClonedBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClonedBranch", reflect.TypeOf((*MockRepo)(nil).GetClonedBranch))
}

// GetCommitForRev mocks base method.
func (m *MockRepo) GetCommitForRev(ctx context.Context, rev string) (git.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitForRev", ctx, rev)
	ret0, _ := ret[0].(git.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitForRev indicates an expected call of GetCommitForRev.
func (mr *MockRepoMockRecorder) GetCommitForRev(ctx, rev any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitForRev", reflect.TypeOf((*MockRepo)(nil).GetCommitForRev), ctx, rev)
}

// GetLatestCommit mocks base method.
func (m *MockRepo) GetLatestCommit(ctx context.Context) (git.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestCommit", ctx)
	ret0, _ := ret[0].(git.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestCommit indicates an expected call of GetLatestCommit.
func (mr *MockRepoMockRecorder) GetLatestCommit(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestCommit", reflect.TypeOf((*MockRepo)(nil).GetLatestCommit), ctx)
}

// GetPath mocks base method.
func (m *MockRepo) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockRepoMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockRepo)(nil).GetPath))
}

// ListCommits mocks base method.
func (m *MockRepo) ListCommits(ctx context.Context, visionRange string) ([]git.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", ctx, visionRange)
	ret0, _ := ret[0].([]git.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockRepoMockRecorder) ListCommits(ctx, visionRange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockRepo)(nil).ListCommits), ctx, visionRange)
}

// MergeRemoteBranch mocks base method.
func (m *MockRepo) MergeRemoteBranch(ctx context.Context, branch, commit, mergeCommitMessage string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeRemoteBranch", ctx, branch, commit, mergeCommitMessage)
	ret0, _ := ret[0].(error)
	return ret0
}

// MergeRemoteBranch indicates an expected call of MergeRemoteBranch.
func (mr *MockRepoMockRecorder) MergeRemoteBranch(ctx, branch, commit, mergeCommitMessage any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeRemoteBranch", reflect.TypeOf((*MockRepo)(nil).MergeRemoteBranch), ctx, branch, commit, mergeCommitMessage)
}

// Pull mocks base method.
func (m *MockRepo) Pull(ctx context.Context, branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", ctx, branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockRepoMockRecorder) Pull(ctx, branch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockRepo)(nil).Pull), ctx, branch)
}

// Push mocks base method.
func (m *MockRepo) Push(ctx context.Context, branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", ctx, branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockRepoMockRecorder) Push(ctx, branch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockRepo)(nil).Push), ctx, branch)
}
