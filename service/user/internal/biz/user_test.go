package biz_test

import (
	"user/internal/biz"
	"user/internal/mocks/mrepo"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserUsecase", func() {
	var userCase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userCase = biz.NewUserUsecase(mUserRepo, nil)
	})

	It("Create", func() {
		info := &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: 693629981,
		}
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.Create(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.ID).To(Equal(int64(1)))
		Ω(l.Mobile).To(Equal("13803881388"))
	})

})
