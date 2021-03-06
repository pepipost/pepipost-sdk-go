/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg



/*
 * Structure for the custom type Content
 */
type Content struct {
    Type            TypeEnum        `json:"type,omitempty" form:"type,omitempty"` //TODO: Write general description for this field
    Value           *string         `json:"value,omitempty" form:"value,omitempty"` //HTML content to be sent in your email
}

/*
 * Structure for the custom type Personalizations
 */
type Personalizations struct {
    Attributes      *interface{}    `json:"attributes,omitempty" form:"attributes,omitempty"` //Dynamic attributes
    Headers         *interface{}    `json:"headers,omitempty" form:"headers,omitempty"` //Dynamic headers attributes
    Attachments     []*Attachments  `json:"attachments,omitempty" form:"attachments,omitempty"` //Attachments to individuals recipient
    To              []*EmailStruct  `json:"to" form:"to"` //To email-address
    Cc              []*EmailStruct  `json:"cc,omitempty" form:"cc,omitempty"` //CC email-address
    Bcc             []*EmailStruct  `json:"bcc,omitempty" form:"bcc,omitempty"` //Bcc email-addresses
    TokenTo         *string         `json:"token_to,omitempty" form:"token_to,omitempty"` //token to which is json string
    TokenCc         *string         `json:"token_cc,omitempty" form:"token_cc,omitempty"` //token cc which is json string
    TokenBcc        *string         `json:"token_bcc,omitempty" form:"token_bcc,omitempty"` //token bcc which is json string
}

/*
 * Structure for the custom type From
 */
