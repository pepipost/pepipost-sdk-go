![pepipostlogo](https://pepipost.com/wp-content/uploads/2017/07/P_LOGO.png)

[![Twitter Follow](https://img.shields.io/twitter/follow/pepi_post.svg?style=social&label=Follow)](https://twitter.com/pepi_post)

# Official Go Library for [Pepipost](http://www.pepipost.com/?utm_campaign=GitHubSDK&utm_medium=GithubSDK&utm_source=GithubSDK)

This library contains methods for easily interacting with the Pepipost Email Sending API to send emails within few seconds.

We are trying to make our libraries Community Driven- which means we need your help in building the right things in proper order we would request you to help us by sharing comments, creating new [issues](https://github.com/pepipost/pepipost-sdk-go/issues) or [pull requests](https://github.com/pepipost/pepipost-sdk-go/pulls).

We welcome any sort of contribution to this library.

The latest 2.5.0 version of this library provides is fully compatible with the latest Pepipost v2.0 API.

For any update of this library check [Releases](https://github.com/pepipost/pepipost-sdk-go/releases).

# Table of Content
  
* [Installation](#installation)
* [Quick Start](#quick-start)
* [Sample Example](#example)
* [Announcements](#announcements)
* [Roadmap](#roadmap)
* [About](#about)
* [License](#license)

<a name="installation"></a>
# Installation

* In order to successfully build and run your SDK files, you are required to have the following Prerequisites in you system:

<a name="prereq"></a>
## Prerequisites

* **GO** (visit [Go installation page](https://golang.org/doc/install) for details on how to install Go)
* Ensure that **```GOROOT```** & **```GOPATH```** enviroment variable is set in the system variables. if not, set it to your working directory where you add your Go projects.

<a name='quick-start'></a>
## Quick Start

Quick guide for installing Pepipost Go library 

* Check **GOPATH** using below command

  ``` echo $GOPATH ```
  
  ![pic1](http://app1.falconide.com/integration_imgs/goimg/capture(24).png)

* Change your directory to **$GOPATH** 
  
  ``` cd $GOPATH```

* make project using below command 

  ```  mkdir $GOPATH/src/pepipost_test && cd $GOPATH/src/ ```
  
  ![pic7](http://app1.falconide.com/integration_imgs/goimg/capture(28).png)

* Clone or download the repository using below command

  ``` git clone https://github.com/pepipost/pepipost-sdk-go.git ```
  
  **OR**
  
  Download **[zip](https://github.com/pepipost/pepipost-sdk-go/archive/master.zip)**


* Create sample Go file named **main.go** in 

``` cd $GOPATH/src/pepipost_test ```

  copy the below [sample example](#example) in your ```main.go``` file.
  
* Change your api-key and sending domain respectively

    * **apikey** will be available under Login to Pepipost -> Settings -> Integration  
    * **Sending Domain** will be available under Login to Pepiost -> Settings -> Sending Domains 
  
    ```
  *Note :: Domains showing with Active status on Sending Domain dashboard are only allowed to send any sort of emails.* In case there are no Sending Domain added under your account, then first add the domain, get the DNS (SPF/DKIM) settings done and get it reviewed by our compliance team for approval. Once the domain is approved, it will be in ACTIVE status and will be ready to send any sort of emails. 
    ```
* Sending first email using Go library

  ```go build main.go```
  
  ```./main.exe ``` (for windows) **OR** ```./main```
  
  ![img](http://app1.falconide.com/integration_imgs/goimg/capture(30).png)
 

<a name='example'></a>
## Sample Example  

```Go
package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/models_pkg"
	"github.com/pepipost/pepipost-sdk-go/pepipost_lib/pepipost_pkg"
)
func main() {
	
  client := PepipostClient.NewPEPIPOST()
	email := client.Email()
	ApiKey := "e8a874fbd6-XXXXXXXX-f0879205"

	username := "toemails@gmail.com"
	usernamecc := []string{"toemail_cc@gmail.com"}
	usernamebcc := []string{"tobcc@netcore.co.in"}
	fromEmail := "punit@register_domain.com"
	subject := "GO-SDK -- [% name %]" 
	body := "This is Go-SDK "

	xhead := map[string]string{
		"x-custom1": "amp-content-header",
		"x-custom2": "google-amp1",
	}

	attribute := map[string]string{
		"name" : "pepipost",
		"test_no" : "1234",
	}

	Body := &models_pkg.EmailBody{}
	Body.Personalizations = make([]*models_pkg.Personalizations, 3)
	Body.Personalizations[0] = &models_pkg.Personalizations{}
	Body.Personalizations[0].Recipient = &username
	Body.Personalizations[0].RecipientCc = &usernamecc
	Body.Personalizations[0].RecipientBcc = &usernamebcc
	Body.Personalizations[0].Xheader = xhead
	Body.Personalizations[0].Attributes = attribute

	Body.From = models_pkg.From{}
	Body.From.FromEmail = &fromEmail
	Body.Subject = &subject
	Body.Content = &body

	b, err1 := ioutil.ReadFile("/tmp/ampcontent") // just pass the file name
	if err1 != nil {
		fmt.Print("error= ")
		fmt.Print(err1)
	}
	//pass your custom URI (if given by PEPIPOST TEAM)
	BASE_URI := ""
	if len(os.Args) == 2 {
		BASE_URI = os.Args[1]
	}
	str := string(b)
	Body.AmpContent = &str

	var err error
	var result *models_pkg.SendEmailResponse
	result, err = email.CreateSendEmail(&ApiKey, Body, BASE_URI)

	if err != nil {
		//TODO: Use err variable here
		fmt.Print("result = ")
		fmt.Println(result)
	} else {
		//TODO: Use result variable here
		fmt.Print("resu=")
		fmt.Println(err)
	}
}
```

<a name="announcements"></a>
# Announcements

v2.6.1 has been released! Please see the [release notes](https://github.com/pepipost/pepipost-sdk-go/releases/) for details.

All updates to this library are documented in our [releases](https://github.com/pepipost/pepipost-sdk-go/releases). For any queries, feel free to reach out us at dx@pepipost.com

<a name="roadmap"></a>
## Roadmap

If you are interested in the future direction of this project, please take a look at our open [issues](https://github.com/pepipost/pepipost-sdk-go/issues) and [pull requests](https://github.com/pepipost/pepipost-sdk-go/pulls). We would love to hear your feedback.

<a name="about"></a>
## About
pepipost-sdk-go library is guided and supported by the [Pepipost Developer Experience Team](https://github.com/orgs/pepipost/teams/pepis/members) .
This pepipost library is maintained and funded by Pepipost Ltd. The names and logos for pepipost are trademarks of Pepipost Ltd.

<a name="license"></a>
## License
This code library was semi-automatically generated by APIMATIC v2.0 and licensed under The MIT License (MIT).



