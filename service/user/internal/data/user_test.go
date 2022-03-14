package data_test

import (
	"user/internal/biz"
	"user/internal/data"
	"user/internal/testdata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User
	BeforeEach(func() {
		// 这里的 Db 是 data_suite_test.go 文件里面定义的
		ro = data.NewUserRepo(Db, nil)
		// 这里你可以引入外部组装好的数据
		uD = testdata.User()
	})

	// 设置 It 块来添加单个规格
	It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		// 组装的数据 mobile 为 13803881388
		Ω(u.Mobile).Should(Equal("13509876789")) // 手机号应该为创建的时候写入的手机号
	})
	It("ListUser", func() {
		user, total, err := ro.ListUser(ctx, 1, 10)
		Ω(err).ShouldNot(HaveOccurred()) // 获取列表不应该出现错误
		Ω(user).ShouldNot(BeEmpty())     // 结果不应该为空
		Ω(total).Should(Equal(1))        // 总数应该为 1，因为上面只创建了一条
		Ω(len(user)).Should(Equal(1))
		Ω(user[0].Mobile).Should(Equal("13509876789"))
	})
})
