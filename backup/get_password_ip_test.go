package backup_test

import (
	"os"

	. "github.com/pivotalservices/cfops/backup"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("get_password_ip", func() {
	Describe("GetPasswordAndIP function", func() {
		Context("when given a valid installation.json", func() {
			var (
				jsonObj     InstallationCompareObject
				product     string = "cf"
				component   string = "ccdb"
				username    string = "admin"
				controlIp   string = "172.16.1.46"
				controlPass string = "e3e89a528625d819160d"
			)

			BeforeEach(func() {
				var fileRef *os.File
				defer fileRef.Close()
				fileRef, _ = os.Open("fixtures/installation.json")
				jsonObj, _ = ReadAndUnmarshal(fileRef)
			})

			AfterEach(func() {
			})

			It("Should return nil error, correct ip & password", func() {
				ip, password, err := GetPasswordAndIP(jsonObj, product, component, username)
				Ω(err).Should(BeNil())
				Ω(ip).Should(Equal(controlIp))
				Ω(password).Should(Equal(controlPass))
			})

			It("Should not panic on complete real world dataset", func() {
				Ω(func() {
					GetPasswordAndIP(jsonObj, product, component, username)
				}).ShouldNot(Panic())
			})
		})
	})

	Describe("IpPasswordParser struct", func() {
		Context("when given a valid installation.json", func() {
			var (
				parser      *IpPasswordParser
				jsonObj     InstallationCompareObject
				product     string = "cf"
				component   string = "ccdb"
				username    string = "admin"
				controlIp   string = "172.16.1.46"
				controlPass string = "e3e89a528625d819160d"
			)

			BeforeEach(func() {
				var fileRef *os.File
				defer fileRef.Close()
				fileRef, _ = os.Open("fixtures/installation.json")
				jsonObj, _ = ReadAndUnmarshal(fileRef)

				parser = &IpPasswordParser{
					Product:   product,
					Component: component,
					Username:  username,
				}
			})

			AfterEach(func() {
			})

			It("Should return nil error, correct ip & password", func() {
				ip, password, err := parser.Parse(jsonObj)
				Ω(err).Should(BeNil())
				Ω(ip).Should(Equal(controlIp))
				Ω(password).Should(Equal(controlPass))
			})

			It("Should not panic on complete real world dataset", func() {
				Ω(func() {
					parser.Parse(jsonObj)
				}).ShouldNot(Panic())
			})
		})

		Context("when no valid component found", func() {
			var (
				parser    *IpPasswordParser
				jsonObj   InstallationCompareObject
				product   string = "cf"
				component string = "aaaa"
				username  string = "admin"
			)

			BeforeEach(func() {
				var fileRef *os.File
				defer fileRef.Close()
				fileRef, _ = os.Open("fixtures/installation.json")
				jsonObj, _ = ReadAndUnmarshal(fileRef)

				parser = &IpPasswordParser{
					Product:   product,
					Component: component,
					Username:  username,
				}
			})

			AfterEach(func() {
			})

			It("Should return error", func() {
				ip, password, err := parser.Parse(jsonObj)
				Ω(err).ShouldNot(BeNil())
				Ω(ip).Should(BeEmpty())
				Ω(password).Should(BeEmpty())
			})

			It("Should not panic", func() {
				Ω(func() {
					parser.Parse(jsonObj)
				}).ShouldNot(Panic())
			})
		})

		Context("when no valid product found", func() {
			var (
				parser    *IpPasswordParser
				jsonObj   InstallationCompareObject
				product   string = "fail"
				component string = "ccdb"
				username  string = "admin"
			)

			BeforeEach(func() {
				var fileRef *os.File
				defer fileRef.Close()
				fileRef, _ = os.Open("fixtures/installation.json")
				jsonObj, _ = ReadAndUnmarshal(fileRef)

				parser = &IpPasswordParser{
					Product:   product,
					Component: component,
					Username:  username,
				}
			})

			AfterEach(func() {
			})

			It("Should return error", func() {
				ip, password, err := parser.Parse(jsonObj)
				Ω(err).ShouldNot(BeNil())
				Ω(ip).Should(BeEmpty())
				Ω(password).Should(BeEmpty())
			})

			It("Should not panic", func() {
				Ω(func() {
					parser.Parse(jsonObj)
				}).ShouldNot(Panic())
			})
		})

		Context("when no valid username found", func() {
			var (
				parser    *IpPasswordParser
				jsonObj   InstallationCompareObject
				product   string = "cf"
				component string = "ccdb"
				username  string = "fail"
			)

			BeforeEach(func() {
				var fileRef *os.File
				defer fileRef.Close()
				fileRef, _ = os.Open("fixtures/installation.json")
				jsonObj, _ = ReadAndUnmarshal(fileRef)

				parser = &IpPasswordParser{
					Product:   product,
					Component: component,
					Username:  username,
				}
			})

			AfterEach(func() {
			})

			It("Should return error", func() {
				ip, password, err := parser.Parse(jsonObj)
				Ω(err).ShouldNot(BeNil())
				Ω(ip).Should(BeEmpty())
				Ω(password).Should(BeEmpty())
			})

			It("Should not panic", func() {
				Ω(func() {
					parser.Parse(jsonObj)
				}).ShouldNot(Panic())
			})
		})

	})
})