type From struct {
    Email           *string         `json:"email,omitempty" form:"email,omitempty"` //TODO: Write general description for this field
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Settings
 */
type Settings struct {
    Footer            *bool           `json:"footer,omitempty" form:"footer,omitempty"` //enable or disable footer
    ClickTrack        *bool           `json:"click_track,omitempty" form:"click_track,omitempty"` //enable or disable click tracking
    OpenTrack         *bool           `json:"open_track,omitempty" form:"open_track,omitempty"` //enable or disable open tracking
    UnsubscribeTrack  *bool           `json:"unsubscribe_track,omitempty" form:"unsubscribe_track,omitempty"` //enable or disable unsubscribe tracking
    Hepf              *bool           `json:"hepf,omitempty" form:"hepf,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type EmailStruct
 */
type EmailStruct struct {
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //Name of recipient
    Email           *string         `json:"email,omitempty" form:"email,omitempty"` //Email of recipient
}

/*
 * Structure for the custom type Attachments
 */
type Attachments struct {
    Content         *string         `json:"content,omitempty" form:"content,omitempty"` //Base64 encoded value of the attached file
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //filename of attachments
}

/*
 * Structure for the custom type Send
 */
type Send struct {
    ReplyTo          *string         `json:"reply_to,omitempty" form:"reply_to,omitempty"` //email address which recipients can reply to.
    From             From            `json:"from" form:"from"` //Email address representing the sender of the mail
    Subject          string          `json:"subject" form:"subject"` //Subject line of the email
    TemplateId       *int64          `json:"template_id,omitempty" form:"template_id,omitempty"` //ID of the template to be used for sending the mail
    Content          []*Content      `json:"content" form:"content"` //content in text/plain format
    Attachments      []*Attachments  `json:"attachments,omitempty" form:"attachments,omitempty"` //attachment information
    Personalizations []*Personalizations `json:"personalizations" form:"personalizations"` //to recipient with some personalized data like to address, attachments and attributes
    Settings         *Settings       `json:"settings,omitempty" form:"settings,omitempty"` //TODO: Write general description for this field
    Tags             *[]string       `json:"tags,omitempty" form:"tags,omitempty"` //define custom tags to organize your emails
    LintPayload      *bool           `json:"lint_payload,omitempty" form:"lint_payload,omitempty"` //TODO: Write general description for this field
    Schedule         *int64          `json:"schedule,omitempty" form:"schedule,omitempty"` //schedule the time of email delivery
    Bcc              []*EmailStruct  `json:"bcc,omitempty" form:"bcc,omitempty"` //Global bcc can be defined here
}

/*
 * Structure for the custom type DomainStruct
 */
type DomainStruct struct {
    Domain          string          `json:"domain" form:"domain"` //The domain you wish to include in the 'From' header of your emails.
    EnvelopeName    *string         `json:"envelopeName,omitempty" form:"envelopeName,omitempty"` //The subdomain which will be used for tracking opens, clicks and unsubscribe
}

/*
 * Structure for the custom type Deletedomain
 */
type Deletedomain struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Name of the domain
}

/*
 * Structure for the custom type AddemailordomaintoSuppressionlist
 */
type AddemailordomaintoSuppressionlist struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //Add the domain to be suppressed here. We will not deliver emails to recipients email addresses with this domain.<br>\nComma separate the values to suppress multiple domains..
    Email           *string         `json:"email,omitempty" form:"email,omitempty"` //Add an email address to be suppressed here. We will not deliver emails to this email address.<br>\nComma separate the values to suppress multiple email addresses
}

/*
 * Structure for the custom type RemoveemailordomaintoSuppressionlist
 */
type RemoveemailordomaintoSuppressionlist struct {
    Domain          *string         `json:"domain,omitempty" form:"domain,omitempty"` //List one or more recipient domains to be removed from the suppression list here. <br>\nComma separate the values to suppress multiple recipient domains.
    Email           *string         `json:"email,omitempty" form:"email,omitempty"` //List one or more email addresses to be removed from the suppression list here. <br>\nComma separate the values to suppress multiple email addresses.
}

/*
 * Structure for the custom type Createsubaccount
 */
type Createsubaccount struct {
    Username        string          `json:"username" form:"username"` //provide a username for the subaccount
    Email           string          `json:"email" form:"email"` //email address to be registered with the account.
    Setpassword     string          `json:"setpassword" form:"setpassword"` //You can opt to set the password for the subaccount.\nIf set as 1, please provide a value in password parameter.\nIf set as 0, the email confirmation link will act as a password reset link.
    Password        *string         `json:"password,omitempty" form:"password,omitempty"` //It is required to pass this value, if setpassword is set as 1.\nThe password must comprise minimum of 8 characters and include one uppercase character, one lowercase character, one numeric character.
    CreditType      *string         `json:"credit_type,omitempty" form:"credit_type,omitempty"` //Allowed values one_time_credit or unlimited by default, all subaccounts are created with credit type as unlimited.
}

/*
 * Structure for the custom type Updatesubaccount
 */
type Updatesubaccount struct {
    Username         *string         `json:"username,omitempty" form:"username,omitempty"` //The username for the subaccount
    NewEmail         *string         `json:"new_email,omitempty" form:"new_email,omitempty"` //The new email address to be registered with the subaccount
    NewPassword      *string         `json:"new_password,omitempty" form:"new_password,omitempty"` //The new password of the subaccount
    ConfirmPassword  *string         `json:"confirm_password,omitempty" form:"confirm_password,omitempty"` //reconfirm the new password for the subaccount
    CreditType       *string         `json:"credit_type,omitempty" form:"credit_type,omitempty"` //Allowed values one_time_credit or unlimited by default, all subaccounts are created with credit type as unlimited.
}

/*
 * Structure for the custom type UpdateCredisofsubaccount
 */
type UpdateCredisofsubaccount struct {
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //The username of the subaccount
    Action          ActionEnum      `json:"action,omitempty" form:"action,omitempty"` //Indicate the action (add or subtract) to be taken.Allowed values increase, decrease
    Amount          *int64          `json:"amount,omitempty" form:"amount,omitempty"` //Amount of credits
}

/*
 * Structure for the custom type UpdaterecurringCredisofsubaccount
 */
type UpdaterecurringCredisofsubaccount struct {
    Username         *string         `json:"username,omitempty" form:"username,omitempty"` //The username of the subaccount
    RecurringCredit  *int64          `json:"recurring_credit,omitempty" form:"recurring_credit,omitempty"` //The amount to be added periodically
    Timeperiod       TimeperiodEnum  `json:"timeperiod,omitempty" form:"timeperiod,omitempty"` //The periodic \n\nAllowed values  \"daily\", \"weekly\", \"monhtly\"
    StartDate        *string         `json:"start_date,omitempty" form:"start_date,omitempty"` //The date from which the credit allocation will commence
    EndDate          *string         `json:"end_date,omitempty" form:"end_date,omitempty"` //The last date of credit allocation
    Status           *string         `json:"status,omitempty" form:"status,omitempty"` //Flag to enable or disable the recurring credit allocation
}

/*
 * Structure for the custom type Deletesubacoount
 */
type Deletesubacoount struct {
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //The username of the subaccount
}

/*
 * Structure for the custom type Enableordisablesubacoount
 */
type Enableordisablesubacoount struct {
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //The username of the subaccount
    Disabled        *bool           `json:"disabled,omitempty" form:"disabled,omitempty"` //Flag to indicate whether the subaccount should be enabled or disabled.
}
