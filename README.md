![pepipostlogo](https://pepipost.com/wp-content/uploads/2017/07/P_LOGO.png)

[![Twitter Follow](https://img.shields.io/twitter/follow/pepi_post.svg?style=social&label=Follow)](https://twitter.com/pepi_post)

# Official Go Library for [Pepipost](http://www.pepipost.com/?utm_campaign=GitHubSDK&utm_medium=GithubSDK&utm_source=GithubSDK)

This library contains methods for easily interacting with the Pepipost Email Sending API to send emails within few seconds.

We are trying to make our libraries Community Driven- which means we need your help in building the right things in proper order we would request you to help us by sharing comments, creating new [issues](https://github.com/pepipost/pepipost-sdk-go/issues) or [pull requests](https://github.com/pepipost/pepipost-sdk-go/pulls).

We welcome any sort of contribution to this library.

The latest 5.0.0 version of this library provides is fully compatible with the latest Pepipost v5.0 API.

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

import(
  "fmt"
  "github.com/pepipost/pepipost-sdk-go/pepipost_lib/pepipost_pkg"
  "github.com/pepipost/pepipost-sdk-go/pepipost_lib/models_pkg"
)

func main(){

  ApiKey := "dfan3n4lk4212bjk39012hi3nrj1qk"

  client := PepipostClient.NewPEPIPOSTCLENT(ApiKey)
  send := client.Send()

  fromEmail := "hello@your-registered-domain-with-pepipost"
  fromName := "Pepipost"
  htmlBody := "<html><body>Hello [%NAME%], Email testing is successful. <br> Hope you enjoyed this integration. <br></html>"
  trueValue := true
  falseValue := false
  toName := "to-address"
  toEmail := "to-address@mydomain.name"

  Body := &models_pkg.Send{}

  Body.From = models_pkg.From{}
  Body.From.Email = &fromEmail
  Body.From.Name = &fromName

  Body.Subject = "Pepipost Test Email from GOlang SDK"
  Body.Content = make([]*models_pkg.Content,3)
  Body.Content[0] = &models_pkg.Content{}
  Body.Content[0].Type = models_pkg.Type_HTML
  Body.Content[0].Value = &htmlBody


  Body.Personalizations = make([]*models_pkg.Personalizations,3)
  Body.Personalizations[0] = &models_pkg.Personalizations{}

  var input interface{}
  Body.Personalizations[0].To = make([]*models_pkg.EmailStruct,3)
  Body.Personalizations[0].To[0] = &models_pkg.EmailStruct{}
  Body.Personalizations[0].To[0].Name = &toName
  Body.Personalizations[0].To[0].Email = &toEmail

  Body.Settings = &models_pkg.Settings{}
  Body.Settings.Footer = &trueValue
  Body.Settings.ClickTrack = &trueValue
  Body.Settings.OpenTrack = &falseValue
  Body.Settings.UnsubscribeTrack = &trueValue

  Body.Tags = &[]string{ "campaign" }

  var err error
  var result []string

  result, err = send.CreateGenerateTheMailSendRequest(Body)

  if err != nil{
    fmt.Println(err)
    //TODO: Use err variable here
  }else{
    //TODO: Use result variable here
    fmt.Println(result)
  }
}

```

<a name="announcements"></a>
# Announcements

v5.0.0 has been released! Please see the [release notes](https://github.com/pepipost/pepipost-sdk-go/releases/) for details.

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



