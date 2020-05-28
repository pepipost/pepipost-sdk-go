/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package mailsend_pkg

import "pepipost_lib/configuration_pkg"
import "pepipost_lib/models_pkg"

/*
 * Interface for the MAILSEND_IMPL
 */
type MAILSEND interface {
    CreateGeneratethemailsendrequest (*models_pkg.Send) (interface{}, error)
}

/*
 * Factory for the MAILSEND interaface returning MAILSEND_IMPL
 */
func NewMAILSEND(config configuration_pkg.CONFIGURATION) *MAILSEND_IMPL {
    client := new(MAILSEND_IMPL)
    client.config = config
    return client
}