/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package setrecurringcreditddetails_pkg

import "pepipost_lib/configuration_pkg"
import "pepipost_lib/models_pkg"

/*
 * Interface for the SETRECURRINGCREDITDDETAILS_IMPL
 */
type SETRECURRINGCREDITDDETAILS interface {
    CreateSetrecurringcreditddetailsPOST (*models_pkg.UpdaterecurringCredisofsubaccount) (interface{}, error)
}

/*
 * Factory for the SETRECURRINGCREDITDDETAILS interaface returning SETRECURRINGCREDITDDETAILS_IMPL
 */
func NewSETRECURRINGCREDITDDETAILS(config configuration_pkg.CONFIGURATION) *SETRECURRINGCREDITDDETAILS_IMPL {
    client := new(SETRECURRINGCREDITDDETAILS_IMPL)
    client.config = config
    return client
}