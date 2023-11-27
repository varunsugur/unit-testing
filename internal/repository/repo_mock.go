// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go
//
// Generated by this command:
//
//	mockgen -source=repo.go -destination=repo_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	models "golang/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CheckEmail mocks base method.
func (m *MockUserRepo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmail", ctx, email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmail indicates an expected call of CheckEmail.
func (mr *MockUserRepoMockRecorder) CheckEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmail", reflect.TypeOf((*MockUserRepo)(nil).CheckEmail), ctx, email)
}

// CreatCompany mocks base method.
func (m *MockUserRepo) CreatCompany(ctx context.Context, data models.Company) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatCompany", ctx, data)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatCompany indicates an expected call of CreatCompany.
func (mr *MockUserRepoMockRecorder) CreatCompany(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatCompany", reflect.TypeOf((*MockUserRepo)(nil).CreatCompany), ctx, data)
}

// CreatUser mocks base method.
func (m *MockUserRepo) CreatUser(ctx context.Context, userData models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatUser", ctx, userData)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatUser indicates an expected call of CreatUser.
func (mr *MockUserRepoMockRecorder) CreatUser(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatUser", reflect.TypeOf((*MockUserRepo)(nil).CreatUser), ctx, userData)
}

// CreateJob mocks base method.
func (m *MockUserRepo) CreateJob(ctx context.Context, jobData models.Job) (models.ResponseJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, jobData)
	ret0, _ := ret[0].(models.ResponseJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockUserRepoMockRecorder) CreateJob(ctx, jobData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockUserRepo)(nil).CreateJob), ctx, jobData)
}

// FindAllJobs mocks base method.
func (m *MockUserRepo) FindAllJobs(ctx context.Context) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllJobs", ctx)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllJobs indicates an expected call of FindAllJobs.
func (mr *MockUserRepoMockRecorder) FindAllJobs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllJobs", reflect.TypeOf((*MockUserRepo)(nil).FindAllJobs), ctx)
}

// FindJob mocks base method.
func (m *MockUserRepo) FindJob(ctx context.Context, cid uint64) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindJob", ctx, cid)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindJob indicates an expected call of FindJob.
func (mr *MockUserRepoMockRecorder) FindJob(ctx, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindJob", reflect.TypeOf((*MockUserRepo)(nil).FindJob), ctx, cid)
}

// GetTheJobData mocks base method.
func (m *MockUserRepo) GetTheJobData(jobid uint) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheJobData", jobid)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTheJobData indicates an expected call of GetTheJobData.
func (mr *MockUserRepoMockRecorder) GetTheJobData(jobid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheJobData", reflect.TypeOf((*MockUserRepo)(nil).GetTheJobData), jobid)
}

// ResetPassword mocks base method.
func (m *MockUserRepo) ResetPassword(email, resetpassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", email, resetpassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserRepoMockRecorder) ResetPassword(email, resetpassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserRepo)(nil).ResetPassword), email, resetpassword)
}

// VerifyUser mocks base method.
func (m *MockUserRepo) VerifyUser(vu models.VerifyUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", vu)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockUserRepoMockRecorder) VerifyUser(vu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockUserRepo)(nil).VerifyUser), vu)
}

// ViewCompanies mocks base method.
func (m *MockUserRepo) ViewCompanies(ctx context.Context) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompanies", ctx)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewCompanies indicates an expected call of ViewCompanies.
func (mr *MockUserRepoMockRecorder) ViewCompanies(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanies", reflect.TypeOf((*MockUserRepo)(nil).ViewCompanies), ctx)
}

// ViewCompanyById mocks base method.
func (m *MockUserRepo) ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompanyById", ctx, cid)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewCompanyById indicates an expected call of ViewCompanyById.
func (mr *MockUserRepoMockRecorder) ViewCompanyById(ctx, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanyById", reflect.TypeOf((*MockUserRepo)(nil).ViewCompanyById), ctx, cid)
}

// ViewJobDetailsBy mocks base method.
func (m *MockUserRepo) ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobDetailsBy", ctx, jid)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobDetailsBy indicates an expected call of ViewJobDetailsBy.
func (mr *MockUserRepoMockRecorder) ViewJobDetailsBy(ctx, jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobDetailsBy", reflect.TypeOf((*MockUserRepo)(nil).ViewJobDetailsBy), ctx, jid)
}
