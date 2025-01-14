package unit
import(
	"T1/entity"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"testing"
	"fmt"
)
func TestPhoneNumber(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Check Phone`,func(t *testing.T){
		user := entity.User{
			Phone: "",
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is required"))
	})

	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		user := entity.User{
			Phone:     "080800000000", // ผิดตรงนี้ มี 11 ตัว
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", user.Phone)))

	})
}
