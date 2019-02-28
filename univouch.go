package main

import (
	"github.com/dim13/unifi"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"strconv"
)

var (
	serve    = kingpin.Flag("serve", "Serve as printable codes").Bool()
	generate = kingpin.Flag("generate", "Generate vouchers").Bool()
	host     = kingpin.Flag("controller", "Controller IP").Required().IP()
	username = kingpin.Flag("username", "Controller username").Required().String()
	password = kingpin.Flag("password", "Controller password").Required().String()
	port     = kingpin.Flag("port", "Controller port").Default("8443").Int()
	num      = kingpin.Flag("vouchers", "Amount of vouchers to generate").Default("1").Int()
	duration = kingpin.Flag("expiration", "Expiration in minutes").Default("67").Int()
	note     = kingpin.Flag("note", "Voucher note").Default("Generated Vouchers").String()
)

var api *unifi.Unifi

func its(t int) string {
	return strconv.Itoa(t)
}

func main() {
	kingpin.Parse()
	var err error
	api, err = unifi.Login(*username, *password, (*host).String(), its(*port), "default", 4)
	if err != nil {
		logrus.Fatal("Failed to login: ", err.Error())
	}
	defer api.Logout()

	r := gin.Default()

	sites, err := api.Sites()
	if err != nil {
		logrus.Fatal("Failed to fetch sites: ", err.Error())
	}

	if *generate {
		_, err = api.NewVoucher(&sites[0], unifi.NewVoucher{
			Cmd:          "create-voucher",
			Expire:       "custom",
			ExpireNumber: its(*duration),
			ExpireUnit:   "1",
			N:            its(*num),
			Note:         *note,
			Quota:        its(*num),
		})
		if err != nil {
			logrus.Fatal("Failed to create voucher: ", err.Error())
		}
	}

	r.LoadHTMLGlob("ui/*.html")
	r.Static("/assets", "assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "vouchers.html", nil)
	})
	r.GET("/api/vouchers", func(c *gin.Context) {

		allVouchers, err := api.VoucherMap(&sites[0])
		if err != nil {
			logrus.Fatal("Failed to get vouchers")
		}

		var vouchers []unifi.Voucher
		for e := range allVouchers {
			if allVouchers[e].Used == 0 {
				vouchers = append(vouchers, allVouchers[e])
			}
		}

		c.IndentedJSON(200, vouchers)
	})

	if *serve {
		if r.Run(":3000") != nil {
			logrus.Fatal("Failed to run server")
		}
	}
}
