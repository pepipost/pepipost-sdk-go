/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package domaindelete_pkg

import "pepipost_lib/configuration_pkg"
import "pepipost_lib/models_pkg"

/*
 * Interface for the DOMAINDELETE_IMPL
 */
type DOMAINDELETE interface {
    Deletedomain (*models_pkg.Deletedomain) (interface{}, error)
}

/*
 * Factory for the DOMAINDELETE interaface returning DOMAINDELETE_IMPL
 */
func NewDOMAINDELETE(config configuration_pkg.CONFIGURATION) *DOMAINDELETE_IMPL {
    client := new(DOMAINDELETE_IMPL)
    client.config = config
    return client
}